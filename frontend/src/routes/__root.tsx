import { Suspense, lazy, useEffect } from 'react'

import AppBar from '$/component/AppBar.tsx'
import AppFooter from '$/component/AppFooter.tsx'
import useDarkMode from '$/lib/useDarkMode.ts'
import { Outlet, createRootRoute } from '@tanstack/react-router'
import { useTheme } from 'react-daisyui'

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
      <div className="flex min-h-screen w-full max-w-[100vw] flex-col overflow-clip">
        <AppBar />
        <main className="flex flex-1 flex-col justify-center">
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
