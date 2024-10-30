/**
 * HTTP error, when response invalid, including a query failed, parse failed, etc.
 */
export class HttpError extends Error {
  public response: Response

  constructor(response: Response, message: string | undefined) {
    super(message)
    this.name = this.constructor.name

    this.response = response
  }
}

/**
 * Business error, when response valid but business logic failed.
 */
export class BusinessError extends Error {
  public code: number
  public message: string

  constructor(code: number, message: string) {
    super(message)
    this.name = this.constructor.name

    this.code = code
    this.message = message
  }
}
