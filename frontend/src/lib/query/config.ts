import { HttpError } from '$/lib/api/error.ts'
import { QueryError } from '$/lib/query/error.ts'
import type { Middleware, SWRConfiguration } from 'swr'

const wrapError: Middleware = (useSWRNext) => (key, fetcher, config) => {
  const newFetcher = fetcher ? (...args: any[]) => {
    const result = fetcher(...args)
    if (result instanceof Promise) {
      return result.catch((error) => {
        throw new QueryError(error, key)
      })
    } else {
      return result
    }
  }: null
  return useSWRNext(key, newFetcher, config)
}

const swrConfig: SWRConfiguration = {
  errorRetryCount: 5,
  onErrorRetry: (error, _key, config, revalidate, { retryCount }) => {
    if (import.meta.env.VITE_ENABLE_MOCK) {
      return
    }

    const maxRetryCount = config.errorRetryCount ?? 5

    if (retryCount >= maxRetryCount) {
      return
    }

    if (error instanceof HttpError && !!error.response) {
      // don't retry on 4xx errors
      if (error.response.status >= 400 && error.response.status < 500) {
        return
      }
    }

    const delay = ~~((Math.random() + 0.5) * (1 << (retryCount < 8 ? retryCount : 8))) * config.errorRetryInterval
    setTimeout(() => revalidate({ retryCount }), delay)
  },

  use: [wrapError],
}

export default swrConfig
