package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	//err := os.Mkdir("subdir", 0755)
	//checkError(err)

	//checkError(os.Mkdir("subdir1", 0755))
	//os.WriteFile("subdir1/file", []byte(""), 0644)

	//checkError(os.MkdirAll("subdir/paret/child", 0755))
	//checkError(os.MkdirAll("subdir/paret/child1", 0755))
	//checkError(os.MkdirAll("subdir/paret/child2", 0755))
	//checkError(os.MkdirAll("subdir/paret/child3", 0755))
	//
	//os.WriteFile("subdir/paret/file", []byte(""), 0755)
	//os.WriteFile("subdir/paret/child/file", []byte(""), 0755)

	result, err := os.ReadDir("subdir/parent")
	checkError(err)

	for _, entry := range result {
		fmt.Println(entry.Name(), entry.IsDir(), entry.Type())
	}

	pathfile := "subdir"
	fmt.Println("Walking Directory")
	err = filepath.WalkDir(pathfile, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}

		fmt.Println(path)
		return nil
	})

	checkError(err)

	checkError(os.RemoveAll("subdir"))
	checkError(os.Remove("subdir1"))
}
