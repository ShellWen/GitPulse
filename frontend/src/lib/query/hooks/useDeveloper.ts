import { useMemo } from 'react'

import {
  getDeveloper,
  searchDevelopers,
  subscribeDeveloperRegion, subscribeDeveloperSummary,
  subscriptDeveloperLanguages,
  subscriptDeveloperPulsePoint,
} from '$/lib/api/endpoint/developer.ts'
import useSWR from 'swr'
import useSWRSubscription from 'swr/subscription'

export const useDeveloper = (username: string) =>
  useSWR(['developers', username], ([, username]) => getDeveloper(username))

export const useSuspenseDeveloper = (username: string) =>
  useSWR(['developers', username], ([, username]) => getDeveloper(username), { suspense: true })

export const useDeveloperPulsePoint = (username: string) => {
  const fetcher = useMemo(() => subscriptDeveloperPulsePoint(username), [username])
  return useSWRSubscription(['developers', username, 'pulse-point'], fetcher)
}

export const useDeveloperLanguages = (username: string) => {
  const fetcher = useMemo(() => subscriptDeveloperLanguages(username), [username])
  return useSWRSubscription(['developers', username, 'languages'], fetcher)
}

export const useDeveloperRegion = (username: string) => {
  const fetcher = useMemo(() => subscribeDeveloperRegion(username), [username])
  return useSWRSubscription(['developers', username, 'region'], fetcher)
}

export const useDeveloperSummary = (username: string) => {
  const fetcher = useMemo(() => subscribeDeveloperSummary(username), [username])
  return useSWRSubscription(['developers', username, 'summary'], fetcher)
}

export const useSearchDevelopers = (limit: number, language?: string, region?: string) =>
  useSWR(['developers', limit, language, region], ([, limit, language, region]) =>
    searchDevelopers(limit, language, region),
  )

export const useSuspenseSearchDevelopers = (limit: number, language?: string, region?: string) =>
  useSWR(
    ['developers', limit, language, region],
    ([, limit, language, region]) => searchDevelopers(limit, language, region),
    {
      suspense: true,
    },
  )
