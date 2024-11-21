import { BusinessError, HttpError } from '$/lib/api/error.ts'
import { errorResponse } from '$/lib/api/types.ts'
import { z } from 'zod'
import { BASE_URL, DEFAULT_TIMEOUT } from '$/lib/api/constants.ts'
import { task } from '$/lib/api/endpoint/types.ts'

const delay = (ms: number) => new Promise((resolve) => setTimeout(resolve, ms))

export const buildUrl = (path: string, params: Record<string, string | number | boolean> = {}): URL => {
  const url = new URL(BASE_URL + path)
  for (const [key, value] of Object.entries(params)) {
    url.searchParams.set(key, value.toString())
  }
  return url
}

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

const fetchWithBaseUrl: typeof fetch = async (input, init) => {
  const baseUrl = BASE_URL
  if (typeof input === 'string' && !input.startsWith('http')) {
    input = baseUrl + input
  } else if (input instanceof URL && !input.toString().startsWith('http')) {
    input = new URL(baseUrl + input.href)
  }

  return fetch(input, init)
}

type TimeoutRequestInit = RequestInit & {
  gp_timeout?: number
}

type FetchWithTimeout = (
  input: RequestInfo | URL,
  init?: TimeoutRequestInit,
) => Promise<Response>

const fetchWithTimeout: FetchWithTimeout = async (input, init) => {
  const { gp_timeout, ...restInit } = init ?? {}
  const controller = new AbortController()
  const timer = setTimeout(() => controller.abort(), gp_timeout ?? DEFAULT_TIMEOUT)
  const response = await fetchWithBaseUrl(input, { ...restInit, signal: controller.signal })
  clearTimeout(timer)
  return response
}

/**
 * Fetch with type validating, error handling, and base URL
 */
export const typedFetch = async <T extends z.ZodType>(
  input: RequestInfo | URL,
  respZodObject: T,
  init?: TimeoutRequestInit
): Promise<z.infer<T>> => {
  const response = await fetchWithTimeout(input, init)
  return parseResponse(response, respZodObject)
}

export const typedFetchAsync = async <T extends z.ZodType>(
  input: RequestInfo | URL,
  respZodObject: T,
  abortController: AbortController,
  next: (error: Error | null, data?: z.infer<T>) => void,
): Promise<void> => {
  console.info('typedFetchAsync', 'startFetch', input)
  const taskId = (await typedFetch(input, task, { method: 'POST' })).task_id
  console.info('typedFetchAsync', 'taskId', input, taskId)
  while (!abortController.signal.aborted) {
    try {
      const response = await fetchWithTimeout(buildUrl(input.toString(), {
        task_id: taskId,
      }))
      if (response.status === 202) {
        console.debug('typedFetchAsync', 'wait', taskId)
        await delay(1000)
        continue
      }
      const result = await parseResponse(response, respZodObject)
      next(null, result)
      console.log('typedFetchAsync', 'result', taskId, result)
      break
    } catch (e) {
      console.warn('typedFetchAsync', 'error', taskId, e)
      next(<Error>e)
      throw e
    }
  }
}
