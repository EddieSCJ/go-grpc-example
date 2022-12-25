# Go with gRPC

## Installing

You'll need to install some stuff before start.
* [Go 1.19](https://go.dev/dl/)
* [Protobuf compiler](https://grpc.io/docs/protoc-installation/)
* [Go plugins for protobuf and codegen](https://grpc.io/docs/languages/go/quickstart/)

## First Steps
Once you have everything installed, create a repo for your go code
and after this clone it.

Inside your git project type
```bash
go mod init <link-to-your-repo>

Ex:

go mod init github.com/EddieSCJ/go-grpc-example-example
```

## Tips

Once you created your .proto file you can generate the code
with protoc. Just type in your terminal:
```bash
protoc --go_out=. --go-grpc_out=. proto/category/course_category.proto

```
---
You'll need to install [evans](https://github.com/ktr0731/evans#macos) to test you grpc api.

If you are in mac do the following, if not, click in the link and follow the instructions
```bash
brew tap ktr0731/evans
brew install evans
```

---
We are running it in port `50051` because of evans (the client we are using to 
test our server). If you want to change it, you'll need to change it in the `main.go` file.

