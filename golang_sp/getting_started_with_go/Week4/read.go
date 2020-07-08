package main

import(
	"fmt"
	"io/ioutil"
	"strings"
)

type person struct{
	fname string
	lname string
}

func main(){
	sli := make([]person,0,3)
    fmt.Printf("Enter the name of file: ")
    var filename string
    fmt.Scan(&filename)
    dat,err := ioutil.ReadFile("./"+filename)
    if err != nil{ panic(err)}
	data := strings.Split(string(dat),"\n")
	data = data[:len(data)-1]
	for _,v := range data{
		var p person
		p.fname = strings.Split(v," ")[0]
		p.lname = strings.Split(v," ")[1]
		sli= append(sli,p)
		}
	for _,v:= range sli{
		fmt.Println(v)
		}
    }
