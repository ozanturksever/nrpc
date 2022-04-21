// This code was autogenerated from helloworld.proto, do not edit.
package helloworld

import (
	"context"
	"log"
	"time"

	"google.golang.org/protobuf/proto"
	"github.com/nats-io/nats.go"
	"github.com/ozanturksever/nrpc"
)

// GreeterServer is the interface that providers of the service
// Greeter should implement.
type GreeterServer interface {
	SayHello(ctx context.Context, req HelloRequest) (resp HelloReply, err error)
}

// GreeterHandler provides a NATS subscription handler that can serve a
// subscription using a given GreeterServer implementation.
type GreeterHandler struct {
	ctx     context.Context
	workers *nrpc.WorkerPool
	nc      nrpc.NatsConn
	server  GreeterServer

	encodings []string
}

func NewGreeterHandler(ctx context.Context, nc nrpc.NatsConn, s GreeterServer) *GreeterHandler {
	return &GreeterHandler{
		ctx:    ctx,
		nc:     nc,
		server: s,

		encodings: []string{"protobuf"},
	}
}

func NewGreeterConcurrentHandler(workers *nrpc.WorkerPool, nc nrpc.NatsConn, s GreeterServer) *GreeterHandler {
	return &GreeterHandler{
		workers: workers,
		nc:      nc,
		server:  s,
	}
}

// SetEncodings sets the output encodings when using a '*Publish' function
func (h *GreeterHandler) SetEncodings(encodings []string) {
	h.encodings = encodings
}

func (h *GreeterHandler) Subject() string {
	return "Greeter.>"
}

func (h *GreeterHandler) Handler(msg *nats.Msg) {
	var ctx context.Context
	if h.workers != nil {
		ctx = h.workers.Context
	} else {
		ctx = h.ctx
	}
	request := nrpc.NewRequest(ctx, h.nc, msg.Subject, msg.Reply)
	// extract method name & encoding from subject
	_, _, name, tail, err := nrpc.ParseSubject(
		"", 0, "Greeter", 0, msg.Subject,
	)
	if err != nil {
		log.Printf("GreeterHanlder: Greeter subject parsing failed: %v", err)
		return
	}

	request.MethodName = name
	request.SubjectTail = tail

	// call handler and form response
	var immediateError *nrpc.Error
	switch name {
	case "SayHello":
		_, request.Encoding, err = nrpc.ParseSubjectTail(0, request.SubjectTail)
		if err != nil {
			log.Printf("SayHelloHanlder: SayHello subject parsing failed: %v", err)
			break
		}
		var req HelloRequest
		if err := nrpc.Unmarshal(request.Encoding, msg.Data, &req); err != nil {
			log.Printf("SayHelloHandler: SayHello request unmarshal failed: %v", err)
			immediateError = &nrpc.Error{
				Type:    nrpc.Error_CLIENT,
				Message: "bad request received: " + err.Error(),
			}
		} else {
			request.Handler = func(ctx context.Context) (proto.Message, error) {
				innerResp, err := h.server.SayHello(ctx, req)
				if err != nil {
					return nil, err
				}
				return &innerResp, err
			}
		}
	default:
		log.Printf("GreeterHandler: unknown name %q", name)
		immediateError = &nrpc.Error{
			Type:    nrpc.Error_CLIENT,
			Message: "unknown name: " + name,
		}
	}
	if immediateError == nil {
		if h.workers != nil {
			// Try queuing the request
			if err := h.workers.QueueRequest(request); err != nil {
				log.Printf("nrpc: Error queuing the request: %s", err)
			}
		} else {
			// Run the handler synchronously
			request.RunAndReply()
		}
	}

	if immediateError != nil {
		if err := request.SendReply(nil, immediateError); err != nil {
			log.Printf("GreeterHandler: Greeter handler failed to publish the response: %s", err)
		}
	} else {
	}
}

type GreeterClient struct {
	nc       nrpc.NatsConn
	Subject  string
	Encoding string
	Timeout  time.Duration
}

func NewGreeterClient(nc nrpc.NatsConn) *GreeterClient {
	return &GreeterClient{
		nc:       nc,
		Subject:  "Greeter",
		Encoding: "protobuf",
		Timeout:  5 * time.Second,
	}
}

func (c *GreeterClient) SayHello(req HelloRequest) (resp HelloReply, err error) {

	subject := c.Subject + "." + "SayHello"

	// call
	err = nrpc.Call(&req, &resp, c.nc, subject, c.Encoding, c.Timeout)
	if err != nil {
		return // already logged
	}

	return
}

type Client struct {
	nc              nrpc.NatsConn
	defaultEncoding string
	defaultTimeout  time.Duration
	Greeter         *GreeterClient
}

func NewClient(nc nrpc.NatsConn) *Client {
	c := Client{
		nc:              nc,
		defaultEncoding: "protobuf",
		defaultTimeout:  5 * time.Second,
	}
	c.Greeter = NewGreeterClient(nc)
	return &c
}

func (c *Client) SetEncoding(encoding string) {
	c.defaultEncoding = encoding
	if c.Greeter != nil {
		c.Greeter.Encoding = encoding
	}
}

func (c *Client) SetTimeout(t time.Duration) {
	c.defaultTimeout = t
	if c.Greeter != nil {
		c.Greeter.Timeout = t
	}
}