package main

import (
	"context"
	"encoding/json"
	grpc_gateway_demo "grpc-gateway-demo"
	"grpc-gateway-demo/api/demopb"
	"grpc-gateway-demo/internal/handler"
	"grpc-gateway-demo/third_party/swagger_ui"
	"html/template"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc/credentials/insecure"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":9090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer func() {
		_ = conn.Close()
	}()
	client := demopb.NewDemoClient(conn)

	demoHandler := handler.NewDemoHandler(client)
	serve(":8080", ":8081", demoHandler)
}

func serve(grpcPort, gatewayPort string, srv demopb.DemoServer) {
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	demopb.RegisterDemoServer(s, srv)
	log.Printf("Serve Grpc on %s", grpcPort)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	gwMux := runtime.NewServeMux()
	mux := http.NewServeMux()
	mux.Handle("/", gwMux)
	serveSwagger(mux)

	err = demopb.RegisterDemoHandlerFromEndpoint(context.Background(),
		gwMux,
		grpcPort,
		[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		log.Fatal(err)
	}
	gwServer := &http.Server{
		Addr:    gatewayPort,
		Handler: mux,
	}
	log.Printf("Serve Gateway on %s", gatewayPort)
	if err = gwServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func serveSwagger(mux *http.ServeMux) {
	fs := grpc_gateway_demo.SwaggerJsonFS

	mux.Handle("/doc/", http.StripPrefix("/doc/", http.FileServer(http.FS(fs))))

	mux.HandleFunc("/swagger/index.html", func(writer http.ResponseWriter, request *http.Request) {
		t1, err := template.ParseFS(swagger_ui.SwaggerUriFS, "index.tpl")
		if err != nil {
			log.Fatal(err)
		}

		_, err = fs.Open("swagger/demo.swagger.json")
		if err != nil {
			log.Fatal(err)
		}

		type Site struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		}
		sites := []*Site{
			&Site{
				Name: "demo",
				URL:  "/doc/swagger/demo.swagger.json",
			}}
		siteByte, _ := json.Marshal(sites)
		err = t1.Execute(writer, string(siteByte))
		if err != nil {
			log.Fatal(err)
		}
	})

	mux.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(http.FS(swagger_ui.SwaggerUriFS))))
}

func swaggerFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "")
}
