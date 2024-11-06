import { BusinessError, HttpError } from '$/lib/api/error.ts'
import { errorResponse } from '$/lib/api/types.ts'
import { z } from 'zod'

const parseResponse = async <T extends z.ZodType>(response: Response, respZodObject: T): Promise<z.infer<T>> => {
  let resp: unknown
  try {
    resp = await response.json()
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
  } catch (e) {
    throw new HttpError(response, 'Failed to parse response: ' + response)
  }

  if (response.ok) {
    return respZodObject.parse(resp)
  }

  // parse as errorResponse
  const errorResp = errorResponse.safeParse(resp)
  if (!errorResp.success) {
    throw new HttpError(response, 'Failed to parse response: ' + resp)
  }
  throw new BusinessError(errorResp.data.code, errorResp.data.message)
}

/**
 * Fetch with type validating, error handling, and base URL
 */
export const fetchWrapped = async <T extends z.ZodType>(
  input: RequestInfo | URL,
  respZodObject: T,
  init?: RequestInit,
): Promise<z.infer<T>> => {
  const response = await fetch(input, init)
  return parseResponse(response, respZodObject)
}

export default fetchWrapped
