# application-interface-sample

### applications
|    name                                      | port | 
|--------------                                |----- |
|[client-react](./client-react)                | 9100 |
|[server-grpc-go](./server-grpc-go)            | 9140 |
|[server-grpc-envoy-go](./server-grpc-envoy-go)| 9141 |                    
|database-mysql                                | 9190 |

### client start
```shell
docker-compose -f docker-compose-application-client.yml build --no-cache && docker-compose -f docker-compose-application-client.yml up
```

### server start
```shell
docker-compose -f docker-compose-application-server.yml build --no-cache && docker-compose -f docker-compose-application-server.yml up
```

### generate command
```shell
docker-compose -f docker-compose-generate-interface.yml build --no-cache && docker-compose -f docker-compose-generate-interface.yml up 
```

### mysql login
```shell
mysql -uroot -h 127.0.0.1 -ppassword -P9190 -D sample
```
