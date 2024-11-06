import type { ReactNode } from 'react'

import { type Developer } from '$/lib/api/endpoint/types.ts'
import { Avatar } from 'react-daisyui'

const UserInfo = ({ developer, rightBlock }: { developer: Developer; rightBlock?: ReactNode }) => {
  return (
    <section className="relative flex w-full max-w-6xl flex-col flex-wrap gap-8 sm:flex-row">
      {rightBlock && <div className="absolute bottom-0 right-0 -z-10">{rightBlock}</div>}
      <Avatar src={developer.avatar_url} shape="square" size="lg" />
      <div className="flex flex-col gap-2">
        <div className="text-4xl font-bold leading-none">{developer.name}</div>
        <div className="text-xl font-light leading-none">{developer.login}</div>
        <div className="flex-1" />
        <div>{developer.bio}</div>
      </div>
    </section>
  )
}

export default UserInfo
