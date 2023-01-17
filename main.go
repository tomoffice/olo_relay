package main

import (
	"fmt"
	"net/http"
	"relay/handle"
	PkgJson "relay/jsonwrap"
)

func main() {
	config := &PkgJson.Config{}
	decoder := PkgJson.NewConfig(config)
	err := decoder.DecodeFile("./config.json")
	if err != nil {
		fmt.Println("Read config file error:", err)
		fmt.Scanln() // wait for Enter Key
	}
	var concernNames []string
	for _, v := range config.TargetNames {
		concernNames = append(concernNames, v.Key)
	}
	var handlefunc handle.Handler = handle.New()
	handlefunc.Config(config.HttpToken, concernNames...)
	postFunc := handlefunc.GetFunc(config.NotifyToken)
	http.HandleFunc(config.Pattern, postFunc)
	err = http.ListenAndServe(fmt.Sprintf("%s:%d", config.Ip, config.Port), nil)
	if err != nil {
		fmt.Println(err)
	}

}
