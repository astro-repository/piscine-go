package main

import (
	"os"
	"fmt"
	"ioutil"
)

func main(){
	dat, err := ioutil.ReadFile(os.Args[0])
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("test")
	}
}
