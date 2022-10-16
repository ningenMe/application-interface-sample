# application-interface-sample

### applications
|    name                                      | port | type     | 
|--------------                                |----- | ----     |
|[client-react](./client-react)                | 9100 | client   |
|[client-java](./client-java)                  | 9101 | client   |
|[server-grpc-go](./server-grpc-go)            | 9140 | server   |
|[server-grpc-envoy-go](./server-grpc-envoy-go)| 9141 | server   |
|[database-mysql](./database-mysql)            | 9190 | database |

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
