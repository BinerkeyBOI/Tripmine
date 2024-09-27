package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Libraries have a limit of two in depth.")

	platform := flag.String("p", "windows", "Choose which os to build the project with.")
	flag.Parse()

	project := flag.Arg(0)
	tpmfile, err := os.Create("build.tpm")
	if err != nil {
		panic(err)
	}
	defer tpmfile.Close()

	makeEverythingElse(project, *tpmfile)
	makeMain(project, *tpmfile, *platform)
}

func makeEverythingElse(p string, t os.File) {
	f, err := os.ReadDir(p)
	if err != nil {
		panic(err)
	}
	for i := range f {
		entry := f[i]
		if !(entry.IsDir()) {
			if !(entry.Name() == "main.bat" || entry.Name() == "main.sh") {
				makeFile(p+"/"+entry.Name(), t)
				continue
			}
			continue
		}
		makeFolder(p+"/"+entry.Name(), t)
		continue
	}
}

func makeFolder(p string, t os.File) {
	pa := strings.Split(p, "/")
	folder := pa[len(pa)-1]
	cReturn := []byte{0x1f, 0xff}
	cReturn = append(cReturn, []byte(folder)...)
	cReturn = append(cReturn, 0xff, 0x02)

	listr, err := os.ReadDir(p)
	if err != nil {
		panic(err)
	}
	for i := range listr {
		entry := listr[i]
		if !(entry.IsDir()) {
			if !(entry.Name() == "main.bat" || entry.Name() == "main.sh") {
				makeFile(entry.Name(), t)
				continue
			}
			continue
		}
		makeFolder(entry.Name(), t)
		continue
	}

	cReturn = append(cReturn, 0x02, 0xfb)
	_, err = t.Write(cReturn)
	if err != nil {
		panic(err)
	}
}

func makeFile(p string, t os.File) {
	file, err := os.Open(p)
	if err != nil {
		panic(err)
	}
	rr := strings.Split(file.Name(), "/")
	name := []byte(rr[len(rr)-1])
	defer file.Close()

	contents, err := os.ReadFile(p)
	cReturn := []byte{0x80, 0xff}
	if err != nil {
		panic(err)
	}

	cReturn = append(cReturn, name...)
	cReturn = append(cReturn, 0xff, 0x01)
	for i := range contents {
		contents[i]--
	}
	for i := range contents {
		cReturn = append(cReturn, contents[i])
	}
	cReturn = append(cReturn, 0x01, 0xfa)

	_, err = t.Write(cReturn)
	if err != nil {
		panic(err)
	}
}

func makeMain(p string, t os.File, plat string) {
	switch plat {
	case "windows":
		main, err := os.ReadFile(p + "/main.bat")
		if err != nil {
			panic(err)
		}

		content := []byte{0xee}
		for i := range main {
			main[i]--
		}
		content = append(content, 0x03)
		for i := range main {
			content = append(content, main[i])
		}
		content = append(content, 0x03, 0xfc)

		_, err = t.Write(content)
		if err != nil {
			panic(err)
		}

	case "linux", "mac":
		main, err := os.ReadFile(p + "/main.sh")
		if err != nil {
			panic(err)
		}

		content := []byte{0xee}
		for i := range main {
			main[i]--
		}
		content = append(content, 0x03)
		for i := range main {
			content = append(content, main[i])
		}
		content = append(content, 0x03, 0xfc)

		_, err = t.Write(content)
		if err != nil {
			panic(err)
		}
	}
}
