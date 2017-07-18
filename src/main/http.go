package main

import (
	"net/http"
	"fmt"
	"strings"
	"log"
)

func httpSayHello(w http.ResponseWriter, r *http.Request)  {
	//解析参数, 默认是不会解析的
	r.ParseForm()
	// 这些信息是输出到服务器端的打印信息
	fmt.Println("request map:", r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ";"))
	}

	// 这个写入到w的信息是输出到客户端的
	fmt.Fprintf(w, "Hello danxingxi lizhi!\n")
}

func main()  {

	// 设置访问的路由
	http.HandleFunc("/lizhi/", httpSayHello)
	//->http://127.0.0.1:9090/lizhi/

	// 设置监听的端口
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}