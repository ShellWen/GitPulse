import { getLanguages } from '$/lib/api/endpoint/languages.ts'
import { useQuery, useSuspenseQuery } from '@tanstack/react-query'

export const useLanguages = () =>
  useQuery({
    queryKey: ['languages'],
    queryFn: () => getLanguages(),
  })

export const useSuspenseLanguages = () =>
  useSuspenseQuery({
    queryKey: ['languages'],
    queryFn: () => getLanguages(),
  })
