import { getPulsePoint } from '$/lib/api/endpoint/pulse-point.ts'
import { useQuery, useSuspenseQuery } from '@tanstack/react-query'

export const usePulsePoint = (username: string) =>
  useQuery({
    queryKey: ['pulse-point', username],
    queryFn: () => getPulsePoint(username),
  })

export const useSuspensePulsePoint = (username: string) =>
  useSuspenseQuery({
    queryKey: ['pulse-point', username],
    queryFn: () => getPulsePoint(username),
  })
