import { type PropsWithChildren, useCallback } from 'react'

import { HttpError } from '$/lib/api/error.ts'
import { QueryErrorResetBoundary } from '@tanstack/react-query'
import { Button } from 'react-daisyui'
import { ErrorBoundary, type FallbackProps } from 'react-error-boundary'

const QueryErrorBoundaryBlock = ({ children }: PropsWithChildren) => {
  const renderer = useCallback(({ resetErrorBoundary, error }: FallbackProps) => {
    if (error instanceof HttpError && error.response.status === 404) {
      throw error
    }
    const errorMsg = error instanceof Error ? error.message : '未知错误'
    return (
      <div>
        请求失败：{errorMsg}
        <Button onClick={() => resetErrorBoundary()}>重试</Button>
      </div>
    )
  }, [])
  return (
    <QueryErrorResetBoundary>
      {({ reset }) => (
        <ErrorBoundary onReset={reset} fallbackRender={renderer}>
          {children}
        </ErrorBoundary>
      )}
    </QueryErrorResetBoundary>
  )
}

export default QueryErrorBoundaryBlock
