import type {NextPage} from 'next'
import Head from 'next/head'
import styles from 'styles/Chat.module.css'
import commonStyles from 'styles/Common.module.css'
import {serverGrpcGoClient} from "../../repository/server";
import {useGrpcChatGet} from "../../hooks/useGrpcChatGet";
import {useGrpcChatPost} from "../../hooks/useGrpcChatPost";

const Index: NextPage = () => {

  const chatList = useGrpcChatGet(serverGrpcGoClient)
  const {chatBody,onChange,onSubmit} = useGrpcChatPost(serverGrpcGoClient)

  const chatListNode = chatList.map(chat => {
    return (
      <div className={styles.chatBody} key={chat.getChatid()}>
        <p> {chat.getChatid()} : {chat.getCreatedat()} : {chat.getBody()} </p>
      </div>
    );
  })

  return (
    <>
      <Head>
        <title>server-grpc-go - client-react</title>
        <meta name="description" content="client-react" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <div className={commonStyles.container}>
        <main className={styles.main}>
          <h1 className={commonStyles.title}>
            server-grpc-go
          </h1>
          <div className={styles.chatWindow} >
            {chatListNode}
          </div>
          <div className={styles.chatForm}>
            <form onSubmit={onSubmit}>
              <input type="text" value={chatBody} onChange={onChange} className={styles.chatFormInput} />
              <button type="submit" className={styles.chatFormButton}> {"▲"} </button>
            </form>
          </div>
        </main>
      </div>
    </>
  )
}

export default Index
