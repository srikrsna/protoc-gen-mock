# protoc-gen-mock

Generates mock implementation for gRPC services that allows to use gRPC services in UI Testing and Mock Testing.

### Example

Lets take a simple proto file like below,
```proto
service EchoService {
    rpc Echo(EchoRequest) returns (EchoResponse);
}
```
This will generate a file containing this,

```go
// MockEchoServiceServer is the mock implementation of the EchoServiceServer. Use this to create mock services that
// return random data. Useful in UI Testing.
type MockEchoServiceServer struct{}

// Echo is mock implementation of the method Echo
func (MockEchoServiceServer) Echo(context.Context, *EchoRequest) (*EchoResponse, error) {
	var res EchoResponse
	fuzzer.Fuzz(&res)
	return &res, nil
}

```

### How to use it?

Just go get it and it will add the plugin to your path,

`go get -u github.com/srikrsna/protoc-gen-mock`

Then use the plugin like below,

`protoc -I ./example --mock_out=:./example ./example/example.proto`

