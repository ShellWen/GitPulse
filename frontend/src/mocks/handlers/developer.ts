import {
  type Developer,
  type DeveloperLanguages,
  type DeveloperPulsePoint,
  type DeveloperRegion,
  type Language,
  developer,
  developerLanguages,
  developerPulsePoint,
  developerRegion, type DeveloperWithPulsePoint,
} from '$/lib/api/endpoint/types.ts'
import { type ErrorResponse, errorResponse } from '$/lib/api/types.ts'
import { HttpHandler, HttpResponse, delay, http } from 'msw'

import { BASE_URL } from '../constants.ts'
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

// noinspection JSUnusedGlobalSymbols
export const handlers = [
  http.get(`${BASE_URL}/developers/:username`, async ({ params }) => {
    await delay(1000)

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

    const resp = developer.parse(fakeDeveloper)
    return HttpResponse.json<Developer>(resp)
  }),
  http.get(`${BASE_URL}/developers/:username/pulse-point`, async ({ params }) => {
    await delay(4000)

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

    const resp = developerPulsePoint.parse({
      pulse_point: {
        id: fakeDeveloper.id,
        pulse_point: 233,

        updated_at: new Date('2024-10-24T11:45:14Z'),
      },
    } satisfies DeveloperPulsePoint)
    return HttpResponse.json(resp)
  }),
  http.get(`${BASE_URL}/developers/:username/languages`, async ({ params }) => {
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

    const resp = developerLanguages.parse({
      languages: {
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
      },
    } satisfies DeveloperLanguages)

    return HttpResponse.json(resp)
  }),
  http.get(`${BASE_URL}/developers/:username/region`, async ({ params }) => {
    await delay(7000)

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

    const resp = developerRegion.parse({
      region: {
        id: fakeDeveloper.id,
        region: 'cn',
        confidence: 0.9,
      },
    } satisfies DeveloperRegion)
    return HttpResponse.json(resp)
  }),
  http.get(`${BASE_URL}/developers`, async ({ request }) => {
    const url = new URL(request.url)
    const languageId = url.searchParams.get('language')
    const region = url.searchParams.get('region')
    const limit = Number(url.searchParams.get('limit'))

    // just let the linter happy
    void(languageId)
    void(region)

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

    return HttpResponse.json<Array<DeveloperWithPulsePoint>>([{
      developer: fakeDeveloper,
      pulse_point: {
        pulse_point: {
          id: fakeDeveloper.id,
          pulse_point: 233,

          updated_at: new Date('2024-10-24T11:45:14Z'),
        },
      }
    }], {
      status: 200,
    })
  }),
] satisfies Array<HttpHandler>
