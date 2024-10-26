import type { Config } from 'tailwindcss'
import daisyui from 'daisyui'

export default {
  content: [
    './index.html',
    './src/**/*.{html,js,jsx,tsx,js,ts}',
    'node_modules/daisyui/dist/**/*.js',
    'node_modules/react-daisyui/dist/**/*.js',
  ],
  theme: {
    extend: {},
  },
  plugins: [daisyui],

  daisyui: {
    themes: ['light', 'night'],
    darkTheme: 'night',
  },
} satisfies Config

