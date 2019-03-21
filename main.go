package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
	_ "petmate/config"
	"petmate/router"
	"time"
)

func main() {
	g := gin.New()

	router.Load(g, nil)

	log.Printf("[%s] Start to listening the incoming requests on http address: [%s]", viper.GetString("app.name"), viper.GetString("app.addr"))
	if err := http.ListenAndServe(viper.GetString("app.addr"), g); err != nil {
		log.Printf("api server start fail, cause by: %s", err)
	}
}

func pingServer() error {
	for i := 0; i < 2; i++ {
		if resp, err := http.Get("http://127.0.0.1:8080" + "/sys/ping"); resp.StatusCode == 200 && err == nil {

			defer resp.Body.Close()
			if body, err := ioutil.ReadAll(resp.Body); err == nil && body != nil {
				log.Print("ping router success, resp body:" + string(body))
			}
			return nil
		}
		log.Print("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second * 1)
	}

	return errors.New("Cannot connect to the router.")
}
