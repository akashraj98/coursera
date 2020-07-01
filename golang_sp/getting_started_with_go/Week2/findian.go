package main

import (
	"fmt"
	"strings"
	)

func main(){
	var value string
	fmt.Printf("Enter your String: ")
	fmt.Scan(&value)
	value = strings.ToLower(value)
	value = strings.ReplaceAll(value," ","")
	start_i := strings.HasPrefix(value,"i")
	contain_a := strings.Contains(value,"a")
	end_n := strings.HasSuffix(value,"n")
	if (start_i && (contain_a && end_n) ){
		fmt.Printf("Found!\n")
	}else{
		fmt.Printf("Not Found!\n")
	}
}
