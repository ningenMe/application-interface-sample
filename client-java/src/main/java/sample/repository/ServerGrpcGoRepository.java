package sample.repository;

import com.google.protobuf.Empty;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.grpc.StatusRuntimeException;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Repository;
import sample.generated.Chat;
import sample.generated.ChatServiceGrpc;
import sample.generated.GetResponse;
import sample.generated.PostRequest;

import java.util.Iterator;

@Repository
public class ServerGrpcGoRepository {

    private final ClientJavaConfig clientJavaConfig;
    private final ChatServiceGrpc.ChatServiceBlockingStub chatServiceBlockingStub;

    public ServerGrpcGoRepository(ClientJavaConfig clientJavaConfig) {
        this.clientJavaConfig = clientJavaConfig;
        this.chatServiceBlockingStub = ChatServiceGrpc.newBlockingStub(
            ManagedChannelBuilder
                .forAddress("127.0.0.1", clientJavaConfig.getAisServerGrpcGoPort())
                .usePlaintext()
                .build()
        );
    }

    public void get() {
        try {
            final Iterator<GetResponse> iterator = chatServiceBlockingStub.get(Empty.newBuilder().build());
            while(iterator.hasNext()) {
                final GetResponse getResponse = iterator.next();
                for (final Chat chat:getResponse.getChatListList()) {
                    System.out.println(chat.toString());
                }
            }
        } catch (StatusRuntimeException ex) {
            System.out.println(ex.getMessage());
        }
    }

    public void post(final String body) {
        try {
            chatServiceBlockingStub.post(
                PostRequest.newBuilder().setBody(body).build()
            );
        } catch (StatusRuntimeException ex) {
            System.out.println(ex.getMessage());
        }
    }

}
