import { createLazyFileRoute, getRouteApi } from '@tanstack/react-router'
import { Button, Hero, Input } from 'react-daisyui'
import { type ChangeEvent, type FormEvent, useCallback, useState } from 'react'

const route = getRouteApi('/')

export const Route = createLazyFileRoute('/')({
  component: Index,
})

function Index() {
  const navigate = route.useNavigate()
  const [userName, setUserName] = useState('')
  const onInputChange = useCallback((e: ChangeEvent<HTMLInputElement>) => {
    setUserName(e.target.value)
  }, [setUserName])
  const onSubmit = useCallback((e: FormEvent) => {
    e.preventDefault()

    const processedUserName = userName.trim().toLowerCase()
    if (!processedUserName) {
      return
    }
    navigate({
      to: `/u/${processedUserName}`,
    })
  }, [navigate, userName])
  // const redirectToUserPage = () => {}
  return (
    <section className="w-full h-screen flex flex-col items-center justify-center">
      <Hero>
        <Hero.Content className="text-center">
          <div className="max-w-md">
            <h1 className="text-5xl font-bold">GitPulse</h1>
            <p className="py-6 text-center">
              GitPulse 利用 GitHub 开源项目数据，精准评估开发者的技术水平。通过分析项目影响力和贡献度，生成精准的
              TalentRank，并推测开发者的地域和专长，助您轻松发现全球优秀开发者与技术专家。
            </p>

            <form onSubmit={onSubmit} className="flex flex-col md:flex-row md:justify-center gap-4">
              <Input bordered placeholder="GitHub 用户名" onInput={onInputChange} />
              <Button>提交</Button>
            </form>
          </div>
        </Hero.Content>
      </Hero>
    </section>
  )
}