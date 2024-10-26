import { QueryClient } from '@tanstack/react-query'

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      retry: !!import.meta.env.VITE_ENABLE_MOCK,
    },
  },
})

export default queryClient
