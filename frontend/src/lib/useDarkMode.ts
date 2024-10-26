import { useMediaQuery } from 'react-responsive'

const useDarkMode = () => {
  return useMediaQuery({
    query: '(prefers-color-scheme: dark)',
  })
}

export default useDarkMode