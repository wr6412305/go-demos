package main

import (
	"context"
	hello "go-demos/micro-demos/greeter/srv/proto"
	"log"

	"github.com/emicklei/go-restful"
	"github.com/micro/go-grpc/client"
	"github.com/micro/go-web"
)

// Say ...
type Say struct{}

var (
	cl hello.SayService
)

// Anything ...
func (s *Say) Anything(req *restful.Request, rsp *restful.Response) {
	log.Print("Received Say.Anything API request")
	rsp.WriteEntity(map[string]string{
		"message": "Hi, this is the Greeter API",
	})
}

// Hello ...
func (s *Say) Hello(req *restful.Request, rsp *restful.Response) {
	log.Print("Received Say.Hello API request")

	name := req.PathParameter("name")

	response, err := cl.Hello(context.TODO(), &hello.Request{
		Name: name,
	})

	if err != nil {
		rsp.WriteError(500, err)
	}

	rsp.WriteEntity(response)
}

func main() {
	// Create service
	service := web.NewService(
		web.Name("go.micro.api.greeter"),
	)

	service.Init()

	// setup Greeter Server Client
	cl = hello.NewSayService("go.micro.srv.greeter", client.DefaultClient)

	// Create RESTful handler
	say := new(Say)
	ws := new(restful.WebService)
	wc := restful.NewContainer()
	ws.Consumes(restful.MIME_XML, restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON, restful.MIME_XML)
	ws.Path("/greeter")
	ws.Route(ws.GET("/").To(say.Anything))
	ws.Route(ws.GET("/{name}").To(say.Hello))
	wc.Add(ws)

	// Register Handler
	service.Handle("/", wc)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
