import { ComponentProps, Suspense, useMemo } from 'react'

import UserGlance from '$/component/user/UserGlance.tsx'
import UserInfo from '$/component/user/UserInfo.tsx'
import UserInfoSkeleton from '$/component/user/UserInfoSkeleton.tsx'
import { useSuspenseUser } from '$/lib/query/useUser.ts'
import useDarkMode from '$/lib/useDarkMode.ts'
import { Pie } from '@ant-design/plots'
import { createLazyFileRoute, getRouteApi } from '@tanstack/react-router'
import { Skeleton } from 'react-daisyui'

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
        <section className="h-64 w-full rounded bg-base-200">
          <DemoPie />
        </section>
        <Skeleton className="h-64 w-full rounded bg-base-200" />
        <Skeleton className="h-64 w-full rounded bg-base-200" />
        <Skeleton className="h-64 w-full rounded bg-base-200" />
        <Skeleton className="h-64 w-full rounded bg-base-200" />
        <Skeleton className="h-64 w-full rounded bg-base-200" />
      </div>
    </div>
  )
}

const UserPage = () => {
  const { userName } = route.useParams()
  return (
    <section className="flex w-full flex-col items-center gap-8 px-4 pt-8">
      <Suspense fallback={<UserInfoSkeleton />}>
        <UserInfoWrapper userName={userName} />
      </Suspense>
      <Suspense fallback={<></>}>
        <UserGlanceWrapper userName={userName} />
      </Suspense>
      <Suspense fallback={<>TODO</>}>
        <UserTable />
      </Suspense>
    </section>
  )
}

export const Route = createLazyFileRoute('/u_/$userName')({
  component: UserPage,
})
