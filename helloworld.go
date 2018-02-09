package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"log"
)

func Index (w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	w.Write( []byte ("Welcome!\n"))
} 

func Hello (w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	w.Write([]byte ("Hello " + ps.ByName("name") + " ,you are using this path " + r.URL.Path + "\n"))
} 

func Bye (w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	w.Write( []byte ("Bye bye!\n"))
} 

func main (){
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	router.GET("/bye", Bye)

	log.Fatal(http.ListenAndServe(":8090", router))
} 