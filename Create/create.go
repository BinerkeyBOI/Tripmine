package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	pOS := flag.String("p", "windows", "Choose which os to create the project with.")
	flag.Parse()

	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	startCreation(pwd, *pOS)
}

func startCreation(dir string, f string) {
	err := os.Mkdir(flag.Arg(0), 0755)
	if err != nil {
		panic(err)
	}

	switch f {
	case "windows":
		mainV, err := os.Create(fmt.Sprintf("%v/%v/main.bat", dir, flag.Arg(0)))
		if err != nil {
			panic(err)
		}
		defer mainV.Close()

		mainV.Chmod(0755)
		mainV.Write([]byte("python3 main.py"))

	case "linux", "mac":
		mainV, err := os.Create(fmt.Sprintf("%v/%v/main.sh", dir, flag.Arg(0)))
		if err != nil {
			panic(err)
		}
		defer mainV.Close()

		mainV.Chmod(0755)
		mainV.Write([]byte("python3 main.py"))
	}

	pyMain, _ := os.Create(fmt.Sprintf("%v/main.py", flag.Arg(0)))
	pyMain.Write([]byte(`print("Hello, Tripmine!")`))
	pyMain.Close()
}
