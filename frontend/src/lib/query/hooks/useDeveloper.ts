import { getDeveloper } from '$/lib/api/endpoint/developer.ts'
import { useQuery, useSuspenseQuery } from '@tanstack/react-query'

export const useDeveloper = (username: string) =>
  useQuery({
    queryKey: ['developer', username],
    queryFn: () => getDeveloper(username),
  })

export const useSuspenseDeveloper = (username: string) =>
  useSuspenseQuery({
    queryKey: ['developer', username],
    queryFn: () => getDeveloper(username),
  })
