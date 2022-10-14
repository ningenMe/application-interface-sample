import {ChatServiceClient} from "../interface-grpc/ChatServiceClientPb";

const getHostName = (port: string) => {
  return process.env.NEXT_PUBLIC_SERVER_HOST!! + ":" + port;
}

export const serverGrpcGoClient = new ChatServiceClient(getHostName(process.env.NEXT_PUBLIC_SERVER_GRPC_ENVOY_GO_PORT!!))
