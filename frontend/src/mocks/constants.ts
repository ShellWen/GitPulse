export const BASE_URL = import.meta.env.VITE_BACKEND_BASEURL!
if (!BASE_URL) {
  throw new Error('VITE_BACKEND_BASEURL is not defined')
}
