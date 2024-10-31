import { BusinessError, HttpError } from '$/lib/api/error.ts'
import { z } from 'zod'

import { baseResponse, baseResponseNull } from './types.ts'

const parseResponse = async <T extends z.ZodType>(response: Response, respZodObject: T): Promise<z.infer<T>> => {
  let resp: unknown
  try {
    resp = await response.json()
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
  } catch (e) {
    throw new HttpError(response, 'Failed to parse response: ' + response)
  }

  const respWrapper = baseResponse(respZodObject)
  const respObj = respWrapper.safeParse(resp)
  if (respObj.success) return respObj.data.data

  // if failed to parse response, we parse it as an error
  const errorResp = baseResponseNull.safeParse(resp)
  if (errorResp.success && errorResp.data.code !== 0) {
    throw new BusinessError(errorResp.data.code, errorResp.data.message)
  }

  throw new HttpError(response, 'Failed to parse response: ' + resp)
}

/**
 * Fetch with type validating, error handling, and base URL
 */
export const fetchWrapped = async <T extends z.ZodType>(
  input: RequestInfo,
  respZodObject: T,
  init?: RequestInit,
): Promise<z.infer<T>> => {
  const response = await fetch(input, init)
  return parseResponse(response, respZodObject)
}

export default fetchWrapped