import { baseResponse } from '$/types/base.ts'
import { user } from '$/types/user.ts'
import { HttpHandler, HttpResponse, delay, http } from 'msw'
import { z } from 'zod'

import { BASE_URL } from './index.ts'

// noinspection JSUnusedGlobalSymbols
export const handlers = [
  http.get(`${BASE_URL}/user/:userName`, async () => {
    await delay(1000)

    const userResponse = baseResponse(user)
    type UserResponse = z.infer<typeof userResponse>

    const resp = userResponse.parse({
      code: 0,
      message: '',
      data: {
        userId: 38996248,
        name: 'ShellWen | 颉文',
        login: 'ShellWen',
        avatarUrl: 'https://avatars.githubusercontent.com/u/38996248?v=4',
        bio: 'Another Furry/🌈/Coder/Student',

        followers: 114514,
        following: 1919810,

        stars: 2333,

        repositories: 233,
        gists: 233,

        created_at: new Date('2018-05-05T02:44:13Z'),
        updated_at: new Date('2024-10-24T01:14:19Z'),
      },
    } satisfies UserResponse)
    return HttpResponse.json<UserResponse>(resp)
  }),
] satisfies Array<HttpHandler>
