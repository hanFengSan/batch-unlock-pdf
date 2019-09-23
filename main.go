package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

func readFiles() []string {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		fmt.Println(err)
	}
	result := []string{}
	for _, f := range files {
		if strings.Contains(f.Name(), ".pdf") || strings.Contains(f.Name(), ".PDF") {
			result = append(result, f.Name())
		}
	}
	return result
}

func getPwd() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter user password: ")
	pwd, _ := reader.ReadString('\n')
	pwd = strings.Trim(pwd, "\n")
	pwd = strings.Trim(pwd, "\r")
	return pwd
}

func wait() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("please enter any key to exit...")
	reader.ReadString('\n')
}

func mkdir() {
	if _, err := os.Stat("./output"); os.IsNotExist(err) {
		os.MkdirAll("./output", 0777)
	}
}

func main() {
	pwd := getPwd()
	succeedFileNum := 0
	files := readFiles()
	mkdir()
	for index, name := range files {
		conf := pdfcpu.NewAESConfiguration("", pwd, 256)
		i := strconv.Itoa(index + 1)
		if err := api.DecryptFile(name, "./output/"+name, conf); err == nil {
			fmt.Println(i + ". " + name + ": Done")
			succeedFileNum++
		} else {
			fmt.Println(err)
			fmt.Println(i + ". " + name + ": Failed!!!")
		}
	}
	fmt.Println(strconv.Itoa(succeedFileNum) + "/" + strconv.Itoa(len(files)) + " done")
	wait()
}
