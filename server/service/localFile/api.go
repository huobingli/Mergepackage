package localFile

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func GetLocalPath(path string) {
	files, _:= ioutil.ReadDir(path)
	for _, f := range files{
		fmt.Println(f.Name())
	}
}

func GetLocalPath2(path string) {
	files, _ := filepath.Glob(path)
	fmt.Println(files)
}