import * as jspb from 'google-protobuf'

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';


export class PostRequest extends jspb.Message {
  getBody(): string;
  setBody(value: string): PostRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PostRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PostRequest): PostRequest.AsObject;
  static serializeBinaryToWriter(message: PostRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PostRequest;
  static deserializeBinaryFromReader(message: PostRequest, reader: jspb.BinaryReader): PostRequest;
}

export namespace PostRequest {
  export type AsObject = {
    body: string,
  }
}

export class GetResponse extends jspb.Message {
  getChatlistList(): Array<Chat>;
  setChatlistList(value: Array<Chat>): GetResponse;
  clearChatlistList(): GetResponse;
  addChatlist(value?: Chat, index?: number): Chat;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetResponse): GetResponse.AsObject;
  static serializeBinaryToWriter(message: GetResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetResponse;
  static deserializeBinaryFromReader(message: GetResponse, reader: jspb.BinaryReader): GetResponse;
}

export namespace GetResponse {
  export type AsObject = {
    chatlistList: Array<Chat.AsObject>,
  }
}

export class Chat extends jspb.Message {
  getBody(): string;
  setBody(value: string): Chat;

  getCreatedat(): string;
  setCreatedat(value: string): Chat;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Chat.AsObject;
  static toObject(includeInstance: boolean, msg: Chat): Chat.AsObject;
  static serializeBinaryToWriter(message: Chat, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Chat;
  static deserializeBinaryFromReader(message: Chat, reader: jspb.BinaryReader): Chat;
}

export namespace Chat {
  export type AsObject = {
    body: string,
    createdat: string,
  }
}

