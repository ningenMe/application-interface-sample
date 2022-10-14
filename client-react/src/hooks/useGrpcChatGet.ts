import {useState} from "react";
import {Chat} from "../interface-grpc/chat_pb";
import {useDidUpdateEffect} from "./useDidUpdateEffect";
import {Empty} from "google-protobuf/google/protobuf/empty_pb";
import {ChatServiceClient} from "../interface-grpc/ChatServiceClientPb";

export const useGrpcChatGet = (client: ChatServiceClient):ReadonlyArray<Chat>  => {

  const [chatList, setChatList] = useState<ReadonlyArray<Chat>>([]);

  useDidUpdateEffect(() => {
      const stream = client.get(new Empty());
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
    }, [client]
  )

  return chatList
}
