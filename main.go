package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	flag.Parse()
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// var n uint64 = 0x89d30e0ce6154c37
	// fmt.Println(n)
	// fmt.Println(int64(n))
	//
	// var n2 int64 = -8515447021864924105
	// fmt.Println(n2)
	// fmt.Println(uint64(n2))
}

func run() error {
	if len(flag.Args()) == 0 {
		return errors.New("require 1 argument")
	}
	n, err := strconv.ParseInt(flag.Args()[0], 0, 64)
	if err != nil {
		nuint, err := strconv.ParseUint(flag.Args()[0], 0, 64)
		if err != nil {
			return err
		}
		n = int64(nuint)
	}
	var (
		nUint64 = uint64(n)
		nInt64  = int64(n)
	)
	fmt.Fprintf(os.Stdout, "uint64:\t%d\n", nUint64)
	fmt.Fprintf(os.Stdout, "int64:\t%d\n", nInt64)
	fmt.Fprintf(os.Stdout, "uint64-hex:\t%#08x\n", nUint64)
	fmt.Fprintf(os.Stdout, "int64-hex:\t%#08x\n", nInt64)
	return nil
}
