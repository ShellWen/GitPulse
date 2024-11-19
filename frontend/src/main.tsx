import { StrictMode } from 'react'

import swrConfig from '$/lib/query/config.ts'
import { RouterProvider, createRouter } from '@tanstack/react-router'
import ReactDOM from 'react-dom/client'
import { SWRConfig } from 'swr'

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
      <SWRConfig value={swrConfig}>
        <RouterProvider router={router} />
      </SWRConfig>
    </StrictMode>,
  )
})
