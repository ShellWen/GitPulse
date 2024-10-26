import { createFileRoute, redirect } from '@tanstack/react-router'

export const Route = createFileRoute('/u_/$userName')({
  beforeLoad: async ({ params }) => {
    if (params.userName !== params.userName.toLowerCase()) {
      // Redirect to the lowercase version of the username
      redirect({
        to: `/u/${params.userName.toLowerCase()}`,
        throw: true,
      })
    }
  },
})
