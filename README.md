# protoc-gen-defaults

Generates base implementation for gRPC services that allows to add methods to gRPC services without breaking existing builds.

### Example

Lets take a simple proto file like below,
```proto
service EchoService {
    rpc Echo(EchoRequest) returns (EchoResponse);
}
```
This will generate a file containing this,

```go
// BaseEchoServiceServer is the dummy implementation of the EchoServiceServer. Embed this into your own implementation
// to add new methods without breaking builds.
type BaseEchoServiceServer struct{}

// Echo is an unimplemented form of the method Echo
func (BaseEchoServiceServer) Echo(context.Context, *EchoRequest) (*EchoResponse, error) {
	return nil, status.Error(codes.Unimplemented, codes.Unimplemented.String())
}

```

It can then be used like this, Embed into your implementation.

```go
type server struct {
    pb.BaseEchoServiceServer	
}
```

Now `server` implements `EchoServiceServer`. 

### How to use it?

Just go get it and it will add the plugin to your path,

`go get -u github.com/srikrsna/protoc-gen-defaults`

Then use the plugin like below,

`protoc -I ./example --defaults_out=:./example ./example/example.proto`

