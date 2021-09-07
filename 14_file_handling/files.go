package main

import (
	"fmt"
	"bufio"
	"io/ioutil"
	"os"
) 

func createFile(filename, text string) {

	fmt.Println("Create file and write to it!")
	
	file, err := os.Create(filename)
	
	if err != nil {
		fmt.Printf("Can't create file %s. Reason %s\n",filename,err)
		os.Exit(1)
	}

	defer file.Close()
	
	len, err := file.WriteString(text)
	if err != nil {
		fmt.Printf("Can't write to file %s. Reason %s\n",filename,err)
		os.Exit(1)
	}
	
	fmt.Printf("File name:%s\n",file.Name())
	fmt.Printf("Length:%d bytes\n",len)
}

func readFile(filename string) {

	fmt.Println("Reading a file\n")
	
	data, err := ioutil.ReadFile(filename)
	
	if err != nil {
		fmt.Printf("Can't read from file %s. Reason %s\n",filename,err)
		os.Exit(1)
	}
	fmt.Printf("File name:%s\n",filename)
	fmt.Printf("Size:%d bytes\n",len(data))
	fmt.Printf("Data:%s\n",data)
	
}

func main() {

	var filename string
	
	//user input for filename
	fmt.Println("Enter filename:")
	fmt.Scanln(&filename)
	
	//user input for content
	fmt.Println("Enter text:")
	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')
	
	//create file and read it
	createFile(filename,input)
	readFile(filename)
}