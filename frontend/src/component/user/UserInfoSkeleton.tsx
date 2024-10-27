import { Skeleton } from 'react-daisyui'

const UserInfoSkeleton = () => {
  return (
    <section className="flex w-full max-w-6xl flex-row flex-wrap gap-8">
      <Skeleton className="h-32 w-32" />
      <div className="flex flex-1 flex-col gap-2">
        <Skeleton className="h-8 w-3/4" />
        <Skeleton className="h-6 w-1/4" />
        <div className="flex-1" />
        <Skeleton className="h-4 w-1/2" />
      </div>
    </section>
  )
}

export default UserInfoSkeleton
