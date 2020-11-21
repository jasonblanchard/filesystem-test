package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"syscall"
	"time"
)

func main() {
	if _, err := os.Stat("./output.txt"); os.IsNotExist(err) {
		err := ioutil.WriteFile("./output.txt", []byte(""), 0644)
		if err != nil {
			panic(nil)
		}
	}

	// file, err := os.OpenFile("./output.txt", os.O_WRONLY, os.ModeAppend)
	// if err != nil {
	// 	panic(err)
	// }

	// for {
	// 	data := []byte("chello\n")
	// 	fmt.Println(fmt.Sprintf("%s: Printing data %s", time.Now(), data))
	// 	_, err := file.Write(data)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}

	// 	time.Sleep(time.Millisecond * 500)
	// }

	fd, err := syscall.Open("./output.txt", syscall.O_RDWR, 755)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fd)

	for {
		data := []byte("chello\n")
		fmt.Println(fmt.Sprintf("%s: Printing data %s", time.Now(), data))

		_, err := syscall.Write(fd, data)
		if err != nil {
			fmt.Println(err)
		}

		time.Sleep(time.Millisecond * 200)
	}
}
