package main

import (
	"flag"
	"os"
	xtrafor "xtraFor/xtraFor"
)

var expectingByte byte
var contents byte[]

func iteration(i int) {
	switch decode(contents[i]) {
		
	}
}

func main() {
	flag.Parse()
	tg := flag.Arg(0)
	contents, err := os.ReadFile(tg)
	floop := xtrafor.Loop(0, false, iteration)
	if err != nil {
		panic(err)
	}
	floop.Step()
}

func decode(b byte) (int, int) {
	switch b {
	case 0x80:
		return 1, 1
	case 0x1f:
		return 2, 1
	case 0xee:
		return 3, 1
	case 0xff:
		return 4, 2
	case 0x01:
		return 5, 2
	case 0x02:
		return 6, 2
	case 0x03:
		return 10, 2
	case 0xfa:
		return 7, 3
	case 0xfb:
		return 8, 3
	case 0xfc:
		return 9, 3
	default:
		return 0, 0
	}
}
