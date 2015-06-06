package main

import(
	"fmt"
	"net/http"
)


func handler(writer http.ResponseWriter, r *http.Request){
	fmt.Fprintf(writer, "hello world!")
}

func main(){
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
