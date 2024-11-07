import { type ComponentProps, type PropsWithChildren, Suspense, lazy, useCallback, useMemo } from 'react'

import QueryErrorBoundaryBlock from '$/component/QueryErrorBoundaryBlock.tsx'
import DeveloperGlance from '$/component/developer/DeveloperGlance.tsx'
import DeveloperInfo from '$/component/developer/DeveloperInfo.tsx'
import DeveloperInfoSkeleton from '$/component/developer/DeveloperInfoSkeleton.tsx'
import type { DeveloperLanguages } from '$/lib/api/endpoint/types.ts'
import { HttpError } from '$/lib/api/error.ts'
import {
  useDeveloperPulsePoint,
  useSuspenseDeveloper,
  useSuspenseDeveloperLanguages,
  useSuspenseDeveloperRegion,
} from '$/lib/query/hooks/useDeveloper.ts'
import useDarkMode from '$/lib/useDarkMode.ts'
import type { PieConfig } from '@ant-design/plots/es/components/pie'
import { createLazyFileRoute, getRouteApi } from '@tanstack/react-router'
import { type TCountryCode, getEmojiFlag } from 'countries-list'
import { Skeleton } from 'react-daisyui'
import { ErrorBoundary, type FallbackProps } from 'react-error-boundary'

const route = getRouteApi('/u_/$username')

const DeveloperInfoWrapper = ({ username }: { username: string }) => {
  const { data: user } = useSuspenseDeveloper(username)
  const { data: pulsePoint } = useDeveloperPulsePoint(username)

  return (
    <DeveloperInfo
      developer={user}
      rightBlock={
        <div className="text-clip whitespace-nowrap text-8xl font-bold italic tracking-wider text-base-content/20 sm:text-9xl">
          {pulsePoint ? `${pulsePoint.pulse_point.pulse_point.toFixed(2)}pp` : '计算中...'}
        </div>
      }
    />
  )
}

const DeveloperGlanceWrapper = ({ username }: { username: string }) => {
  const { data: user } = useSuspenseDeveloper(username)

  return <DeveloperGlance developer={user} />
}

interface LanguagePieItem {
  id: string
  color?: string
  name: string
  percentage: number
}

// It's too large to bundle the whole antd
const Pie = lazy(() => import('@ant-design/plots/es/components/pie'))

