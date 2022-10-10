import type { NextPage } from 'next'
import Head from 'next/head'
import styles from '../styles/Home.module.css'

const Home: NextPage = () => {
  return (
    <div className={styles.container}>
      <Head>
        <title>Application Interface Sample</title>
        <meta name="description" content="application interface sample" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className={styles.main}>
        <h2>
            Application Interface Sample

        </h2>
        <h1 className={styles.title}>
            client-react
        </h1>

        <p className={styles.description}>
            implemented by {' '}
          <code className={styles.code}>next</code>
            {' '} and {' '}
          <code className={styles.code}>typescript</code>
        </p>

        <div className={styles.grid}>
          <a href="https://nextjs.org/docs" className={styles.card}>
            <h2>Interface: Grpc</h2>
            <h2>Language: Go &rarr;</h2>
          </a>

          <a href="https://nextjs.org/learn" className={styles.card}>
            <h2>Interface: Websocket</h2>
            <h2>Language: Kotlin &rarr;</h2>
          </a>

          <a href="https://nextjs.org/learn" className={styles.card}>
            <h2>Interface: Openapi</h2>
            <h2>Language: Rust &rarr;</h2>
          </a>

          <a href="https://nextjs.org/learn" className={styles.card}>
            <h2>Interface: Graphql</h2>
            <h2>Language: Kotlin &rarr;</h2>
          </a>

        </div>
      </main>

    </div>
  )
}

export default Home
