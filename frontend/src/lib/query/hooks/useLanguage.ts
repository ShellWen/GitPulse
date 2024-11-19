import { getLanguages } from '$/lib/api/endpoint/languages.ts'
import useSWR from 'swr'

export const useLanguages = () =>
  useSWR(['languages'], getLanguages)

export const useSuspenseLanguages = () =>
  useSWR(['languages'], getLanguages, { suspense: true })
