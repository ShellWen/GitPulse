import {
  type Developer,
  type DeveloperLanguages,
  type DeveloperRegion,
  type DeveloperWithPulsePoint,
  type Language,
  type PulsePoint,
  type Task,
  developer,
  developerLanguages,
  developerRegion,
  pulsePoint,
  task,
} from '$/lib/api/endpoint/types.ts'
import { type ErrorResponse, errorResponse } from '$/lib/api/types.ts'
import { getTask, newTask } from '$/mocks/tasks.ts'
import { HttpHandler, HttpResponse, delay, http } from 'msw'
import type { HttpRequestHandler } from 'msw/core/http'

import { BASE_URL, DEFAULT_ASYNC_DELAY } from '../constants.ts'
import languagesJson from './languages.json'

const languages: Array<Language> = languagesJson

const fakeDeveloper: Developer = {
  id: 38996248,
  name: 'ShellWen | 颉文',
  login: 'ShellWen',
  avatar_url: 'https://avatars.githubusercontent.com/u/38996248?v=4',
  company: 'ShellWen Company',
  location: 'Utopia',
  bio: 'Another Furry/🌈/Coder/Student',
  blog: 'https://shellwen.com',
  email: 'me@shellwen.com',

  followers: 114514,
  following: 1919810,

  stars: 2333,
  repos: 233,
  gists: 233,

  created_at: new Date('2018-05-05T02:44:13Z'),
  updated_at: new Date('2024-10-24T01:14:19Z'),
}

const developerNotFoundResponse = HttpResponse.json(
  errorResponse.parse({
    code: 404,
    message: 'Developer not found',
  } satisfies ErrorResponse),
  { status: 404 },
)

// they don't expose this type so...
type ResponseResolverInfo = Parameters<Parameters<HttpRequestHandler>[1]>[0]

const httpAsyncApi = <ProduceType extends Response>(
  route: string,
  producer: (req: ResponseResolverInfo) => Promise<ProduceType>,
): HttpHandler[] => {
  return [
    http.post(route, async (req) => {
      const [taskId, completeTask] = newTask()
      void producer(req).then(
        (resp) => {
          completeTask(resp)
        },
        (err) => {
          completeTask(
            HttpResponse.json(
              errorResponse.parse({
                code: -1,
                message: `Failed to produce response: ${err}`,
              } satisfies ErrorResponse),
              {
                status: 500,
              },
            ),
          )
        },
      )
      const resp = task.parse({
        task_id: taskId,
      } satisfies Task)
      return HttpResponse.json(resp)
    }),
    http.get(route, async ({ request }) => {
      const url = new URL(request.url)
      const task_id = url.searchParams.get('task_id')
      if (!task_id) {
        return HttpResponse.json(
          errorResponse.parse({
            code: 400,
            message: 'Task ID is required',
          } satisfies ErrorResponse),
          {
            status: 400,
          },
        )
      }
      const result = getTask(task_id)
      if (result === undefined) {
        return HttpResponse.json(
          errorResponse.parse({
            code: 404,
            message: 'Task not found',
          } satisfies ErrorResponse),
          {
            status: 404,
          },
        )
      }
      if (result === null) {
        return HttpResponse.json(
          // TODO: type safe
          {
            state: 'active',
            reason: '',
          },
          {
            status: 202,
          },
        )
      }
      return result
    }),
  ]
}

// noinspection JSUnusedGlobalSymbols
export const handlers = [
  http.get(`${BASE_URL}/developers/:username`, async ({ params }) => {
    await delay(1000)

    if (params.username !== fakeDeveloper.login.toLowerCase()) {
      return developerNotFoundResponse.clone()
    }

    const resp = developer.parse(fakeDeveloper)
    return HttpResponse.json<Developer>(resp)
  }),
  ...httpAsyncApi(`${BASE_URL}/developers/:username/pulse-point`, async ({ params }) => {
    if (params.username !== fakeDeveloper.login.toLowerCase()) {
      return HttpResponse.json(
        errorResponse.parse({
          code: 404,
          message: 'Developer not found',
        } satisfies ErrorResponse),
        {
          status: 404,
        },
      )
    }

    await delay(DEFAULT_ASYNC_DELAY)

    const resp = pulsePoint.parse({
      id: fakeDeveloper.id,
      pulse_point: 233,

      updated_at: new Date('2024-10-24T11:45:14Z'),
    } satisfies PulsePoint)
    return HttpResponse.json(resp)
  }),
  ...httpAsyncApi(`${BASE_URL}/developers/:username/languages`, async ({ params }) => {
    if (params.username !== fakeDeveloper.login.toLowerCase()) {
      return developerNotFoundResponse.clone()
    }

    await delay(DEFAULT_ASYNC_DELAY)

    const resp = developerLanguages.parse({
      id: fakeDeveloper.id,

      languages: [
        {
          language: languages.find((l) => l.id === 'typescript')!,
          percentage: 60.9,
        },
        {
          language: languages.find((l) => l.id === 'kotlin')!,
          percentage: 19.1,
        },
        {
          language: languages.find((l) => l.id === 'rust')!,
          percentage: 7.9,
        },
        {
          language: languages.find((l) => l.id === 'go')!,
          percentage: 7.1,
        },
        {
          language: languages.find((l) => l.id === 'java')!,
          percentage: 5,
        },
      ],

      updated_at: new Date('2024-10-24T11:45:14Z'),
    } satisfies DeveloperLanguages)
    return HttpResponse.json(resp)
  }),
  ...httpAsyncApi(`${BASE_URL}/developers/:username/region`, async ({ params }) => {
    if (params.username !== fakeDeveloper.login.toLowerCase()) {
      return developerNotFoundResponse.clone()
    }

    await delay(DEFAULT_ASYNC_DELAY)

    const resp = developerRegion.parse({
      id: fakeDeveloper.id,
      region: 'cn',
      confidence: 0.9,
    } satisfies DeveloperRegion)
    return HttpResponse.json(resp)
  }),

  http.get(`${BASE_URL}/rank`, async ({ request }) => {
    const url = new URL(request.url)
    const language = url.searchParams.get('language')
    const region = url.searchParams.get('region')
    const limit = Number(url.searchParams.get('limit'))

    // just let the linter happy
    void language
    void region

    if (!limit) {
      return HttpResponse.json(
        errorResponse.parse({
          code: 400,
          message: 'Invalid query parameters',
        } satisfies ErrorResponse),
        {
          status: 400,
        },
      )
    }

    return HttpResponse.json<Array<DeveloperWithPulsePoint>>(
      [
        {
          developer: fakeDeveloper,
          pulse_point: {
            id: fakeDeveloper.id,
            pulse_point: 233,

            updated_at: new Date('2024-10-24T11:45:14Z'),
          },
        },
      ],
      {
        status: 200,
      },
    )
  }),
] satisfies Array<HttpHandler>
