import { BASE_URL } from '$/lib/api/constants.ts'
import { type Language, language } from '$/lib/api/endpoint/types.ts'
import fetchWrapped from '$/lib/api/fetchWrapped.ts'

export const getLanguages = async (): Promise<Array<Language>> => {
  return fetchWrapped(`${BASE_URL}/languages`, language.array())
}
