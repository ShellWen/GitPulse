import { type Developer } from '$/types/developer.ts'
import { Avatar } from 'react-daisyui'

const UserInfo = ({ developer }: { developer: Developer }) => {
  return (
    <section className="flex w-full max-w-6xl flex-row flex-wrap gap-8">
      <Avatar src={developer.avatarUrl} shape="square" size="lg" />
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