const LanguagePie = ({ data }: { data: DeveloperLanguages }) => {
  const isDarkMode = useDarkMode()
  const flattenedData: LanguagePieItem[] = useMemo(
    () => [
      ...data.languages.languages.map((language) => ({
        id: language.language.id,
        color: language.language.color,
        name: language.language.name,
        percentage: language.percentage,
      })),
    ],
    [data],
  )

  // Sort the data by percentage
  const sortedData = useMemo(() => {
    const newData = [...flattenedData]
    newData.sort((a, b) => b.percentage - a.percentage)
    return newData
  }, [flattenedData])

  // Limit the data to 10 items
  const limitedData = useMemo(() => {
    if (sortedData.length <= 10) {
      return sortedData
    }
    const top10 = sortedData.slice(0, 10)
    const rest = sortedData.slice(10)
    const restPercentage = rest.reduce((prev, current) => prev + current.percentage, 0)
    return [
      ...top10,
      {
        id: 'rest',
        name: '其它',
        percentage: restPercentage,
      },
    ]
  }, [sortedData])

  const config: PieConfig = useMemo(
    () =>
      ({
        data: limitedData,
        angleField: 'percentage',
        colorField: 'name',
        radius: 0.75,
        label: {
          text: (d: LanguagePieItem) => `${d.name} ${d.percentage.toFixed(2)}%`,
          position: 'spider',
        },
        tooltip: {
          title: 'name',
          items: [
            {
              name: '百分比',
              field: 'percentage',
              valueFormatter: (v: number) => `${v.toFixed(2)}%`,
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
    [flattenedData, isDarkMode],
  )
  return <Pie {...config} />
}

const DeveloperLanguageBlock = ({ username }: { username: string }) => {
  const { data: developer } = useSuspenseDeveloper(username)
  const { data } = useSuspenseDeveloperLanguages(username)
  const mostUsedLanguage = useMemo(() => {
    if (!data) {
      return null
    }
    return data.languages.languages.reduce((prev, current) => (prev.percentage > current.percentage ? prev : current))
  }, [data])
  return (
    <>
      <section className="w-full rounded bg-base-200 p-8 lg:col-span-2 lg:h-96">
        <p>
          {`${developer.name} 使用最多的语言是 ${mostUsedLanguage?.language.name}，占比 ${mostUsedLanguage?.percentage.toFixed(2)}%。`}
          {/* TODO: Styles */}
        </p>
      </section>
      <section className="w-full rounded bg-base-200 lg:col-span-3 lg:h-96">
        <LanguagePie data={data} />
      </section>
    </>
  )
}

const RegionNotSure = Symbol('RegionNotSure')

const DeveloperRegionBlock = ({ username }: { username: string }) => {
  const { data: developer } = useSuspenseDeveloper(username)
  const { data } = useSuspenseDeveloperRegion(username)
  const { region, confidence } = useMemo(() => {
    if (!data) {
      return { region: RegionNotSure, confidence: 0 }
    }
    if (data.region.confidence < 0.1) {
      return { region: RegionNotSure, confidence: data.region.confidence }
    }
    return { region: data.region.region, confidence: data.region.confidence }
  }, [data])
  const regionName = useMemo(() => {
    if (region === RegionNotSure) {
      return '未知'
    }

    if (region === 'cn') {
      return '中国大陆'
    } else if (region === 'hk') {
      return '中国香港'
    } else if (region === 'mo') {
      return '中国澳门'
    } else if (region === 'tw') {
      return '中国台湾'
    } else {
      return ` ${(region as string).toUpperCase()} `
    }
  }, [region])
  const regionFlag = useMemo(() => {
    if (region === RegionNotSure) {
      return '❓'
    }
    if (region === 'cn') {
      return '🇨🇳'
    } else if (region === 'hk') {
      return '🇭🇰'
    } else if (region === 'mo') {
      return '🇲🇴'
    } else if (region === 'tw') {
      return '🇨🇳'
    } else {
      return getEmojiFlag((region as string).toUpperCase() as TCountryCode)
    }
  }, [region])

  return (
    <>
      <section className="flex w-full flex-col items-center justify-center rounded bg-base-200 p-8 text-8xl lg:col-span-3 lg:h-96">
        {regionFlag}
      </section>
      <section className="w-full rounded bg-base-200 p-8 lg:col-span-2 lg:h-96">
        {region === RegionNotSure ? (
          <p>我们没有找到开发者可能所处的地区。</p>
        ) : (
          <p>{`${developer.name} 来自 ${regionName}，置信度为 ${(confidence * 100).toFixed(2)}%。`}</p>
        )}
      </section>
    </>
  )
}

const DeveloperBlockSuspense = ({ children }: PropsWithChildren) => {
  return (
    <Suspense fallback={<Skeleton className="h-64 w-full rounded bg-base-200 lg:col-span-5" />}>{children}</Suspense>
  )
}

const DeveloperTable = ({ username }: { username: string }) => {
  return (
    <div className="flex w-full max-w-6xl flex-col">
      <div className="grid grid-cols-1 gap-4 lg:grid-cols-5">
        <DeveloperBlockSuspense>
          <DeveloperLanguageBlock username={username} />
        </DeveloperBlockSuspense>
        <DeveloperBlockSuspense>
          <DeveloperRegionBlock username={username} />
        </DeveloperBlockSuspense>
      </div>
    </div>
  )
}
const DeveloperNotFoundErrorBoundary = ({ children }: PropsWithChildren) => {
  const renderer = useCallback(({ error }: FallbackProps) => {
    if (!(error instanceof HttpError) || error.response.status !== 404) {
      // Only handle 404 error
      throw error
    }
    // TODO: Add a button to redirect to home page
    return <div>您请求的用户名不存在</div>
  }, [])
  return <ErrorBoundary fallbackRender={renderer}>{children}</ErrorBoundary>
}

const DeveloperPage = () => {
  const { username } = route.useParams()

  return (
    <div className="flex w-full flex-col items-center gap-8 px-4 pt-8">
      <DeveloperNotFoundErrorBoundary>
        <>
          <QueryErrorBoundaryBlock>
            <Suspense fallback={<DeveloperInfoSkeleton />}>
              <DeveloperInfoWrapper username={username} />
              <DeveloperGlanceWrapper username={username} />
            </Suspense>
          </QueryErrorBoundaryBlock>
          <QueryErrorBoundaryBlock>
            <DeveloperTable username={username} />
          </QueryErrorBoundaryBlock>
        </>
      </DeveloperNotFoundErrorBoundary>
    </div>
  )
}

export const Route = createLazyFileRoute('/u_/$username')({
  component: DeveloperPage,
})
