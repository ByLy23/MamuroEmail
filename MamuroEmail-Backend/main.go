package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type ZincSearchRecord struct {
	Username  string `json:"username"`
	Directory string `json:"directory"`
	File 	string `json:"file_name"`
	Content string `json:"content"`
}

type IndexDirectory struct {
	Name string `json:"index"`
	Directory []*ZincSearchRecord `json:"records"`
}


type File struct {
	Name string `json:"Mail"`
	Content string `json:"content"`
}
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func indexingDirectory(path string) []*IndexDirectory{
	var index= []*IndexDirectory{}
	filepath.Walk(path, func(path1 string, info os.FileInfo, err error) error {
		check(err)
		if info.IsDir() {
			lenDir:= strings.Split(path1, "\\")
			if len(strings.Split(path1, "\\")) == 2 {
				index= append(index, &IndexDirectory{
				Name: lenDir[0], Directory: dir_to_json(path1,info.Name())})
		}
	}
		return nil
	})
	return index
}

func dir_to_json(path string,root string) []*ZincSearchRecord {
	var newRecord []*ZincSearchRecord
	filepath.Walk(path, func(path1 string, info os.FileInfo, err error) error {
		check(err)
		if !info.IsDir() {
			nueva:= strings.Split(path1, "\\")
			if len(nueva) >= 3 {
				newFile:= readFiles(path1)
				newRecord= append(newRecord, &ZincSearchRecord{
					Username: nueva[1],
					Directory: nueva[2],
					File: newFile[0].Name,
					Content: newFile[0].Content,
				})
			}
		}
		return nil
	})
	return newRecord
}

func readFiles(path string) []*File{
	var nameFile []*File
			nueva:= strings.Split(path, "\\")
			lenNueva:= len(nueva)
			dat, err := os.ReadFile(path)
			if err != nil {
				panic(err)
			}
			nameFile= append(nameFile, &File {
				Name: nueva[lenNueva-2]+"_"+nueva[lenNueva-1],
				Content: string(dat),
			})
		
	return nameFile
}

func createIndex (path string) {

	// var dir []*Directory
	// indexingDirectory("enron_mail_20110402/maildir", dir)
}

func main() {
	// user:="admin"
	// password:="Complexpass#123"
	// index:="maildir"
	// rootFiles:="enron_mail_20110402/maildir"
	// zinc_host := "http://localhost:4080"

	// router:= chi.NewRouter()
	// router.Get("/api/v1/index", func(w http.ResponseWriter, r *http.Request) {
		
	// })

	//se crea el index

	nameFile:= indexingDirectory("maildir")

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
	
}