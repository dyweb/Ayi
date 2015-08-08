package check

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func getv(sw string, vmap map[string]string) (version string, version2 string) {
	argv := []string{"-v", "-V", "-version", "--version"}
	list := []string{"php", "mysql", "java", "python"}
	vstr := "-v"
	for i, v := range argv {
		if list[i] == sw {
			vstr = v
		}
	}
	if true {
		cmd := exec.Command(sw, vstr)
		stdout, _ := cmd.StdoutPipe()
		stderr, _ := cmd.StderrPipe()
		if err := cmd.Start(); err != nil {
			//fmt.Println("Start: ", err.Error())
			//return
		}
		bytesErr, _ := ioutil.ReadAll(stderr)
		if len(bytesErr) != 0 {
			//fmt.Printf("%s", bytesErr)
			//return
		}
		bytes, _ := ioutil.ReadAll(stdout)
		if err := cmd.Wait(); err != nil {
			//fmt.Println("Wait: ", err.Error())
			//return
		}
		return string(bytes), string(bytesErr)
	}
	return "no", "error"
}

func WeaIns(sw string, vmap map[string]string) (f string, err error) {
	f, err = exec.LookPath(sw)
	fmt.Println(sw, "------------------------install check")
	if err != nil {
		fmt.Println("Not install", sw)
	} else {
		fmt.Println("Already installed")
		fmt.Println("PATH: ", f)

		fmt.Println(sw, "--------Version check")
		version, version2 := getv(sw, vmap)
		fmt.Println("VERSION INFO1: \n" + version)
		fmt.Println("VERSION INFO2: \n" + version2)
		if version == vmap[sw] {
			fmt.Println(sw, "'s version is right")
			fmt.Println(sw, "------------------------is working")
		} else {
			fmt.Println(sw, "'s version is not right, we need install version ", vmap[sw])
			fmt.Println(sw, "------------------------need fix")
		}
	}
	return f, err
}

func main() {
	var vmap map[string]string
	vmap = make(map[string]string)
	vmap["php"] = "5.5.9"
	vmap["mysql"] = "5"
	vmap["java"] = "7"
	vmap["python"] = "2.7.3"
	list := []string{"php", "mysql", "java", "python"}
	for _, sw := range list {
		fmt.Println("")
		WeaIns(sw, vmap)
		fmt.Println("")
	}
}
