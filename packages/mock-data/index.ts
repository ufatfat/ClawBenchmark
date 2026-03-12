import type { Agent, BenchmarkResult, User, Comment } from '@clawbenchmark/types'
import agentsData from './agents.json'
import benchmarksData from './benchmarks.json'
import usersData from './users.json'
import commentsData from './comments.json'

export const agents = agentsData as Agent[]
export const benchmarks = benchmarksData as BenchmarkResult[]
export const users = usersData as User[]
export const comments = commentsData as Comment[]
