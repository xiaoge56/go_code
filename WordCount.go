package main

import (
	"fmt"
	
	"strings"
)

func WordCount(s string) map[string]int {
	returnmap:=map[string]int{}
	temp:=strings.Fields(s)
	//fmt.Printf("Fields are: %q", strings.Fields(s))
	for _,v:=range temp{
		_,ok:=returnmap[v]
		if ok==true{
		returnmap[v]+=1
		}else{
			returnmap[v]=1
		}
	}
	fmt.Println(returnmap)
	return returnmap
}

func main() {
	WordCount("wo xi huan nui")
}
