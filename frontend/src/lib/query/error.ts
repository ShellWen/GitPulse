import type { Arguments as SWRArguments } from 'swr'

/**
 * Query error, wrap the inner error and query key.
 */
export class QueryError extends Error {
  public innerError: Error
  public key: SWRArguments

  constructor(innerError: Error, key: SWRArguments) {
    super(`Query error: ${innerError.message}`)
    this.name = this.constructor.name

    this.innerError = innerError
    this.key = key
  }
}