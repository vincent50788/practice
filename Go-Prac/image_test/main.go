package main

import (
	"bufio"
	"bytes"
	"fmt"
	"image"
	_"image/jpeg"
	_"image/png"
	_"image/gif"
	"log"
	"os"
	"strings"
)

func ImageFormat(imageData []byte) (format string) {
	_, format, err := image.Decode(bytes.NewReader(imageData))
	if err != nil {
		log.Print("err: ", err)
	}
	return format
}

func main(){
	file, err := os.Open("13.png")

	defer file.Close()
	if err != nil {
		log.Print(err)
	}

	fileInfo, _ := file.Stat()
	var size  = fileInfo.Size()
	bytes := make([]byte, size)
	buffer := bufio.NewReader(file)

	buffer.Read(bytes)
	res := strings.Replace(strings.Trim(fmt.Sprint(bytes), "[]"), " ", ", ", -1)
	log.Print("res", res, "!!!")
	result := ImageFormat(bytes)
	log.Print("result: ", result)
}
