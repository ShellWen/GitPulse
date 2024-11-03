import { pulsePoint } from '$/lib/api/endpoint/pulse-point.ts'
import { baseResponse } from '$/lib/api/types.ts'
import { HttpHandler, HttpResponse, delay, http } from 'msw'
import { z } from 'zod'

import { BASE_URL } from '../constants.ts'

// noinspection JSUnusedGlobalSymbols
export const handlers = [
  http.get(`${BASE_URL}/pulse-point/:username`, async ({ params }) => {
    await delay(4000)

    if (params.username !== 'shellwen') {
      return HttpResponse.json(
        baseResponse(z.null()).parse({
          code: 404,
          message: 'Developer not found',
          data: null,
        }),
        {
          status: 404,
        },
      )
    }

    const pulsePointResponse = baseResponse(pulsePoint)
    type PulsePointResponse = z.infer<typeof pulsePointResponse>

    const resp = pulsePointResponse.parse({
      code: 0,
      message: '',
      data: {
        pulse_point: 233,

        created_at: new Date('2024-10-24T11:45:14Z'),
        updated_at: new Date('2024-10-24T11:45:14Z'),
      },
    } satisfies PulsePointResponse)
    return HttpResponse.json<PulsePointResponse>(resp)
  }),
] satisfies Array<HttpHandler>
