package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"petmate/router"
	"time"
)

func main() {
	g := gin.New()

	router.Load(g, nil)

	/*	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Print("The router has been deployed successfully.")
	}()*/

	log.Printf("Start to listening the incoming requests on http address: %s", ":8080")
	if err := http.ListenAndServe(":8080", g); err != nil {
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
