package main

import (
	"net/http"
	"os"
	"server/api"
	"server/config"
)

func main() {
	Config, err := config.NewConfig("server", "json", ".")
	if err != nil {
		os.Exit(1)
	}
	println(Config.Cf.Get("apple").(string), "sdf")

	server := api.NewServer()
	http.ListenAndServe(":8080", server)
}
