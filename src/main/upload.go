package main

import (
	"net/http"
	"log"
)

/**
 *上传文件
 */
import (
	"html/template"
	"io/ioutil"
)

var uploadTemplate = template.Must(template.ParseFiles("index.html"))

func indexHandle(w http.ResponseWriter, r *http.Request) {
	if err := uploadTemplate.Execute(w, nil); err != nil {
		log.Fatal("Execute: ", err.Error())
		return
	}
}

func uploadHandle(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		log.Fatal("FormFile: ", err.Error())
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal("Close: ", err.Error())
			return
		}
	}()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("ReadAll: ", err.Error())
		return
	}

	w.Write(bytes)
}

func main() {
	http.HandleFunc("/", indexHandle)
	http.HandleFunc("/upload", uploadHandle)
	//http://127.0.0.1:8899/
	http.ListenAndServe(":8899", nil)
}