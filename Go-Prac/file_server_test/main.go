package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main(){
//	f, err := fileserver.New(fileserver.Config{
//		Host:     "http://squirrel-dev.paradise-soft.com.tw:12345",
//		Username: "arch",
//		Password: "1qaz!QAZ",
//	})
//	if err != nil{
//		log.Println(err)
//	}
//
//	ok, err := f.Delete("/zqb/app/deposit/18.png")
//	log.Print(ok)
//f.
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

	//filebuffer := bufio.NewReader

	log.Print("resss:  ", strings.Replace(strings.Trim(fmt.Sprint(bytes), "[]"), " ", ", ", -1))



	//e := f.Upload("/zqb/test", "teees.png", )
	//if e != nil {
	//	log.Println(e)
	//}


}