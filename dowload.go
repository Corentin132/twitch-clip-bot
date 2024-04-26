package main

import (
	"fmt"
	"os"
)

func DowloadVideo(url string) {
	urlFile, err := os.Create("url.txt")
	if err != nil {
		fmt.Println("Can't writting file : ", err)
	}
	defer urlFile.Close()

	_, err2 := urlFile.WriteString(url)
	if err2 != nil {
		fmt.Println("Can't write on file : ", err2)
	}

}
