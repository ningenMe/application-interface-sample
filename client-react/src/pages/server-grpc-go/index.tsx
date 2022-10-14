import type {GetServerSideProps, NextPage} from 'next'
import Head from 'next/head'
import styles from 'styles/Chat.module.css'
import commonStyles from 'styles/Common.module.css'
import {ChatServiceClient} from "interface-grpc/ChatServiceClientPb";
import {useEffect, useState} from "react";
import { Empty } from 'google-protobuf/google/protobuf/empty_pb'
import {Chat} from "interface-grpc/chat_pb";

type ServerConfig ={
  host: string
  port: string
}
const getHostName = (serverConfig: ServerConfig) => {
  return serverConfig.host + ":" + serverConfig.port;
}

export const getServerSideProps: GetServerSideProps = async () => {
  return {
    props: {
      host: process.env.SERVER_HOST!!,
      port: process.env.SERVER_GRPC_ENVOY_GO_PORT!!
    }
  }
};

const Index: NextPage<ServerConfig> = (props: ServerConfig) => {

  const [chatList, setChatList] = useState<Chat[]>([]);

  const client = new ChatServiceClient(getHostName(props))

  useEffect(
    () => {
      const stream = client.get(new Empty());
      stream.on("data", response => {
        const list = response.getChatlistList();
        console.log(list)

        //新しいデータがないならreturn
        if (!list.length) {
          return;
        }
        setChatList([...chatList, ...list])
      })
    },
    [client]
  )

  const chatNodeList = chatList.map(chat => {
    return (
      <p> {chat.getChatid()} : {chat.getCreatedat()} : {chat.getBody()} </p>
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
          <div className={styles.chatWindow} >
            <p> TODO メッセージ送信ボックスを作る </p>
            {chatNodeList}
          </div>
        </main>
      </div>
    </>
  )
}

export default Index
