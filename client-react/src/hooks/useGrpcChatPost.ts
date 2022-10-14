import { useState, useCallback, SyntheticEvent } from "react";
import {ChatServiceClient} from "../interface-grpc/ChatServiceClientPb";
import {PostRequest} from "../interface-grpc/chat_pb";

export const useGrpcChatPost = (client: ChatServiceClient) => {
  const [chatBody, setChatBody] = useState<string>("");

  const onChange = useCallback(
    (event: SyntheticEvent) => {
      const target = event.target as HTMLInputElement;
      setChatBody(target.value);
    },
    [setChatBody]
  );

  const onSubmit = useCallback(
    (event: SyntheticEvent) => {
      event.preventDefault();
      if (chatBody == "") {
        return;
      }
      const req = new PostRequest();
      req.setBody(chatBody);
      client.post(req, null, res => console.log(res));
      setChatBody("");
    },
    [client, chatBody]
  );

  return {
    chatBody,
    onChange,
    onSubmit
  };
};
