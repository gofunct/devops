# gRPC-Go

### Description

<details><summary>show</summary>
<p>

The Go implementation of gRPC: A high performance, open source, general RPC framework that puts mobile and HTTP/2 first. For more information see the gRPC Quick Start: Go guide.

</p>
</details>

Installation
To install this package, you need to install Go and setup your Go workspace on your computer. The simplest way to install the library is to run:

$ go get -u google.golang.org/grpc
Prerequisites
This requires Go 1.6 or later. Go 1.7 will be required soon.

Constraints
The grpc package should only depend on standard Go packages and a small number of exceptions. If your contribution introduces new dependencies which are NOT in the list, you need a discussion with gRPC-Go authors and consultants.

Documentation
See API documentation for package and API descriptions and find examples in the examples directory.

Performance
See the current benchmarks for some of the languages supported in this dashboard.

Status
General Availability Google Cloud Platform Launch Stages.

FAQ
Compiling error, undefined: grpc.SupportPackageIsVersion
Please update proto package, gRPC package and rebuild the proto files:

go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u google.golang.org/grpc
protoc --go_out=plugins=grpc:. *.proto


# gRPC command line tool
Overview
This document describes the command line tool that comes with gRPC repository. It is desirable to have command line tools written in other languages roughly follow the same syntax and flags.

At this point, the tool needs to be built from source, and it should be moved out to grpc-tools repository as a stand alone application once it is mature enough.

Core functionality
The command line tool can do the following things:

Send unary rpc.
Attach metadata and display received metadata.
Handle common authentication to server.
Infer request/response types from server reflection result.
Find the request/response types from a given proto file.
Read proto request in text form.
Read request in wire form (for protobuf messages, this means serialized binary form).
Display proto response in text form.
Write response in wire form to a file.
The command line tool should support the following things:

List server services and methods through server reflection.
Fine-grained auth control (such as, use this oauth token to talk to the server).
Send streaming rpc.
Code location
To use the tool, you need to get the grpc repository and make sure your system has the prerequisites for building grpc from source, given in the installation instructions.

In order to build the grpc command line tool from a fresh clone of the grpc repository, you need to run the following command to update submodules:

git submodule update --init
You also need to have the gflags library installed on your system. gflags can be installed with the following command: Linux:

sudo apt-get install libgflags-dev
Mac systems with Homebrew:

brew install gflags
Once the prerequisites are satisfied, you can build the command line tool with the command:

$ make grpc_cli
The main file can be found at https://github.com/grpc/grpc/blob/master/test/cpp/util/grpc_cli.cc

Prerequisites
Most grpc_cli commands need the server to support server reflection. See guides for Java , C++ and Go

Usage
List services
grpc_cli ls command lists services and methods exposed at a given port

List all the services exposed at a given port

$ grpc_cli ls localhost:50051
output:

helloworld.Greeter
grpc.reflection.v1alpha.ServerReflection
The localhost:50051 part indicates the server you are connecting to.

List one service with details

grpc_cli ls command inspects a service given its full name (in the format of <package>.<service>). It can print information with a long listing format when -l flag is set. This flag can be used to get more details about a service.

$ grpc_cli ls localhost:50051 helloworld.Greeter -l
helloworld.Greeter is full name of the service.

output:

filename: helloworld.proto
package: helloworld;
service Greeter {
  rpc SayHello(helloworld.HelloRequest) returns (helloworld.HelloReply) {}
}
List methods
List one method with details

grpc_cli ls command also inspects a method given its full name (in the format of <package>.<service>.<method>).

$ grpc_cli ls localhost:50051 helloworld.Greeter.SayHello -l
helloworld.Greeter.SayHello is full name of the method.

output:

rpc SayHello(helloworld.HelloRequest) returns (helloworld.HelloReply) {}
Inspect message types
We can use grpc_cli type command to inspect request/response types given the full name of the type (in the format of <package>.<type>).

Get information about the request type

$ grpc_cli type localhost:50051 helloworld.HelloRequest
helloworld.HelloRequest is the full name of the request type.

output:

message HelloRequest {
  optional string name = 1;
}
Call a remote method
We can send RPCs to a server and get responses using grpc_cli call command.

Call a unary method Send a rpc to a helloworld server at localhost:50051:

$ grpc_cli call localhost:50051 SayHello "name: 'gRPC CLI'"
output: sh message: "Hello gRPC CLI"

SayHello is (part of) the gRPC method string. Then "name: 'world'" is the text format of the request proto message. For information on more flags, look at the comments of grpc_cli.cc.

Use local proto files

If the server does not have the server reflection service, you will need to provide local proto files containing the service definition. The tool will try to find request/response types from them.

$ grpc_cli call localhost:50051 SayHello "name: 'world'" \
  --protofiles=examples/protos/helloworld.proto
If the proto file is not under the current directory, you can use --proto_path to specify a new search root.

Send non-proto rpc

For using gRPC with protocols other than protobuf, you will need the exact method name string and a file containing the raw bytes to be sent on the wire.

$ grpc_cli call localhost:50051 /helloworld.Greeter/SayHello \
  --input_binary_file=input.bin \
  --output_binary_file=output.bin
On success, you will need to read or decode the response from the output.bin file.