import { createFileRoute, redirect } from '@tanstack/react-router'

export const Route = createFileRoute('/u_/$username')({
  beforeLoad: async ({ params }) => {
    if (params.username !== params.username.toLowerCase()) {
      // Redirect to the lowercase version of the username
      redirect({
        to: `/u/${params.username.toLowerCase()}`,
        throw: true,
      })
    }
  },
})
