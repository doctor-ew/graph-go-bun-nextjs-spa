// src/pages/_app.tsx
import { ApolloProvider } from '@apollo/client';
import { apolloClient } from '@/lib/apolloClient';
import type { AppProps } from 'next/app';
import '@/styles/globals.css';

export default function MyApp({ Component, pageProps }: AppProps) {
    return (
        <ApolloProvider client={apolloClient}>
            <Component {...pageProps} />
        </ApolloProvider>
    );
}
