import {
  type Developer,
  type DeveloperLanguages,
  type DeveloperRegion,
  type DeveloperWithPulsePoint,
  type Language,
  type PulsePoint,
  developer,
  developerLanguages,
  developerRegion,
  developerWithPulsePoint,
  pulsePoint,
} from '$/lib/api/endpoint/types.ts'
import { buildUrl, typedFetch, typedFetchAsync } from '$/lib/api/fetcher.ts'
import type { Key } from 'swr'
import { type SWRSubscription, type SWRSubscriptionOptions } from 'swr/subscription'

export const getDeveloper = async (username: string): Promise<Developer> => {
  return typedFetch(`/developers/${username}`, developer)
}

export const subscriptDeveloperPulsePoint =
  (username: string) =>
  (_: Key, { next }: SWRSubscriptionOptions<PulsePoint, Error>) => {
    const abortController = new AbortController()
    typedFetchAsync(`/developers/${username}/pulse-point`, pulsePoint, abortController, next)
    return () => {
      abortController.abort()
    }
  }

export const subscriptDeveloperLanguages =
  (username: string) =>
  (_: Key, { next }: SWRSubscriptionOptions<DeveloperLanguages, Error>): SWRSubscription => {
    const abortController = new AbortController()
    typedFetchAsync(`/developers/${username}/languages`, developerLanguages, abortController, next)
    return () => {
      abortController.abort()
    }
  }

export const subscribeDeveloperRegion =
  (username: string) =>
  (_: Key, { next }: SWRSubscriptionOptions<DeveloperRegion, Error>) => {
    const abortController = new AbortController()
    typedFetchAsync(`/developers/${username}/region`, developerRegion, abortController, next)
    return () => {
      abortController.abort()
    }
  }

export const searchDevelopers = async (
  limit: number,
  language?: Language['id'],
  region?: DeveloperRegion['region'],
): Promise<Array<DeveloperWithPulsePoint>> => {
  const url = buildUrl('/rank', {
    limit: (limit ?? '').toString(),
    language: language ?? '',
    region: region ?? '',
  })
  return typedFetch(url, developerWithPulsePoint.array(), {
    method: 'GET',
  })
}
