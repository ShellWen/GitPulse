import {
  getDeveloper,
  getDeveloperLanguages,
  getDeveloperPulsePoint,
  getDeveloperRegion,
  searchDevelopers,
} from '$/lib/api/endpoint/developer.ts'
import useSWR from 'swr'

export const useDeveloper = (username: string) =>
  useSWR(['developers', username], ([, username]) => getDeveloper(username))

export const useSuspenseDeveloper = (username: string) =>
  useSWR(['developers', username], ([, username]) => getDeveloper(username), { suspense: true })

export const useDeveloperPulsePoint = (username: string) =>
  useSWR(['developers', username, 'pulse-point'], ([, username]) => getDeveloperPulsePoint(username))

export const useSuspenseDeveloperPulsePoint = (username: string) =>
  useSWR(['developers', username, 'pulse-point'], ([, username]) => getDeveloperPulsePoint(username), {
    suspense: true,
  })

export const useDeveloperLanguages = (username: string) =>
  useSWR(['developers', username, 'languages'], ([, username]) => getDeveloperLanguages(username))

export const useSuspenseDeveloperLanguages = (username: string) =>
  useSWR(['developers', username, 'languages'], ([, username]) => getDeveloperLanguages(username), {
    suspense: true,
  })

export const useDeveloperRegion = (username: string) =>
  useSWR(['developers', username, 'region'], ([, username]) => getDeveloperRegion(username))

export const useSuspenseDeveloperRegion = (username: string) =>
  useSWR(['developers', username, 'region'], ([, username]) => getDeveloperRegion(username), { suspense: true })

export const useSearchDevelopers = (limit: number, languageId?: string, region?: string) =>
  useSWR(['developers', limit, languageId, region], ([, limit, languageId, region]) =>
    searchDevelopers(limit, languageId, region),
  )

export const useSuspenseSearchDevelopers = (limit: number, languageId?: string, region?: string) =>
  useSWR(
    ['developers', limit, languageId, region],
    ([, limit, languageId, region]) => searchDevelopers(limit, languageId, region),
    {
      suspense: true,
    },
  )
