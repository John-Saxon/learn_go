package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func Download(w http.ResponseWriter, r *http.Request) {
	file := "/tmp/holmes/test"
	// if exist := isExist(file); !exist {
	// 	http.NotFound(w, r)
	// }
	http.ServeFile(w, r, file)
}

func isExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.Methods("GET").
		Path("/test").
		Name("TEST").
		Handler(http.HandlerFunc(Download))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println(err)
	}
}
