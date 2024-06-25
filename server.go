package main

import (
	"fmt"
	"main/helper"
	"net/http"
	"path"
	"text/template"
)

func handlerIndex(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Welcome"))
}

func handlerPage(w http.ResponseWriter, r *http.Request){
	filepath := path.Join("views","index.html")

	tmpl,err:= template.ParseFiles(filepath)

	if err!= nil {
		helper.HttpErrorHandler(w,err)
    }

	data:= map[string]interface{}{
		"title":"Hello",
        "content":"Hello, world!",
	}

	err = tmpl.Execute(w,data)
	if err!= nil {
       helper.HttpErrorHandler(w,err)
    }
}



func handlerHello(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello, world!"))
}


func main() {

	http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/",handlerPage)
	http.HandleFunc("/hello",handlerHello)
	address := "localhost:9000"
	fmt.Printf("Server start at %s\n",address)

	err:= http.ListenAndServe(address,nil)
	if err != nil{
		helper.ErrorHelper(err)
	}
}