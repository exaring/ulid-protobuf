Encode [ulid](https://github.com/oklog/ulid/)s as protobuf directly.

# Usage

1) import the proto file in your proto code  

   `mymessage.proto`:  
   ```proto
   syntax = "proto3";
   
   package mypackage;
   
   import "ulid.proto";
   
   option go_package = "example.org/mypackage";
   
   message MyMessage {
     ulid.ULID id = 1;
   }
   
   ```

2) add this library to your code:
   ```
   $ go get github.com/exaring/ulid-protobuf@latest
   ```

3) Then compile your proto files with `protoc`, e.g.:
   ```
   $ protoc --proto_path=$(go list -m -f '{{.Dir}}' github.com/exaring/ulid-protobuf) \
            --go_out=paths=source_relative:. \
            --proto_path=. ./mymessage.proto
   ```

   Note the trick with `--proto_path`: it allows `protoc` to find the `ulid.proto` file inside the go module's source.