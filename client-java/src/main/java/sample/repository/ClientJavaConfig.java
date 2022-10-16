package sample.repository;

import lombok.Getter;
import lombok.Setter;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.context.annotation.Configuration;

@Configuration
@ConfigurationProperties(prefix = "client-java")
@Getter
@Setter
public class ClientJavaConfig {

    private String aisServerHost;

    private Integer aisServerGrpcGoPort;
}
