import { ComponentProps, PropsWithChildren, Suspense, lazy, useCallback, useMemo } from 'react'

import UserGlance from '$/component/user/UserGlance.tsx'
import UserInfo from '$/component/user/UserInfo.tsx'
import UserInfoSkeleton from '$/component/user/UserInfoSkeleton.tsx'
import { BusinessError } from '$/lib/query/error.ts'
import { useSuspenseUser } from '$/lib/query/useUser.ts'
import useDarkMode from '$/lib/useDarkMode.ts'
import { QueryErrorResetBoundary } from '@tanstack/react-query'
import { createLazyFileRoute, getRouteApi } from '@tanstack/react-router'
import { Button, Skeleton } from 'react-daisyui'
import { ErrorBoundary, FallbackProps } from 'react-error-boundary'

const route = getRouteApi('/u_/$userName')

const UserInfoWrapper = ({ userName }: { userName: string }) => {
  const { data: user } = useSuspenseUser(userName)

  return <UserInfo user={user} />
}

const UserGlanceWrapper = ({ userName }: { userName: string }) => {
  const { data: user } = useSuspenseUser(userName)

  return <UserGlance user={user} />
}

const data = [
  { type: '分类一', value: 27 },
  { type: '分类二', value: 25 },
  { type: '分类三', value: 18 },
  { type: '分类四', value: 15 },
  { type: '分类五', value: 10 },
  { type: '其他', value: 5 },
]

// It's too large to bundle the whole antd
const Pie = lazy(() => import('@ant-design/plots/es/components/pie'))

const DemoPie = () => {
  const isDarkMode = useDarkMode()
  const config = useMemo(
    () =>
      ({
        data,
        angleField: 'value',
        colorField: 'type',
        radius: 0.8,
        label: {
          text: (d) => `${d.type}\n ${d.value}`,
          position: 'spider',
        },
        legend: {
          color: {
            title: false,
            position: 'right',
            rowPadding: 5,
          },
        },
        theme: isDarkMode ? 'classicDark' : 'classic',
      }) satisfies ComponentProps<typeof Pie>,
    [isDarkMode],
  )
  return <Pie {...config} />
}

const UserTable = () => {
  return (
    <div className="flex w-full max-w-4xl flex-col">
      <div className="grid grid-cols-1 gap-4 lg:grid-cols-2">
        <Suspense fallback={<Skeleton className="h-64 w-full rounded bg-base-200" />}>
          <section className="h-64 w-full rounded bg-base-200">
            <DemoPie />
          </section>
        </Suspense>
        <Skeleton className="h-64 w-full rounded bg-base-200" />
        <Skeleton className="h-64 w-full rounded bg-base-200" />
        <Skeleton className="h-64 w-full rounded bg-base-200" />
        <Skeleton className="h-64 w-full rounded bg-base-200" />
        <Skeleton className="h-64 w-full rounded bg-base-200" />
      </div>
    </div>
  )
}
const UserNotFoundErrorBoundary = ({ children }: PropsWithChildren) => {
  const renderer = useCallback(({ error }: FallbackProps) => {
    if (!(error instanceof BusinessError) || error.code !== 404) {
      // Only handle 404 error
      throw error
    }
    // TODO: Add a button to redirect to home page
    return <div>您请求的用户名不存在</div>
  }, [])
  return <ErrorBoundary fallbackRender={renderer}>{children}</ErrorBoundary>
}

const UserInfoErrorBoundary = ({ children }: PropsWithChildren) => {
  const renderer = useCallback(({ resetErrorBoundary, error }: FallbackProps) => {
    if (error instanceof BusinessError && error.code === 404) {
      // When user not found, the UserNotFoundErrorBoundary will handle it
      throw error
    }
    const errorMsg = error instanceof Error ? error.message : '未知错误'
    return (
      <div>
        请求用户信息失败：{errorMsg}
        <Button onClick={() => resetErrorBoundary()}>重试</Button>
      </div>
    )
  }, [])
  return (
    <QueryErrorResetBoundary>
      {({ reset }) => (
        <ErrorBoundary onReset={reset} fallbackRender={renderer}>
          {children}
        </ErrorBoundary>
      )}
    </QueryErrorResetBoundary>
  )
}

const UserPage = () => {
  const { userName } = route.useParams()

  return (
    <section className="flex w-full flex-col items-center gap-8 px-4 pt-8">
      <UserNotFoundErrorBoundary>
        <>
          <UserInfoErrorBoundary>
            <Suspense fallback={<UserInfoSkeleton />}>
              <UserInfoWrapper userName={userName} />
              <UserGlanceWrapper userName={userName} />
            </Suspense>
          </UserInfoErrorBoundary>
          {/* TODO: Add ErrorBoundary */}
          <Suspense fallback={<>TODO</>}>
            <UserTable />
          </Suspense>
        </>
      </UserNotFoundErrorBoundary>
    </section>
  )
}

export const Route = createLazyFileRoute('/u_/$userName')({
  component: UserPage,
})
