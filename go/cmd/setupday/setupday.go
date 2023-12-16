package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strconv"
)

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		panic("No input date specified")
	}
	day, err := strconv.Atoi(flag.Args()[0])
	if err != nil {
		panic(err)
	}
	dir := fmt.Sprintf("dec%02d", day)
	err = os.MkdirAll(dir, 0o755)
	if err != nil {
		panic(err)
	}
	copy("template/main.go", path.Join(dir, dir+".go"))
	copy("template/main_test.go", path.Join(dir, dir+"_test.go"))
	copy("template/testinput", path.Join(dir, "testinput"))
}

func copy(srcPath, dstPath string) error {
	data, err := os.ReadFile(srcPath)
	if err != nil {
		return err
	}
	err = os.WriteFile(dstPath, data, 0o644)
	if err != nil {
		return err
	}
	return nil
}
