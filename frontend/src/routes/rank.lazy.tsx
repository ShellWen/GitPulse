import { createLazyFileRoute } from '@tanstack/react-router'

export const Route = createLazyFileRoute('/rank')({
  component: () => <div>TODO: Hello /rank!</div>,
})
