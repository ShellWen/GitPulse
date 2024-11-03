import { type ComponentProps, type PropsWithChildren, Suspense, lazy, useCallback, useMemo } from 'react'

import DeveloperGlance from '$/component/developer/DeveloperGlance.tsx'
import DeveloperInfo from '$/component/developer/DeveloperInfo.tsx'
import DeveloperInfoSkeleton from '$/component/developer/DeveloperInfoSkeleton.tsx'
import { BusinessError } from '$/lib/api/error.ts'
import { useSuspenseDeveloper } from '$/lib/query/hooks/useDeveloper.ts'
import { usePulsePoint } from '$/lib/query/hooks/usePulsePoint.ts'
import useDarkMode from '$/lib/useDarkMode.ts'
import type { PieConfig } from '@ant-design/plots/es/components/pie'
import { QueryErrorResetBoundary } from '@tanstack/react-query'
import { createLazyFileRoute, getRouteApi } from '@tanstack/react-router'
import { Button, Skeleton } from 'react-daisyui'
import { ErrorBoundary, type FallbackProps } from 'react-error-boundary'

const route = getRouteApi('/u_/$username')

const DeveloperInfoWrapper = ({ username }: { username: string }) => {
  const { data: user } = useSuspenseDeveloper(username)
  const { data: pulsePoint } = usePulsePoint(username)

  return (
    <DeveloperInfo
      developer={user}
      rightBlock={
        <div className="text-clip whitespace-nowrap text-8xl font-bold italic tracking-wider text-base-content/20 sm:text-9xl">
          {pulsePoint ? `${pulsePoint.pulse_point}pp` : '计算中...'}
        </div>
      }
    />
  )
}

const DeveloperGlanceWrapper = ({ username }: { username: string }) => {
  const { data: user } = useSuspenseDeveloper(username)

  return <DeveloperGlance developer={user} />
}

interface DataItem {
  type: string
  value: number
}

const data = [
  { type: '分类一', value: 27 },
  { type: '分类二', value: 25 },
  { type: '分类三', value: 18 },
  { type: '分类四', value: 15 },
  { type: '分类五', value: 10 },
  { type: '其他', value: 5 },
] as DataItem[]

// It's too large to bundle the whole antd
const Pie = lazy(() => import('@ant-design/plots/es/components/pie'))

const DemoPie = () => {
  const isDarkMode = useDarkMode()
  const config: PieConfig = useMemo(
    () =>
      ({
        data,
        angleField: 'value',
        colorField: 'type',
        radius: 0.6,
        label: {
          text: (d: DataItem) => `${d.type}\n ${d.value}`,
          position: 'spider',
        },
        tooltip: {
          title: 'type',
          items: [
            {
              name: '值',
              field: 'value',
            },
          ],
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

const DeveloperTable = () => {
  return (
    <div className="flex w-full max-w-6xl flex-col">
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
const DeveloperNotFoundErrorBoundary = ({ children }: PropsWithChildren) => {
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

const DeveloperInfoErrorBoundary = ({ children }: PropsWithChildren) => {
  const renderer = useCallback(({ resetErrorBoundary, error }: FallbackProps) => {
    if (error instanceof BusinessError && error.code === 404) {
      // When the developer not found, the DeveloperNotFoundErrorBoundary will handle it
      throw error
    }
    const errorMsg = error instanceof Error ? error.message : '未知错误'
    return (
      <div>
        请求开发者信息失败：{errorMsg}
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

const DeveloperPage = () => {
  const { username } = route.useParams()

  return (
    <section className="flex w-full flex-col items-center gap-8 px-4 pt-8">
      <DeveloperNotFoundErrorBoundary>
        <>
          <DeveloperInfoErrorBoundary>
            <Suspense fallback={<DeveloperInfoSkeleton />}>
              <DeveloperInfoWrapper username={username} />
              <DeveloperGlanceWrapper username={username} />
            </Suspense>
          </DeveloperInfoErrorBoundary>
          {/* TODO: Add ErrorBoundary */}
          <Suspense fallback={<>TODO</>}>
            <DeveloperTable />
          </Suspense>
        </>
      </DeveloperNotFoundErrorBoundary>
    </section>
  )
}

export const Route = createLazyFileRoute('/u_/$username')({
  component: DeveloperPage,
})
