grpc-entgo-cqrs-example

#### 👨‍💻 Full list what has been used:
[gRPC](https://github.com/grpc/grpc-go) Go implementation of gRPC<br/>
[entgo](https://entgo.io) Go implementation of entgo<br/>

``` mermaid
graph LR
    A[Client]  --> B(Write_SerVice)
    A --> C(Read_SerVice)
    B --> D{DB}
    C --> D
``` 