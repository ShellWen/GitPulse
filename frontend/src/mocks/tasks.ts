// We are just mocking tasks on the browser side, so don't wrong about memory leak.
const tasks: Record<string, Response | null> = {}

export type NewTaskResult = [taskId: string, completeTask: (obj: Response) => void]

const genTaskId = (function* (): Generator<string, never, unknown> {
  let id = 0
  while (true) {
    yield id.toString().padStart(8, '0')
    id++
  }
})()

export const newTask = (): NewTaskResult => {
  const taskId = genTaskId.next().value
  tasks[taskId] = null
  return [taskId, (obj: Response) => {
    tasks[taskId] = obj
  }]
}

export const getTask = (taskId: string): Response | null | undefined => {
  return tasks[taskId]
}

