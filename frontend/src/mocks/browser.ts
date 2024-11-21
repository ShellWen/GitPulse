import { type ErrorResponse, errorResponse } from '$/lib/api/types.ts'
import { BASE_URL } from '$/mocks/constants.ts'
import { HttpResponse, RequestHandler, http } from 'msw'
import { setupWorker } from 'msw/browser'

const catchAll = http.all(`${BASE_URL}/*`, async ({ request }) =>
  HttpResponse.json(
    errorResponse.parse({
      code: -1,
      message: `API route not mocked: ${request.url.startsWith(BASE_URL) ? request.url.slice(BASE_URL.length) : request.url}`,
    } satisfies ErrorResponse),
    { status: 404 },
  ),
)

export const worker = async () => {
  const modules = import.meta.glob('./handlers/*.ts')
  const handlers = (
    await Promise.all(
      Object.values(modules).map(async (module) => {
        const { handlers } = (await module()) as {
          handlers?: Array<RequestHandler>
        }
        if (!handlers) {
          return []
        }
        return handlers
      }),
    )
  ).flat()

  return setupWorker(...handlers, catchAll)
}
