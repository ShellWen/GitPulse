import { QueryClient } from '@tanstack/react-query'
import { HttpError } from '$/lib/api/error.ts'

const MAX_RETRIES = 3

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      retry: (failureCount, error) => {
        if (import.meta.env.VITE_ENABLE_MOCK) {
          return false
        }

        if (failureCount > MAX_RETRIES) {
          return false
        }

        if (
          error instanceof HttpError && !!error.response
        ) {
          // don't retry on 4xx errors
          if (error.response.status >= 400 && error.response.status < 500) {
            return false
          }
        }

        return true
      },
    },
  },
})

export default queryClient
