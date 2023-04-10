package main

import (
	"archive/tar"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("./output.tar")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.close()

	tw := tar.NewWriter(f)
	defer tw.Close()

	fileinfo, err := os.Stat("./test.txt")
	if err != nil {
		fmt.Println(err)
	}
	hdr, err := tar.FileInfoHeader(fileinfo, "")
	if err != nil {
		fmt.Println(err)
	}
	err = tw.WriteHeader(hdr)
	if err != nil {
		fmt.Println(err)
	}


}
