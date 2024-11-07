import { createFileRoute } from '@tanstack/react-router'

export interface RankPageSearch {
  language?: string
  region?: string
  limit: number
}

export const Route = createFileRoute('/rank')({
  component: () => <div>Hello /rank!</div>,
  validateSearch: (search: Record<string, unknown>): RankPageSearch => {
    if (typeof search.language !== 'string' && search.language !== undefined) {
      throw new Error('Invalid language')
    }

    if (typeof search.region !== 'string' && search.region !== undefined) {
      throw new Error('Invalid region')
    }

    if (typeof search.limit !== 'string' && typeof search.limit !== 'number' && search.limit !== undefined) {
      throw new Error('Invalid limit')
    }

    return {
      language: search.language,
      region: search.region,
      limit: Number(search.limit ?? 50),
    }
  }
})
