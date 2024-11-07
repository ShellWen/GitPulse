import {
  getDeveloper,
  getDeveloperLanguages,
  getDeveloperPulsePoint,
  getDeveloperRegion, searchDevelopers,
} from '$/lib/api/endpoint/developer.ts'
import { useQuery, useSuspenseQuery } from '@tanstack/react-query'

export const useDeveloper = (username: string) =>
  useQuery({
    queryKey: ['developers', username],
    queryFn: () => getDeveloper(username),
  })

export const useSuspenseDeveloper = (username: string) =>
  useSuspenseQuery({
    queryKey: ['developers', username],
    queryFn: () => getDeveloper(username),
  })

export const useDeveloperPulsePoint = (username: string) =>
  useQuery({
    queryKey: ['developers', username, 'pulse-point'],
    queryFn: () => getDeveloperPulsePoint(username),
  })

export const useSuspenseDeveloperPulsePoint = (username: string) =>
  useSuspenseQuery({
    queryKey: ['developers', username, 'pulse-point'],
    queryFn: () => getDeveloperPulsePoint(username),
  })

export const useDeveloperLanguages = (username: string) =>
  useQuery({
    queryKey: ['developers', username, 'languages'],
    queryFn: () => getDeveloperLanguages(username),
  })

export const useSuspenseDeveloperLanguages = (username: string) =>
  useSuspenseQuery({
    queryKey: ['developers', username, 'languages'],
    queryFn: () => getDeveloperLanguages(username),
  })

export const useDeveloperRegion = (username: string) =>
  useQuery({
    queryKey: ['developers', username, 'region'],
    queryFn: () => getDeveloperRegion(username),
  })

export const useSuspenseDeveloperRegion = (username: string) =>
  useSuspenseQuery({
    queryKey: ['developers', username, 'region'],
    queryFn: () => getDeveloperRegion(username),
  })

export const useSearchDevelopers = (limit: number, languageId?: string, region?: string) =>
  useQuery({
    queryKey: ['developers', limit, languageId, region],
    queryFn: () => searchDevelopers(limit, languageId, region),
  })

export const useSuspenseSearchDevelopers = (limit: number, languageId?: string, region?: string) =>
  useSuspenseQuery({
    queryKey: ['developers', limit, languageId, region],
    queryFn: () => searchDevelopers(limit, languageId, region),
  })
