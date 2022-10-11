import type { NextPage } from 'next'
import Head from 'next/head'
import styles from '../../styles/Chat.module.css'
import commonStyles from '../../styles/Common.module.css'

const Index: NextPage = () => {
  return (
    <>
      <Head>
        <title>server-grpc-go - client-react</title>
        <meta name="description" content="client-react" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <div className={commonStyles.container}>
        <main className={styles.main}>
          <div className={styles.chatWindow} >
            <p> TODO メッセージ送信ボックスを作る </p>
            <p> TODO メッセージ参照表示を作る </p>
          </div>
        </main>
      </div>
    </>
  )
}

export default Index
