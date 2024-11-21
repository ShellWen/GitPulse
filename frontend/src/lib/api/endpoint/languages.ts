import { type Language, language } from '$/lib/api/endpoint/types.ts'
import { typedFetch } from '$/lib/api/fetcher.ts'

export const getLanguages = async (): Promise<Array<Language>> => {
  return typedFetch('/languages', language.array())
}
