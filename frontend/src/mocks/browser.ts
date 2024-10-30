import { RequestHandler } from 'msw'
import { setupWorker } from 'msw/browser'

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

  return setupWorker(...handlers)
}
