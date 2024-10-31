import { BASE_URL } from '$/lib/api/constants.ts'
import fetchWrapped from '$/lib/api/fetchWrapped.ts'
import { z } from 'zod'

export const developer = z.object({
  id: z.number().nonnegative('Developer ID should be non-negative'),
  name: z.string().optional(),
  login: z.string().min(1, 'Login should not be empty'),
  avatarUrl: z.string().url('Avatar URL should be a valid URL'),
  company: z.string().nullable(),
  location: z.string().nullable(),
  bio: z.string().nullable(),
  blog: z.string().nullable(),
  email: z.string().email('Email should be a valid email').nullable(),

  followers: z.number().nonnegative('Followers count should be non-negative'),
  following: z.number().nonnegative('Following count should be non-negative'),

  stars: z.number().nonnegative('Stars count should be non-negative'),
  repos: z.number().nonnegative('Repos count should be non-negative'),
  gists: z.number().nonnegative('Gists count should be non-negative'),

  created_at: z.coerce.date(),
  updated_at: z.coerce.date(),
})

export type Developer = z.infer<typeof developer>

export const getDeveloper = async (username: string): Promise<Developer> => {
  return fetchWrapped(`${BASE_URL}/developer/${username}`, developer)
}
