package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func webTest(w http.ResponseWriter, r *http.Request) {
	page, err := os.Open("index.html")
	if err != nil {
		log.Print(err)
	}


	fmt.Fprint(w, page)
	// fmt.Fprintln(w, "這是一個開始")
}

func myWeb(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //它還將請求主體解析爲表單，獲得POST Form表單數據，必須先調用這個函數

	for k, v := range r.URL.Query() {
		fmt.Println("key:", k, ", value:", v[0])
	}

	json.Marshal()

	//for k, v := range r.PostForm {
	//	fmt.Println("key:", k, ", value:", v[0])
	//}

	x := r.Method
	fmt.Fprintln(w, x)
	data := fmt.Sprintf("the request parem name is %s", r.URL.Query().Get("name"))

	fmt.Fprintln(w, data)
}

func pointer(w http.ResponseWriter, r *http.Request) {
	mystring := "hi"
	//取指針
	mypointer := &mystring
	//取值
	mystring2 := *mypointer

	fmt.Fprintf(w, mystring, mypointer, mystring2)
}

func main() {
	http.HandleFunc("/", myWeb)
	http.HandleFunc("/pointer", pointer)
	http.HandleFunc("/webTest", webTest)

	fmt.Println("服務器即將開啓，訪問地址 http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("服務器開啓錯誤: ", err)
	}
}
