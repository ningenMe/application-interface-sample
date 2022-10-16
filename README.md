# application-interface-sample

### applications
|    name                                      | port | 
|--------------                                |----- |
|[client-react](./client-react)                | 9100 |
|[server-grpc-go](./server-grpc-go)            | 9140 |
|[server-grpc-envoy-go](./server-grpc-envoy-go)| 9141 |                    
|database-mysql                                | 9190 |

### env set
```shell
export `cat .env | awk '{print $1}' | xargs`
```

### client start
```shell
docker-compose -f docker-compose-application-client.yml build --no-cache --progress plain --remove-orphans && docker-compose -f docker-compose-application-client.yml up
```

### server start
```shell
docker-compose -f docker-compose-application-server.yml build --no-cache --progress plain --remove-orphans && docker-compose -f docker-compose-application-server.yml up
```

### generate command
```shell
docker-compose -f docker-compose-generate-interface.yml build --no-cache --progress plain --remove-orphans && docker-compose -f docker-compose-generate-interface.yml up 
```

### mysql login
```shell
mysql -uroot -h 127.0.0.1 -p${AIS_DATABASE_MYSQL_PASSWORD} -P${AIS_DATABASE_MYSQL_PORT} -D ${AIS_DATABASE_MYSQL_SCHEMA}
```
