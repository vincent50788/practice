package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"vincent/practice/Go-Prac/file_server_test/fileserver"
)

func main(){
	_, err := fileserver.New(fileserver.Config{
		Host:     "http://squirrel-dev.paradise-soft.com.tw:12345",
		Username: "arch",
		Password: "1qaz!QAZ",
	})
	if err != nil{
		log.Println(err)
	}

	fileToBeUploaded := "ALIPAY.png"
	file, err := os.Open(fileToBeUploaded)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()
	fileInfo, _ := file.Stat()
	log.Print("fileName: ", fileInfo.Name())
	var size int64 = fileInfo.Size()
	bytes := make([]byte, size)

	// read file into bytes
	buffer := bufio.NewReader(file)
	_, er := buffer.Read(bytes)
	if er != nil {
		log.Print(er)
	}
	tmp := string(bytes)
	strings.Replace(strings.Trim(fmt.Sprint(tmp), "[]"), " ", ", ", -1)

	log.Print("image", tmp)

	//filebuffer := bufio.NewReader

	//a, err := f.ZQBUpload("tutorials", "test0102", bytes)
	//log.Print(a)



	//e := f.Upload("/zqb/test", "teees.png", )
	//if e != nil {
	//	log.Println(e)
	//}


}