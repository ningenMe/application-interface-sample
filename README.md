# application-interface-sample

### applications
|    name             | port |
|--------------       |----- |
|client-go            | 9000 |
|client-react         | 9001 |
|server-grpc-go       | 9020 |
|server-grpc-envoy-go | 9021 |

### generate command
```shell
docker-compose -f docker-compose-generate-interface.yml build --no-cache &&  docker-compose -f docker-compose-generate-interface.yml up 
```

