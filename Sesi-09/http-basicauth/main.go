package main

import (
	"net/http"
	"fmt"
	"encoding/json"
)

func main() {
	http.HandleFunc("/student", ActionStudent)

	server := new(http.Server)
	server.Addr = ":9000"

	fmt.Println("Server started at localhost:9000")
	server.ListenAndServe()
}


func ActionStudent(w http.ResponseWriter, r *http.Request) {
	if !Auth(w, r)			{return}
	if !AllowOnlyGet(w, r)	{return}

	if id := r.URL.Query().Get("id"); id != "" {
		OutputJSON(w, SelectStudent(id))
		return
	}

	OutputJSON(w, GetStudent())
}

func OutputJSON(w http.ResponseWriter, r *http.Request) {
	res, err := json.Marshal(o)
	if err != nil{
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

