// BlogGo project main.go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/chownplusx/web"
	"io/ioutil"
)

type ServerConfig struct {
	Port string
	Fcgi string
}

func main() {
	configtext, _err := ioutil.ReadFile("config.json")
	if _err != nil {
		fmt.Println(_err)
	}
	var conf ServerConfig
	err := json.Unmarshal(configtext, &conf)
	if err != nil {
		fmt.Println(err)
	}
	//WEB ROUTES
	web.Get("/()", IndexLoadGet)
	web.Get("/static/(.*)", Sendstatic)
	web.Get("/post/(.*)", GetSinglePost)
	//STARTING PROCEDURE
	if conf.Fcgi == "false" {
		fmt.Println("FCGI:", conf.Fcgi)
		fmt.Println("PORT:", conf.Port)
		web.Run("0.0.0.0:" + conf.Port)
	} else {
		fmt.Println("FCGI:", conf.Fcgi)
		fmt.Println("PORT:", conf.Port)
		web.RunFcgi("0.0.0.0:" + conf.Port)
	}
}
