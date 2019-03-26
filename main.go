package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
	"petmate/config"
	_ "petmate/config"
	"petmate/model"
	"petmate/router"
	"time"
)

func init()  {
	log.Print("main init start...")
	config.InitFileConf()
	model.InitRedisClient(viper.GetString("redis.addr"), viper.GetString("password"), viper.GetInt("poolsize"))
	log.Print("main init end...")
}

func main() {
	g := gin.New()

	router.Load(g, nil)

	log.Printf("app [%s] listening at [%s]", viper.GetString("app.name"), viper.GetString("app.addr"))
	if err := http.ListenAndServe(viper.GetString("app.addr"), g); err != nil {
		log.Printf("api server start fail, cause by: %s", err)
	}
}

func pingServer() error {
	for i := 0; i < 2; i++ {
		if resp, err := http.Get("http://127.0.0.1:8080" + "/sysinfo/ping"); resp.StatusCode == 200 && err == nil {

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
