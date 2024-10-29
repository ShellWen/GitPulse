import { useQuery, useSuspenseQuery } from '@tanstack/react-query'
import { user, type User } from '$/types/user.ts'
import { BASE_URL, fetchWrapped } from '$/lib/query/index.ts'

const queryFn = async (username: string): Promise<User> => {
  return fetchWrapped(`${BASE_URL}/user/${username}`, user)
}

export const useUser = (username: string) => useQuery({
  queryKey: ['user', username],
  queryFn: () => queryFn(username),
})

export const useSuspenseUser = (username: string) => useSuspenseQuery({
  queryKey: ['user', username],
  queryFn: () => queryFn(username),
})
