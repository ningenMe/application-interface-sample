// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: chat.proto

package sample.generated;

public final class ChatProto {
  private ChatProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_sample_PostRequest_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_sample_PostRequest_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_sample_GetResponse_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_sample_GetResponse_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_sample_Chat_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_sample_Chat_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\nchat.proto\022\006sample\032\033google/protobuf/em" +
      "pty.proto\"\033\n\013PostRequest\022\014\n\004body\030\001 \001(\t\"-" +
      "\n\013GetResponse\022\036\n\010chatList\030\001 \003(\0132\014.sample" +
      ".Chat\"7\n\004Chat\022\016\n\006chatId\030\001 \001(\003\022\014\n\004body\030\002 " +
      "\001(\t\022\021\n\tcreatedAt\030\003 \001(\t2|\n\013ChatService\0225\n" +
      "\004Post\022\023.sample.PostRequest\032\026.google.prot" +
      "obuf.Empty\"\000\0226\n\003Get\022\026.google.protobuf.Em" +
      "pty\032\023.sample.GetResponse\"\0000\001Bo\n\020sample.g" +
      "eneratedB\tChatProtoP\001ZNgithub.com/ningen" +
      "me/application-interface-sample/server-g" +
      "o-grpc/interface-grpcb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          com.google.protobuf.EmptyProto.getDescriptor(),
        });
    internal_static_sample_PostRequest_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_sample_PostRequest_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_sample_PostRequest_descriptor,
        new java.lang.String[] { "Body", });
    internal_static_sample_GetResponse_descriptor =
      getDescriptor().getMessageTypes().get(1);
    internal_static_sample_GetResponse_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_sample_GetResponse_descriptor,
        new java.lang.String[] { "ChatList", });
    internal_static_sample_Chat_descriptor =
      getDescriptor().getMessageTypes().get(2);
    internal_static_sample_Chat_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_sample_Chat_descriptor,
        new java.lang.String[] { "ChatId", "Body", "CreatedAt", });
    com.google.protobuf.EmptyProto.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
