/// <reference types="vite/client" />

interface ImportMetaEnv {
  // The base URL of the backend API
  readonly VITE_BACKEND_BASEURL: string
  // Enable mock, set to any value to enable
  readonly VITE_ENABLE_MOCK: string
}

interface ImportMeta {
  readonly env: ImportMetaEnv
}
