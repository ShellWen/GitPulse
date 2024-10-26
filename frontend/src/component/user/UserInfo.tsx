import { User } from '$/types/user.ts'
import { Avatar } from 'react-daisyui'

const UserInfo = ({ user }: { user: User }) => {
  return (
    <section className="flex w-full max-w-4xl flex-row flex-wrap gap-8">
      <Avatar src={user.avatarUrl} shape="square" size="lg" />
      <div className="flex flex-col gap-2">
        <div className="text-4xl font-bold leading-none">{user.name}</div>
        <div className="text-xl font-light leading-none">{user.login}</div>
        <div className="flex-1" />
        <div>{user.bio}</div>
      </div>
    </section>
  )
}

export default UserInfo
