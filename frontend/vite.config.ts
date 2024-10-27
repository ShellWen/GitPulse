import { defineConfig } from 'vite'
import { TanStackRouterVite } from '@tanstack/router-plugin/vite'
import react from '@vitejs/plugin-react'

// https://vite.dev/config/
export default defineConfig({
  plugins: [TanStackRouterVite(), react()],

  resolve: {
    alias: {
      '$': '/src',
    },
  },

  build: {
    rollupOptions: {
      output: {
        manualChunks: {
          // Split antd and antd-charts into separate chunks cuz they are too large
          'ant-charts': ['@ant-design/charts', '@ant-design/plots'],
        }
      }
    }
  }
})
