package main

import (
	// "flag"
	"fmt"
	// "io"
	"log"
	"net/http"
	// "os"
	// "path"
	// "strconv"
)

var dir string
var port int
var staticHandler http.Handler

// 初始化参数
func init() {
	// dir = path.Dir(os.Args[0])
	// flag.IntVar(&port, "port", 800, "服务器端口")
	// flag.Parse()
	// fmt.Println("dir:", http.Dir(dir))
	staticHandler = http.FileServer(http.Dir("/tmp/holmes"))
	// staticHandler = staticHandler.ServeHTTP(w, req)
}

func main() {
	http.Handle("/holmes/download/*", http.StripPrefix("/holmes/download", staticHandler))
	http.HandleFunc("/", StaticServer)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// 静态文件处理
func StaticServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("path:" + req.URL.Path)
	// if req.URL.Path != "/down/" {
	staticHandler.ServeHTTP(w, req)
	// return
	// }

	// io.WriteString(w, "hello, world!\n")
}
