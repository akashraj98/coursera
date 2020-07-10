package main

import (
	"fmt"
	"bufio"
	"os"
	"encoding/json"
)

func main(){
idMap := map[string]string{
	"Name":"",
	"Address":"",
	}
fmt.Printf("Enter your name: ")
reader := bufio.NewReader(os.Stdin)
name,_ := reader.ReadString('\n')
idMap["Name"] = name

fmt.Printf("Enter your Address: ")
add,_ := reader.ReadString('\n')
idMap["Address"]= add

//var jsonData []byte
jsonData,_:= json.Marshal(idMap)
fmt.Println(string(jsonData))
}
