import { BASE_URL } from '$/lib/api/constants.ts'
import {
  type Developer,
  type DeveloperLanguages,
  type DeveloperPulsePoint,
  type DeveloperRegion,
  type DeveloperWithPulsePoint,
  type Language,
  developer,
  developerLanguages,
  developerPulsePoint,
  developerRegion,
  developerWithPulsePoint,
} from '$/lib/api/endpoint/types.ts'
import fetchWrapped from '$/lib/api/fetchWrapped.ts'

export const getDeveloper = async (username: string): Promise<Developer> => {
  return fetchWrapped(`${BASE_URL}/developers/${username}`, developer)
}

export const getDeveloperPulsePoint = async (username: string): Promise<DeveloperPulsePoint> => {
  return fetchWrapped(`${BASE_URL}/developers/${username}/pulse-point`, developerPulsePoint)
}

export const getDeveloperLanguages = async (username: string): Promise<DeveloperLanguages> => {
  return fetchWrapped(`${BASE_URL}/developers/${username}/languages`, developerLanguages)
}

export const getDeveloperRegion = async (username: string): Promise<DeveloperRegion> => {
  return fetchWrapped(`${BASE_URL}/developers/${username}/region`, developerRegion)
}

export const searchDevelopers = async (
  limit: number,
  languageId?: Language['id'],
  region?: DeveloperRegion['region']['region'],
): Promise<Array<DeveloperWithPulsePoint>> => {
  const url = new URL(`${BASE_URL}/developers/`)
  url.searchParams.set('limit', (limit ?? "").toString())
  url.searchParams.set('language', languageId ?? "")
  url.searchParams.set('region', region ?? "")
  return fetchWrapped(url, developerWithPulsePoint.array(), {
    method: 'GET',
  })
}
