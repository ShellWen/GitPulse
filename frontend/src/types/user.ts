import { z } from 'zod'

export const user = z.object({
  userId: z.number().nonnegative("User ID should be non-negative"),
  name: z.string().optional(),
  login: z.string().min(1, "Login should not be empty"),
  avatarUrl: z.string().url("Avatar URL should be a valid URL"),
  bio: z.string().nullable(),

  followers: z.number().nonnegative("Followers count should be non-negative"),
  following: z.number().nonnegative("Following count should be non-negative"),

  stars: z.number().nonnegative("Stars count should be non-negative"),

  repositories: z.number().nonnegative("Repositories count should be non-negative"),
  gists: z.number().nonnegative("Gists count should be non-negative"),

  created_at: z.coerce.date(),
  updated_at: z.coerce.date(),
})

export type User = z.infer<typeof user>
