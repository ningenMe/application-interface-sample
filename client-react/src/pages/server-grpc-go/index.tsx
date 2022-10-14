import type {GetServerSideProps, NextPage} from 'next'
import Head from 'next/head'
import styles from 'styles/Chat.module.css'
import commonStyles from 'styles/Common.module.css'
import {useState} from "react";
import { Empty } from 'google-protobuf/google/protobuf/empty_pb'
import {Chat} from "interface-grpc/chat_pb";
import {serverGrpcGoClient} from "../../repository/server";
import {useDidUpdateEffect} from "../../hooks/useDidUpdateEffect";

const Index: NextPage = () => {

  const [chatList, setChatList] = useState<ReadonlyArray<Chat>>([]);

  useDidUpdateEffect(() => {
      const stream = serverGrpcGoClient.get(new Empty());
      stream.on("data", response => {
        const fetchedList = response.getChatlistList();

        setChatList(chatList =>
          Array
            .from(
              new Map(
                [...chatList,...fetchedList].map(chat => [chat.getChatid(), chat])
              ).values())
            .sort((l,r)=> r.getChatid() - l.getChatid())
            .slice(0,5)
            .sort((l,r)=> l.getChatid() - r.getChatid())
        )
      })
    }, [serverGrpcGoClient]
  )

  const chatListNode = chatList.map(chat => {
    return (
      <div className={styles.chatMessage} key={chat.getChatid()}>
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
          <div className={styles.chatWindow} >
            {chatListNode}
          </div>
        </main>
      </div>
    </>
  )
}

export default Index
