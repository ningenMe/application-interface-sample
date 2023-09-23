import {ChatServiceClient} from "../interface-grpc/ChatServiceClientPb";

const getHostName = (port: string) => {
  return "http://localhost:" + port;
}

export const serverGrpcGoClient = new ChatServiceClient(getHostName(process.env.NEXT_PUBLIC_AIS_SERVER_GO_GRPC_ENVOY_PORT!!))
