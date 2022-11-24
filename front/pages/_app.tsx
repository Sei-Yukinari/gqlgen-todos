import '../styles/globals.css'
import type {AppProps} from 'next/app'
import {createUrqlClient} from '@/lib/graphql'
import {Provider} from 'urql'

export default function App({
                                Component,
                                pageProps,
                            }: AppProps) {
    return (
        <Provider value={createUrqlClient()}>
            <Component {...pageProps} />
        </Provider>
    )
}
