package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type File struct {
	Name string `json:"name"`
	Content string `json:"content"`
}

type IndexDirectory struct {
	Name string `json:"name"`
	Directory []*Directory `json:"directories"`
}

type Directory struct {
	Name string `json:"name"`
	File []*File `json:"files"`
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

//TODO: hacer que este metodo no retorne nada
func indexingDirectory(path string,dir []*Directory) []*IndexDirectory {
	var index= []*IndexDirectory{}
	filepath.Walk(path, func(path1 string, info os.FileInfo, err error) error {
		check(err)
		if info.IsDir() {
			if len(strings.Split(path1, "\\")) == 2 {
				index= append(index, &IndexDirectory{
				Name: info.Name(), Directory: dir_to_json(path1, dir,info.Name())})
		}
	}

		return nil
	})
	return index
}

func dir_to_json(path string, dir []*Directory,root string) []*Directory {
	// prueba:=""
	var dir1 =dir
	filepath.Walk(path, func(path1 string, info os.FileInfo, err error) error {
		check(err)
		if info.IsDir() {

			nueva:= strings.Split(path1, "\\")
			if len(nueva) >= 3 {
			dir1= append(dir1, &Directory{
							Name: info.Name(), File: readFiles(path1)})
						}
			if len(nueva) == 3 {
			fmt.Println(root)
			}
						// dir=nil
			// if len(nueva) >= 3 {
			// 	if prueba != nueva[2] {
			// 		prueba = nueva[2]
			// 		dir= append(dir, &Directory{
			// 			Name: info.Name(), File: readFiles(path1)})
			// 		fmt.Println(dir[0].File[0])
			// 		dir=nil
			// 		//crear cada instancia de la carpeta y sus contenidos
			// 		//enviar a la base de datos el nombre de la carpeta con sus contenidos
			// 	}
			// }
			
			// fmt.Println(nameFile)
		}
		return nil
	})
	return dir1
}

func readFiles(path string) []*File{
	var nameFile []*File
	filepath.Walk(path, func(path1 string, info os.FileInfo, err error) error {
		check(err)
		if !info.IsDir() {
			dat, err := os.ReadFile(path1)
			if err != nil {
				panic(err)
			}
			// dat := bufio.NewReader(f)
			// b4, err1 := dat.Peek(5)
			// if err1 != nil {
			// 	panic(err1)
			// }

			nameFile= append(nameFile, &File{
				Name: info.Name(), Content: string(dat)})
				
			// fmt.Println(path1)
		} 
		return nil
	})
	return nameFile
}

func main() {
	var dir []*Directory
	//se crea el index

	nameFile:= indexingDirectory("maildir", dir)

	fmt.Println(nameFile)
	_,err:= os.Stat("maildir.json")
	if os.IsNotExist(err) {
		file, err:= os.Create("maildir.json")
		check(err)
		defer file.Close()
	}
	file, err:= os.OpenFile("maildir.json", os.O_APPEND|os.O_WRONLY, 0644)
	check(err)
	defer file.Close()
	nuevo:=&nameFile
	b, err := json.Marshal(nuevo)
	check(err)
	_, err = file.WriteString(string(b))
	check(err)
	err=file.Sync()
	check(err)
	// ENDPOINT := "http://localhost:4080"
	
}