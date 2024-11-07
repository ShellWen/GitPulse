import { type FormEvent, Suspense, useCallback, useMemo, useState } from 'react'

import QueryErrorBoundaryBlock from '$/component/QueryErrorBoundaryBlock.tsx'
import { useSuspenseSearchDevelopers } from '$/lib/query/hooks/useDeveloper.ts'
import { useLanguages } from '$/lib/query/hooks/useLanguage.ts'
import { Link, createLazyFileRoute, getRouteApi } from '@tanstack/react-router'
import { getCountryDataList } from 'countries-list'
import { Button, Form, Input, Select } from 'react-daisyui'

const route = getRouteApi('/rank')

const RankResultWrapper = ({ language, region, limit }: { language?: string; region?: string; limit: number }) => {
  const { data } = useSuspenseSearchDevelopers(limit, language, region)

  return (
    <>
      {data?.map(({ developer, pulse_point }, index) => (
        <div key={developer.id} className="flex gap-4 items-center">
          <div className="rounded bg-base-300 w-12 h-12 flex justify-center items-center">{index + 1}</div>
          <div className="flex-shrink-0">
            <Link to={`/u/${developer.login}`}>
              <img src={developer.avatar_url} alt={developer.name} className="h-16 w-16 cursor-pointer rounded-full" />
            </Link>
          </div>
          <div>
            <div className="cursor-pointer text-lg">
              <Link to={`/u/${developer.login}`}>
                <span className="font-bold">{developer.name}</span>
                <span className="pl-2">{`@${developer.login}`}</span>
              </Link>
            </div>
            <div className="text-sm">
              <Link to={`/u/${developer.login}`}>{developer.bio}</Link>
            </div>
          </div>
          <div className="flex-1" />
          <div className="text-4xl">
            <Link to={`/u/${developer.login}`}>{`${pulse_point.pulse_point.toFixed(2)}pp`}</Link>
          </div>
        </div>
      ))}
    </>
  )
}

const RankPage = () => {
  const searchParams = route.useSearch()
  const navigate = route.useNavigate()

  const { data: languages } = useLanguages()
  const sortedLanguages = useMemo(() => languages?.sort((a, b) => a.name.localeCompare(b.name)), [languages])

  const regions: Array<[string, string]> = useMemo(() => {
    const l = getCountryDataList()
    const r: Array<[string, string]> = []
    for (const data of l) {
      const iso2LowerCase = data.iso2.toLowerCase()
      if (iso2LowerCase === 'cn' || iso2LowerCase === 'hk' || iso2LowerCase === 'mo' || iso2LowerCase === 'tw') {
        continue
      }
      r.push([iso2LowerCase, data.name])
    }
    r.sort((a, b) => a[1].localeCompare(b[1]))
    r.unshift(['cn', '中国大陆'], ['hk', '中国香港'], ['mo', '中国澳门'], ['tw', '中国台湾'])

    return r
  }, [])

  const [searchLanguage, setSearchLanguage] = useState(searchParams.language)
  const [searchRegion, setSearchRegion] = useState(searchParams.region)
  const [searchLimit, setSearchLimit] = useState(searchParams.limit)

  const onFormSubmit = useCallback(
    (e: FormEvent<HTMLFormElement>) => {
      e.preventDefault()
      navigate({
        search: {
          language: searchLanguage === '_all' ? undefined : searchLanguage,
          region: searchRegion === '_all' ? undefined : searchRegion,
          limit: searchLimit,
        },
      })
    },
    [navigate, searchLanguage, searchLimit, searchRegion],
  )

  return (
    <div className="mx-auto flex w-full max-w-6xl flex-col items-center gap-8 px-4 pt-8 md:flex-row md:items-start">
      <section className="w-full md:max-w-xs">
        {sortedLanguages && (
          <Form className="w-full gap-2" onSubmit={onFormSubmit}>
            <label className="label">
              <span className="label-text">编程语言</span>
            </label>
            <Select
              name="language"
              defaultValue="_all"
              value={searchLanguage}
              onChange={(e) => {
                setSearchLanguage(e.target.value)
              }}
            >
              <option value="_all">所有</option>
              {sortedLanguages.map((language) => (
                // backend use language name as the key, instead of language id
                <option key={language.id} value={language.name}>
                  {language.name}
                </option>
              ))}
            </Select>

            <label className="label">
              <span className="label-text">所在地区</span>
            </label>
            <Select
              name="region"
              defaultValue="_all"
              value={searchRegion}
              onChange={(e) => {
                setSearchRegion(e.target.value)
              }}
            >
              <option value="_all">所有</option>
              {regions.map(([iso2, name]) => (
                <option key={iso2} value={iso2}>
                  {name}
                </option>
              ))}
            </Select>

            <label className="label">
              <span className="label-text">请求数量</span>
            </label>
            <Input
              value={searchLimit}
              onChange={(e) => setSearchLimit(Number(e.target.value))}
              type="number"
              name="limit"
            />

            <Button type="submit" className="mt-2">
              查询
            </Button>
          </Form>
        )}
      </section>
      <QueryErrorBoundaryBlock>
        <Suspense fallback={<section className="skeleton h-96 w-full rounded bg-base-200 md:flex-1" />}>
          <section className="flex min-h-screen w-full flex-col gap-6 rounded bg-base-200 p-8 md:flex-1">
            <RankResultWrapper
              language={searchLanguage === '_all' ? undefined : searchLanguage}
              region={searchRegion === '_all' ? undefined : searchRegion}
              limit={searchLimit}
            />
          </section>
        </Suspense>
      </QueryErrorBoundaryBlock>
    </div>
  )
}

export const Route = createLazyFileRoute('/rank')({
  component: RankPage,
})
