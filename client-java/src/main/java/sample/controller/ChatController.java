package sample.controller;

import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import sample.repository.ServerGrpcGoRepository;

@RestController
@RequiredArgsConstructor
public class ChatController {

    private final ServerGrpcGoRepository serverGrpcGoRepository;

    @GetMapping("server-grpc-go/get")
    public ResponseEntity<String> serverGrpcGoGet() {
        serverGrpcGoRepository.get();
        return ResponseEntity.ok("ok");
    }

    @GetMapping("server-grpc-go/post")
    public ResponseEntity<String> serverGrpcGoPost(
        @RequestParam String body
    ) {
        serverGrpcGoRepository.post(body);
        return ResponseEntity.ok("ok");
    }

}
