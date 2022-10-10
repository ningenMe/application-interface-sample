import type { NextPage } from 'next'
import Head from 'next/head'
import styles from '../../styles/Index.module.css'

const Index: NextPage = () => {
  return (
    <div className={styles.container}>
      <Head>
        <title>server-grpc-go - client-react</title>
        <meta name="description" content="client-react" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main className={styles.main}>
          <p> TODO メッセージ送信ボックスを作る </p>
          <p> TODO メッセージ参照表示を作る </p>
      </main>
    </div>
  )
}

export default Index
