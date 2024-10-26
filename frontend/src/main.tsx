import { StrictMode } from 'react'
import ReactDOM from 'react-dom/client'
import { createRouter, RouterProvider } from '@tanstack/react-router'

import { routeTree } from './routeTree.gen'

import './index.css'
import { QueryClientProvider } from '@tanstack/react-query'
import queryClient from '$/lib/query/client.ts'
import { ReactQueryDevtools } from '@tanstack/react-query-devtools'

async function enableMocking() {
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

enableMocking().then(() => {
  ReactDOM.createRoot(document.getElementById('root')!).render(
    <StrictMode>
      <QueryClientProvider client={queryClient}>
        <ReactQueryDevtools initialIsOpen={false} />

        <RouterProvider router={router} />
      </QueryClientProvider>
    </StrictMode>,
  )
})
