package main

import (
	"net/http"
	"fmt"
	"log"
)

func main()  {

	//采集agent信息
	http.HandleFunc("/",HomePage)
	http.HandleFunc("/monitor/",MonitorAgent)
	log.Fatal(http.ListenAndServe(":8081", nil))
}


func HomePage(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintf(w,"index")
}

func MonitorAgent(w http.ResponseWriter,r *http.Request)  {

	fmt.Fprintf(w,"moniting")
}