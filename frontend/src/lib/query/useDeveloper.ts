import { useQuery, useSuspenseQuery } from '@tanstack/react-query'
import { developer, type Developer } from '$/types/developer.ts'
import { BASE_URL, fetchWrapped } from '$/lib/query/index.ts'

const queryFn = async (username: string): Promise<Developer> => {
  return fetchWrapped(`${BASE_URL}/developer/${username}`, developer)
}

export const useDeveloper = (username: string) => useQuery({
  queryKey: ['developer', username],
  queryFn: () => queryFn(username),
})

export const useSuspenseDeveloper = (username: string) => useSuspenseQuery({
  queryKey: ['developer', username],
  queryFn: () => queryFn(username),
})
