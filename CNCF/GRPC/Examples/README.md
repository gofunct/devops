# GRPC & Protobuf Examples

***

- [GRPC & Protobuf Examples](#grpc--protobuf-examples)
- [Protobuf todo Example](#protobuf-todo-example)
- [Golang todo GRPC Server](#golang-todo-grpc-server)
- [Protobuf speech-text file](#protobuf-speech-text-file)
- [Protoc Makefile](#protoc-makefile)
- [GRPC/HTTP Server](#grpchttp-server)
		- [Packages](#packages)
		- [Variables](#variables)
		- [GRPC & HTTP Handler](#grpc--http-handler)
		- [Start GRPC Server](#start-grpc-server)
		- [Create a new GRPC server stub with credstore auth](#create-a-new-grpc-server-stub-with-credstore-auth)
- [GRPC- Keys](#grpc--keys)
		- [Packages](#packages)
		- [Load Keys](#load-keys)
- [GRPC- Zipkin](#grpc--zipkin)
		- [Packages](#packages)
		- [Variables](#variables)
		- [OpenTracing](#opentracing)
- [GRPC- Client](#grpc--client)
		- [Packages](#packages)
		- [Dial](#dial)

***

# Protobuf todo Example

<details><summary>show</summary>
<p>

```
syntax = "proto3";

package todo;

message Task {
    string text = 1;
    bool done = 2;
}

message TaskList {
    repeated Task tasks = 1;
}

message Text {
    string text = 1;
}

message Void {}

service Tasks {
    rpc List(Void) returns(TaskList) {}
    rpc Add(Text) returns(Task) {}
}

```

</p>
</details>

—-

# Golang todo GRPC Server

<details><summary>show</summary>
<p>

```
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"

	"github.com/campoy/justforfunc/31-grpc/todo"
	"github.com/gogo/protobuf/proto"
	"golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

func main() {
	srv := grpc.NewServer()
	var tasks taskServer
	todo.RegisterTasksServer(srv, tasks)
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("could not listen to :8888: %v", err)
	}
	log.Fatal(srv.Serve(l))
}

type taskServer struct{}

type length int64

const (
	sizeOfLength = 8
	dbPath       = "mydb.pb"
)

var endianness = binary.LittleEndian

func (taskServer) Add(ctx context.Context, text *todo.Text) (*todo.Task, error) {
	task := &todo.Task{
		Text: text.Text,
		Done: false,
	}

	b, err := proto.Marshal(task)
	if err != nil {
		return nil, fmt.Errorf("could not encode task: %v", err)
	}

	f, err := os.OpenFile(dbPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("could not open %s: %v", dbPath, err)
	}

	if err := binary.Write(f, endianness, length(len(b))); err != nil {
		return nil, fmt.Errorf("could not encode length of message: %v", err)
	}
	_, err = f.Write(b)
	if err != nil {
		return nil, fmt.Errorf("could not write task to file: %v", err)
	}

	if err := f.Close(); err != nil {
		return nil, fmt.Errorf("could not close file %s: %v", dbPath, err)
	}
	return task, nil
}

func (taskServer) List(ctx context.Context, void *todo.Void) (*todo.TaskList, error) {
	b, err := ioutil.ReadFile(dbPath)
	if err != nil {
		return nil, fmt.Errorf("could not read %s: %v", dbPath, err)
	}

	var tasks todo.TaskList
	for {
		if len(b) == 0 {
			return &tasks, nil
		} else if len(b) < sizeOfLength {
			return nil, fmt.Errorf("remaining odd %d bytes, what to do?", len(b))
		}

		var l length
		if err := binary.Read(bytes.NewReader(b[:sizeOfLength]), endianness, &l); err != nil {
			return nil, fmt.Errorf("could not decode message length: %v", err)
		}
		b = b[sizeOfLength:]

		var task todo.Task
		if err := proto.Unmarshal(b[:l], &task); err != nil {
			return nil, fmt.Errorf("could not read task: %v", err)
		}
		b = b[l:]
		tasks.Tasks = append(tasks.Tasks, &task)
	}
}

```

</p>
</details>

***

# Protobuf speech-text file

<details><summary>show</summary>
<p>

```
syntax = "proto3";

package say;

service TextToSpeech {
    rpc Say(Text) returns(Speech) {}
}

message Text {
    string text = 1;
}

message Speech {
    bytes audio = 1;
}

```

</p>
</details>

***

# Protoc Makefile

<details><summary>show</summary>
<p>

```
build:
	protoc -I . say.proto --go_out=plugins=grpc:.

```

</p>
</details>

***

# GRPC/HTTP Server

### Packages

<details><summary>show</summary>
<p>

package serverhelpers

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"

	"github.com/golang/glog"
	"github.com/google/credstore/client"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	opentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

</p>
</details>



### Variables

* ListenAddress
* serverCert
* serverKey
* clientCA
* credStoreAddress
* credstoreCA

<details><summary>show</summary>
<p>

var (
	// ListenAddress is the grpc listen address
	ListenAddress = flag.String("listen", "", "GRPC listen address")

	serverCert = flag.String("server-cert", "", "server TLS cert")
	serverKey  = flag.String("server-key", "", "server TLS key")
	clientCA   = flag.String("client-ca", "", "client CA")

	credStoreAddress = flag.String("credstore-address", "", "credstore grpc address")
	credStoreCA      = flag.String("credstore-ca", "", "credstore server ca")
)

</p>
</details>

### GRPC & HTTP Handler

* otherHandler.ServeHTTP(w,r) 
* grpc.Server.ServeHTTP(w,r)

<details><summary>show</summary>
<p>

func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	})
}

</p>
</details>

### Start GRPC Server

* https
* Uses either grpc or http handler

<details><summary>show</summary>
<p>
func ListenAndServe(grpcServer *grpc.Server, otherHandler http.Handler) error {
	lis, err := net.Listen("tcp", *ListenAddress)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	if *serverCert != "" {
		serverCertKeypair, err := tls.LoadX509KeyPair(*serverCert, *serverKey)
		if err != nil {
			return fmt.Errorf("failed to load server tls cert/key: %v", err)
		}

		var clientCertPool *x509.CertPool
		if *clientCA != "" {
			caCert, err := ioutil.ReadFile(*clientCA)
			if err != nil {
				return fmt.Errorf("failed to load client ca: %v", err)
			}
			clientCertPool = x509.NewCertPool()
			clientCertPool.AppendCertsFromPEM(caCert)
		}

		var h http.Handler
		if otherHandler == nil {
			h = grpcServer
		} else {
			h = grpcHandlerFunc(grpcServer, otherHandler)
		}

		httpsServer := &http.Server{
			Handler: h,
			TLSConfig: &tls.Config{
				Certificates: []tls.Certificate{serverCertKeypair},
				NextProtos:   []string{"h2"},
			},
		}

		if clientCertPool != nil {
			httpsServer.TLSConfig.ClientCAs = clientCertPool
			httpsServer.TLSConfig.ClientAuth = tls.RequireAndVerifyClientCert
		} else {
			glog.Warningf("no client ca provided for grpc server")
		}

		glog.Infof("serving on %v", *ListenAddress)
		err = httpsServer.Serve(tls.NewListener(lis, httpsServer.TLSConfig))
		return fmt.Errorf("failed to serve: %v", err)
	}

	glog.Warningf("serving INSECURE on %v", *ListenAddress)
	err = grpcServer.Serve(lis)
	return fmt.Errorf("failed to serve: %v", err)
}

</p>
</details>


### Create a new GRPC server stub with credstore auth

* register w/ prometheus
* store credentials

<details><summary>show</summary>
<p>

func NewServer() (*grpc.Server, *client.CredstoreClient, error) {
	var grpcServer *grpc.Server
	var cc *client.CredstoreClient

	if *credStoreAddress != "" {
		var err error
		cc, err = client.NewCredstoreClient(context.Background(), *credStoreAddress, *credStoreCA)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to init credstore: %v", err)
		}

		glog.Infof("enabled credstore auth")
		grpcServer = grpc.NewServer(
			grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
				otgrpc.OpenTracingServerInterceptor(opentracing.GlobalTracer()),
				grpc_prometheus.UnaryServerInterceptor,
				client.CredStoreTokenInterceptor(cc.SigningKey()),
				client.CredStoreMethodAuthInterceptor(),
			)))
	} else {
		grpcServer = grpc.NewServer(
			grpc.UnaryInterceptor(
				otgrpc.OpenTracingServerInterceptor(opentracing.GlobalTracer())))
	}

	reflection.Register(grpcServer)
	grpc_prometheus.Register(grpcServer)

	return grpcServer, cc, nil
}

</p>
</details>

# GRPC- Keys

### Packages

<details><summary>show</summary>
<p>

package pki

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

</p>
</details>

### Load Keys

<details><summary>show</summary>
<p>

func LoadECKeyFromFile(fileName string) (*ecdsa.PrivateKey, error) {
	privateKeyBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read signing key file: %v", err)
	}

	privateKeyPEM, _ := pem.Decode(privateKeyBytes)
	if privateKeyPEM == nil {
		return nil, fmt.Errorf("failed to decode pem signing key file: %v", err)
	}

	privateKey, err := x509.ParseECPrivateKey(privateKeyPEM.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse signing key file: %v", err)
	}

	return privateKey, nil
}

</p>
</details>

# GRPC- Zipkin

### Packages

<details><summary>show</summary>
<p>


package tracing

import (
	"flag"
	"fmt"

	opentracing "github.com/opentracing/opentracing-go"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
)

</p>
</details>

### Variables

<details><summary>show</summary>
<p>

var (
	zipkinURL = flag.String("zipkin-url", "http://localhost:9411/api/v1/spans", "zipkin url for distributed tracing")
)
</p>
</details>

### OpenTracing

<details><summary>show</summary>
<p>

func InitTracer(hostPort string, serviceName string) error {
	if *zipkinURL == "" {
		return nil
	}

	collector, err := zipkin.NewHTTPCollector(*zipkinURL)
	if err != nil {
		return fmt.Errorf("unable to create Zipkin HTTP collector: %v", err)
	}
	recorder := zipkin.NewRecorder(collector, false, hostPort, serviceName)
	tracer, err := zipkin.NewTracer(
		recorder,
		zipkin.ClientServerSameSpan(false),
		zipkin.TraceID128Bit(true),
	)
	if err != nil {
		return fmt.Errorf("unable to create Zipkin tracer: %v", err)
	}
	opentracing.InitGlobalTracer(tracer)

	return nil
}

</p>
</details>

# GRPC- Client

### Packages

<details><summary>show</summary>
<p>

package clienthelpers

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	opentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

</p>
</details>

### Dial

<details><summary>show</summary>
<p>

// NewGRPCConn is a helper wrapper around grpc.Dial.
func NewGRPCConn(
	address string,
	serverCAFileName string,
	clientCertFileName string,
	clientKeyFileName string,
) (*grpc.ClientConn, error) {
	if serverCAFileName == "" {
		return grpc.Dial(address,
			grpc.WithInsecure(),
			grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())))
	}

	caCert, err := ioutil.ReadFile(serverCAFileName)
	if err != nil {
		return nil, err
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	cfg := &tls.Config{
		RootCAs: caCertPool,
	}

	if clientCertFileName != "" && clientKeyFileName != "" {
		peerCert, err := tls.LoadX509KeyPair(clientCertFileName, clientKeyFileName)
		if err != nil {
			return nil, err
		}
		cfg.Certificates = []tls.Certificate{peerCert}
	}

	return grpc.Dial(address,
		grpc.WithTransportCredentials(credentials.NewTLS(cfg)),
		grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())),
	)
}

</p>
</details>


