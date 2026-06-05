package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	srcDoc     = "../book/"
	srcPreFace = "../book/_index.md"

	dstDoc   = "content/" // copy all src doc to dst doc
	dstAsset = "static/img/"
)

func walkMd(path string, info os.FileInfo, err error) error {
	// rules:
	//	- skip dirs
	if info.IsDir() {
		fmt.Printf("walkmd: skip dir %v\n", path)
		return nil
	}
	fmt.Printf("walkmd: handling %v\n", path)

	// read file
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("walkmd: cannot read %v\n", path)
	}

	// proc markdown file
	// rules:
	//   - replace ./assets to ../assets
	if strings.HasSuffix(path, "md") {
		if path == srcPreFace {
			// for preface log png
			data = bytes.Replace(data, []byte("./assets"), []byte("./img"), -1)
		} else {
			// for normal png
			data = bytes.Replace(data, []byte("./assets"), []byte("../../../../img"), -1)
		}

		dst := dstDoc + strings.TrimPrefix(path, srcDoc)
		fmt.Printf("walkmd: markdown writing %v\n", dst)
		// create directory first
		if _, e := os.Stat(filepath.Dir(dst)); os.IsNotExist(e) {
			e = os.MkdirAll(filepath.Dir(dst), os.ModePerm)
			if e != nil {
				return fmt.Errorf("walkmd: create folder fail err %v\n", e)
			}
		}
		err = ioutil.WriteFile(dst, data, os.ModePerm)
		if err != nil {
			return fmt.Errorf("walkmd: cannot write %v\n", dst)
		}
		return nil
	}

	// proc image file
	// rules:
	//   - copy .png, .svg
	if strings.Contains(path, ".png") || strings.Contains(path, ".svg") {
		parts := strings.Split(strings.TrimPrefix(path, srcDoc), "/")
		n := len(parts)
		if n == 0 {
			return fmt.Errorf("fail split %v\n", path)
		}
		dst := dstAsset + parts[n-1]
		fmt.Printf("walkmd: assest writing %v\n", dst)
		err = ioutil.WriteFile(dst, data, os.ModePerm)
		if err != nil {
			return fmt.Errorf("walkmd: cannot write %v\n", dst)
		}
	}
	return nil
}

func main() {
	dirs := [...]string{dstDoc, dstAsset}
	// 1. create all directory
	for _, d := range dirs {
		err := os.MkdirAll(d, os.ModePerm)
		if err != nil {
			panic(fmt.Errorf("make: failed to create folders: %v", err))
		}
	}

	// 2. walk all files
	err := filepath.Walk(srcDoc, walkMd)
	if err != nil {
		fmt.Printf("main: wakmd err %v\n", err)
	}
}
