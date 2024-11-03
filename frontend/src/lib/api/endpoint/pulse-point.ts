import { BASE_URL } from '$/lib/api/constants.ts'
import fetchWrapped from '$/lib/api/fetchWrapped.ts'
import { z } from 'zod'

export const pulsePoint = z.object({
  pulse_point: z.number().nonnegative(),

  created_at: z.coerce.date(),
  updated_at: z.coerce.date(),
})

export type PulsePoint = z.infer<typeof pulsePoint>

export const getPulsePoint = async (username: string): Promise<PulsePoint> => {
  return fetchWrapped(`${BASE_URL}/pulse-point/${username}`, pulsePoint)
}
