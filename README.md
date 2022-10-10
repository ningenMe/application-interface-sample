# application-interface-sample

### applications
|    name                                      | port | 
|--------------                                |----- |
|client-react                                  | 9100 |
|[server-grpc-go](./server-grpc-go)            | 9140 |
|[server-grpc-envoy-go](./server-grpc-envoy-go)| 9141 |                    
|database-mysql                                | 9190 |

### generate command
```shell
docker-compose -f docker-compose-generate-interface.yml build --no-cache && docker-compose -f docker-compose-generate-interface.yml up 
```

### application start
```shell
docker-compose -f docker-compose-application.yml build --no-cache && docker-compose -f docker-compose-application.yml up
```

### mysql login
```shell
mysql -uroot -h 127.0.0.1 -ppassword -P9190 -D sample
```
