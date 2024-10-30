import { StrictMode } from 'react'

import queryClient from '$/lib/query/client.ts'
import { QueryClientProvider } from '@tanstack/react-query'
import { ReactQueryDevtools } from '@tanstack/react-query-devtools'
import { RouterProvider, createRouter } from '@tanstack/react-router'
import ReactDOM from 'react-dom/client'

import './index.css'
import { routeTree } from './routeTree.gen'

async function enableMockingIfEnabled() {
  // Tree-shake
  if (import.meta.env.VITE_ENABLE_MOCK) {
    const { worker } = await import('./mocks/browser')

    // `worker.start()` returns a Promise that resolves
    // once the Service Worker is up and ready to intercept requests.
    return (await worker()).start()
  }
}

const router = createRouter({ routeTree })
declare module '@tanstack/react-router' {
  interface Register {
    router: typeof router
  }
}

enableMockingIfEnabled().then(() => {
  ReactDOM.createRoot(document.getElementById('root')!).render(
    <StrictMode>
      <QueryClientProvider client={queryClient}>
        <ReactQueryDevtools initialIsOpen={false} />

        <RouterProvider router={router} />
      </QueryClientProvider>
    </StrictMode>,
  )
})
