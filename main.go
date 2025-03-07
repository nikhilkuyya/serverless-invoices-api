package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/nikhilkuyya/invoice-go-app/internal/app"
	"github.com/nikhilkuyya/invoice-go-app/internal/routes"
)

func main() {
	var port int
	flag.IntVar(&port,"port", 8080,"default port")
	flag.Parse()

	appInstance, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	handler := routes.SetupRoutes(appInstance);
	server := http.Server{
		Addr: fmt.Sprintf(":%d",port),
		Handler: handler,
		IdleTimeout: time.Minute,
		ReadTimeout: time.Second * 10,
		WriteTimeout: time.Second * 30,
	}
	appInstance.Logger.Printf("we are running the app in port %d",port)

	err = server.ListenAndServe()
	if err != nil {
		appInstance.Logger.Fatal(err)
	}
}
