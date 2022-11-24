import gql from 'graphql-tag'
import * as Urql from 'urql'
export type Maybe<T> = T | null
export type InputMaybe<T> = Maybe<T>
export type Exact<T extends { [key: string]: unknown }> = {
  [K in keyof T]: T[K]
}
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & {
  [SubKey in K]?: Maybe<T[SubKey]>
}
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & {
  [SubKey in K]: Maybe<T[SubKey]>
}
export type Omit<T, K extends keyof T> = Pick<T, Exclude<keyof T, K>>
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string
  String: string
  Boolean: boolean
  Int: number
  Float: number
  Time: any
}

export type Message = {
  __typename?: 'Message'
  createdAt: Scalars['Time']
  id: Scalars['String']
  text: Scalars['String']
  user: Scalars['String']
}

export type Mutation = {
  __typename?: 'Mutation'
  createTodo: Todo
  noop?: Maybe<NoopPayload>
  postMessage?: Maybe<Message>
}

export type MutationCreateTodoArgs = {
  input: NewTodo
}

export type MutationNoopArgs = {
  input?: InputMaybe<NoopInput>
}

export type MutationPostMessageArgs = {
  text: Scalars['String']
  user: Scalars['String']
}

export type NewTodo = {
  text: Scalars['String']
  userId: Scalars['String']
}

export type NoopInput = {
  clientMutationId?: InputMaybe<Scalars['String']>
}

export type NoopPayload = {
  __typename?: 'NoopPayload'
  clientMutationId?: Maybe<Scalars['String']>
}

export type Query = {
  __typename?: 'Query'
  messages: Array<Message>
  node: Scalars['String']
  todos: Array<Todo>
}

export type QueryNodeArgs = {
  id: Scalars['ID']
}

export type Subscription = {
  __typename?: 'Subscription'
  messagePosted: Message
  noop?: Maybe<NoopPayload>
}

export type SubscriptionMessagePostedArgs = {
  user: Scalars['String']
}

export type SubscriptionNoopArgs = {
  input?: InputMaybe<NoopInput>
}

export type Todo = {
  __typename?: 'Todo'
  done: Scalars['Boolean']
  id: Scalars['ID']
  text: Scalars['String']
  user: User
}

export type User = {
  __typename?: 'User'
  id: Scalars['ID']
  name: Scalars['String']
}

export type FindTodosQueryVariables = Exact<{ [key: string]: never }>

export type FindTodosQuery = {
  __typename?: 'Query'
  todos: Array<{
    __typename?: 'Todo'
    text: string
    done: boolean
    user: { __typename?: 'User'; name: string }
  }>
}

export const FindTodosDocument = gql`
  query findTodos {
    todos {
      text
      done
      user {
        name
      }
    }
  }
`

export function useFindTodosQuery(
  options?: Omit<Urql.UseQueryArgs<FindTodosQueryVariables>, 'query'>,
) {
  return Urql.useQuery<FindTodosQuery, FindTodosQueryVariables>({
    query: FindTodosDocument,
    ...options,
  })
}
