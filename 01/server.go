package main

import "net/http"
type coaster struct{
	Name :=
}

func coasterHandlers(w http.ResponseWriter, r *http.Request){

}
func main() {
	http.HandleFunc("/coasters", coasterHandlers)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}
