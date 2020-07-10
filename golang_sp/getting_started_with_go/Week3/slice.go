package main
import (
	"fmt"
	"sort"
	"strings"
	"strconv"
)

func main(){
num := ""

sli := make([]int ,0, 3)

for true{
	fmt.Printf("Enter number to be added to slice : ")
	fmt.Scan(&num)
	if strings.ToLower(num) == "x" { break }
	x,err := strconv.Atoi(num)
	if err !=nil { continue}
	sli =append(sli,x)
	sort.Ints(sli)
	fmt.Println("Sorted Slice: ",sli)
	}
}
