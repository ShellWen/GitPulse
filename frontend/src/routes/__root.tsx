import { Suspense, lazy, useEffect } from 'react'

import AppBar from '$/component/AppBar.tsx'
import useDarkMode from '$/lib/useDarkMode.ts'
import { Outlet, createRootRoute } from '@tanstack/react-router'
import { useTheme } from 'react-daisyui'
import AppFooter from '$/component/AppFooter.tsx'

const TanStackRouterDevtools = import.meta.env.PROD
  ? () => null // Render nothing in production
  : lazy(() =>
      // Lazy load in development
      import('@tanstack/router-devtools').then((res) => ({
        default: res.TanStackRouterDevtools,
        // For Embedded Mode
        // default: res.TanStackRouterDevtoolsPanel
      })),
    )

const Root = () => {
  const { setTheme } = useTheme()

  const isDarkMode = useDarkMode()

  useEffect(() => {
    setTheme(isDarkMode ? 'dark' : 'light')
  }, [isDarkMode, setTheme])

  return (
    <>
      <div className="w-full max-w-[100vw] overflow-clip">
        <AppBar />
        <main className="min-h-screen w-full">
          <Outlet />
        </main>
        <AppFooter />
      </div>
      <Suspense>
        <TanStackRouterDevtools />
      </Suspense>
    </>
  )
}

export const Route = createRootRoute({
  component: Root,
})
