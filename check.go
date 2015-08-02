package main

import (
	"fmt"
	"os/exec"
)

func WeaIns(sw string) (f string, err error){
	f, err = exec.LookPath(sw)
	if err != nil {
		fmt.Println("Not install", sw)
	}
	fmt.Println(f)
	return f, err
}

func main(){
	//cmd := exec.Command("echo", "Hello, world!");
	//cmd.Run();
	list := []string{"php", "mysql", "java", "python"}
	for  _, sw := range list {
		WeaIns(sw)
	}
}
