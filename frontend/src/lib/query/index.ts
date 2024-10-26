import { baseResponse } from '$/types/base.ts'
import { z } from 'zod'

export const BASE_URL = import.meta.env.VITE_BACKEND_BASEURL
if (!BASE_URL) {
  throw new Error('VITE_BACKEND_BASEURL is not defined')
}

/**
 * Fetch with type validating, error handling, and base URL
 */
export const fetchWrapped = async <T extends z.ZodType>(input: RequestInfo, respZodObject: T, init?: RequestInit): Promise<z.infer<T>> => {
  const response = await fetch(input, init)
  if (!response.ok) {
    throw new Error(response.statusText)
  }
  const respObj = await response.json()
  const respWrapper = baseResponse(respZodObject)
  const resp = respWrapper.parse(respObj)
  if (resp.code !== 0) {
    throw new Error(resp.message)
  }
  return resp.data
}