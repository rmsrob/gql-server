# GQL - SERVER

> Get custom resolvers for each GRPC server to only be called when data is needed.
> Add Custom resolver per fields in `gqlgen.yml` / `graph/schema.resolver`

## Install

#### GO / GRPC

> You'll need to [install Go](https://golang.org/doc/install)
> Install [grpc](https://grpc.io/docs/protoc-installation/)
>
> ```sh
> brew install protobuf
> brew install protoc-gen-go
> brew install protoc-gen-go-grpc
> ```

```sh
make dep
make proto
```

## Launch

We are using 3 grpc servers for our graphql
- launch on an other terminal a grpc-server as inside the example
- `make run`


> open [http://localhost:8080/](http://localhost:8080/)

```js
query servers {
  returnTime {
    server_1
    server_3
  }
}
```

> server_2 should not be called, check Gin log or have any error from server_2

```log
[GIN] 2022/08/02 - 09:30:39 | 200 |     2.68225ms |             ::1 | POST     "/query"
```

> If you don't call that field (that have it's own custom resolver), 
> the resolver itself never get called