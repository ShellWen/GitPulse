import { type PropsWithChildren, useCallback } from 'react'

import { BusinessError } from '$/lib/api/error.ts'
import { QueryError } from '$/lib/query/error.ts'
import { Button } from 'react-daisyui'
import { ErrorBoundary, type FallbackProps } from 'react-error-boundary'
import { type Arguments, useSWRConfig } from 'swr'

const QueryErrorBoundaryBlock = ({ children }: PropsWithChildren) => {
  const { mutate } = useSWRConfig()
  const reset = useCallback(
    (key: Arguments) => {
      mutate(key)
    },
    [mutate],
  )

  const renderer = useCallback(({ resetErrorBoundary, error }: FallbackProps) => {
    if (!(error instanceof QueryError)) {
      throw error
    }
    const innerError = error.innerError
    if (innerError instanceof BusinessError && innerError.code === 404) {
      throw error
    }
    const errorMsg = error.message
    return (
      <div>
        请求失败：{errorMsg}
        <Button onClick={() => resetErrorBoundary(error.key)}>重试</Button>
      </div>
    )
  }, [])
  return (
    <ErrorBoundary onReset={reset} fallbackRender={renderer}>
      {children}
    </ErrorBoundary>
  )
}

export default QueryErrorBoundaryBlock
