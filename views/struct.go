package views

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type FileStruct struct {
	File string `json:"file"`
}

type DataBase struct {
	Db []FileStruct
}

func (d *DataBase) UplData(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.Body)
	r.ParseMultipartForm(10 << 20)
	// fmt.Println(r.FormFile("files"))
	file, handler, err := r.FormFile("files")
	fmt.Println(handler.Filename)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer file.Close()
	temp, err := ioutil.TempFile("temp", "*"+handler.Filename)
	if err != nil {
		panic(fmt.Sprintf("Error:PANIC CALLED", err))
	}
	defer temp.Close()
	fp, err := ioutil.ReadAll(file)
	temp.Write(fp)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("File Uploaded..")
	}

	// return 1
}
