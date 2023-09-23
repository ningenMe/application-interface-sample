import type { NextPage } from 'next'
import Head from 'next/head'
import styles from 'styles/Index.module.css'
import commonStyles from 'styles/Common.module.css'

const Home: NextPage = () => {
  return (
    <>
      <Head>
        <title>client-react</title>
        <meta name="description" content="client-react" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <div className={commonStyles.container}>

        <main className={styles.main}>
          <h2>
            Application Interface Sample
          </h2>
          <h1 className={commonStyles.title}>
            client-react
          </h1>

          <p className={styles.description}>
            implemented by {' '}
            <code className={styles.code}>next</code>
            {' '} and {' '}
            <code className={styles.code}>typescript</code>
          </p>

          <div className={styles.grid}>
            <a href="/server-go-grpc" className={styles.card}>
              <h2>Interface: Grpc</h2>
              <h2>Language: Go &rarr;</h2>
            </a>

            {/*<a href="https://nextjs.org/learn" className={styles.card}>*/}
            {/*  <h2>Interface: Openapi</h2>*/}
            {/*  <h2>Language: Go &rarr;</h2>*/}
            {/*</a>*/}

          </div>
        </main>
      </div>
    </>
  )
}

export default Home
