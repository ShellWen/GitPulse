import { type Language, language } from '$/lib/api/endpoint/types.ts'
import { HttpHandler, HttpResponse, http } from 'msw'

import { BASE_URL } from '../constants.ts'
import languagesJson from './languages.json'

// noinspection JSUnusedGlobalSymbols
export const handlers = [
  http.get(`${BASE_URL}/languages`, async () => {
    const resp = language.array().parse(languagesJson)
    return HttpResponse.json<Array<Language>>(resp)
  }),
] satisfies Array<HttpHandler>
