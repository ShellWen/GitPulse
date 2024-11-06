import { z } from 'zod'

export const errorResponse = z.object({
  code: z.number(),
  message: z.string(),
})

export type ErrorResponse = z.infer<typeof errorResponse>
