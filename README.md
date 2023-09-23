# application-interface-sample

### application port
|    name                                      | port | type     | 
|--------------                                |----- | ----     |
|[client-react](./client-react)                | 9100 | client   |
|[client-java](./client-java)                  | 9101 | client   |
|[server-grpc-go](./server-grpc-go)            | 9140 | server   |
|[server-grpc-envoy-go](./server-grpc-envoy-go)| 9141 | server   |
|[database-mysql](./database-mysql)            | 9190 | database |

### server-roadmap
|  language  | grpc | connect | websocket | graphql | openapi(rest) | quic |
|--          | --- | ----    |        -- | --      | --            | --   |
| go         | [server-grpc-go](./server-grpc-go) | TODO    | TODO | TODO | TODO            | TODO   |
| java         | TODO | TODO    | TODO | TODO | TODO            | TODO   |
| kotlin         | TODO | TODO    | TODO | TODO | TODO            | TODO   |
| rust         | TODO | TODO    | TODO | TODO | TODO            | TODO   |



### architecture

```mermaid
flowchart LR
    subgraph client
        react([react])
        java([java])
    end
    subgraph server
        subgraph grpc-go-and-envoy
            grpc-go-envoy[[grpc-go-envoy]]
            grpc-go[[grpc-go]]
            grpc-go-envoy <--> grpc-go
        end
        connect-go[[connect-go: TODO]]
        websocket-go[[websocket-go: TODO]]
    end
    subgraph database
        mysql[(mysql)]
    end

    react <--> grpc-go-envoy
    java  <--> grpc-go
    react & java <--> connect-go 
    react & java <--> websocket-go

    grpc-go & connect-go & websocket-go <--> mysql
```

### client start
```shell
docker-compose -f docker-compose-application-client.yml build --no-cache --progress plain && docker-compose -f docker-compose-application-client.yml up
```

### server start
```shell
docker-compose -f docker-compose-application-server.yml build --no-cache --progress plain && docker-compose -f docker-compose-application-server.yml up
```

### generate command
```shell
docker-compose -f docker-compose-generate-interface.yml build --no-cache --progress plain && docker-compose -f docker-compose-generate-interface.yml up 
```

### env set (if you want to run application without docker)
```shell
export `cat .env | awk '{print $1}' | xargs`
```
