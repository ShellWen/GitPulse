import { Fragment, useMemo } from 'react'

import { User } from '$/types/user.ts'
import { Divider } from 'react-daisyui'

interface UserGlanceItem {
  name: string
  value: string
}

const UserGlanceBlock = ({ item }: { item: UserGlanceItem }) => {
  return (
    <div className="flex flex-1 flex-col items-center gap-2">
      <div className="text-2xl font-bold leading-none">{item.name}</div>
      <div className="text-lg font-light leading-none">{item.value}</div>
    </div>
  )
}

const UserGlance = ({ user }: { user: User }) => {
  const items: UserGlanceItem[] = useMemo(
    () => [
      { name: 'Repositories', value: user.repositories.toString() },
      { name: 'Stars', value: user.stars.toString() },
      { name: 'Followers', value: user.followers.toString() },
      { name: 'Following', value: user.following.toString() },
      { name: 'Gists', value: user.gists.toString() },
    ],
    [user],
  )

  return (
    <section className="flex w-full max-w-4xl flex-col gap-2 md:flex-row md:gap-0">
      <Divider horizontal />
      {items.map((item, index) => (
        <Fragment key={item.name + item.value}>
          <UserGlanceBlock item={item} />
          {index < items.length - 1 && <Divider horizontal />}
        </Fragment>
      ))}
      <Divider horizontal />
    </section>
  )
}

export default UserGlance
