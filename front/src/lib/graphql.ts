import {
  createClient,
  dedupExchange,
  cacheExchange,
  subscriptionExchange,
  Client,
  Exchange,
} from 'urql'
import { multipartFetchExchange } from '@urql/exchange-multipart-fetch'
import { createClient as createWSClient } from 'graphql-ws'

export const createUrqlClient = (): Client => {
  const graphqlEndpoint = 'http://localhost:8080/query'
  const graphqlWebsocketEndpoint =
    'ws://localhost:8080/query'

  const exchanges: Exchange[] | undefined = [
    dedupExchange,
    cacheExchange,
    multipartFetchExchange,
  ]

  if (typeof window !== 'undefined') {
    const wsClient = createWSClient({
      url: graphqlWebsocketEndpoint,
    })

    const subExchange = subscriptionExchange({
      forwardSubscription: (operation) => ({
        subscribe: (sink) => ({
          unsubscribe: wsClient.subscribe(operation, sink),
        }),
      }),
    })

    exchanges.push(subExchange)
  }

  const client = createClient({
    url: graphqlEndpoint,
    requestPolicy: 'cache-and-network',
    exchanges,
    fetchOptions: () => ({
      credentials: 'include',
    }),
  })

  return client
}
