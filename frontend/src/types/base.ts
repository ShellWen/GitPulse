import { z } from 'zod'

export const baseRequest = z.object({})

export type BaseRequest = z.infer<typeof baseRequest>

export const baseResponse = <T extends z.ZodType>(data: T) => z.object({
  code: z.number(),
  message: z.string(),
  data: data,
})

export type BaseResponse<T extends z.ZodType> = z.infer<ReturnType<typeof baseResponse<T>>>
