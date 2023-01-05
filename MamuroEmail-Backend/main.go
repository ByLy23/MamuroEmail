package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/cors"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

//Structs
// type Query struct {
// 	Term string `json:"term"`
// 	Start_time string `json:"start_time"`
// 	End_time string `json:"end_time"`
// }

// type Search struct {
// 	SearchType string `json:"search_type"`
// 	Query Query `json:"query"`
// 	From int `json:"from"`
// 	Max_results int `json:"max_results"`
// 	Source []string `json:"_source"`
// }
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

// functions
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func indexingDirectory(path string) {
	var index= []*IndexDirectory{}
	filepath.Walk(path, func(path1 string, info os.FileInfo, err error) error {
		check(err)
		if info.IsDir() {
			lenDir:= strings.Split(path1, "\\")
			if len(strings.Split(path1, "\\")) == 2 {
				index= append(index, &IndexDirectory{
				Name: lenDir[0], Directory: dir_to_json(path1,info.Name())})
				createJson(index)
				index=nil
		}
	}
		return nil
	})
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

func readingBody(r io.ReadCloser) string {
	body, err := io.ReadAll(r)
	check(err)
	return string(body)
}

func createJson(index []*IndexDirectory) {
	// var dir []*Directory
	// name:= index[0].Directory[0].Username+".json"
	newJson, err:= json.Marshal(index[0])
	check(err)
	resp, err:= postAPI("/api/_bulkv2",string(newJson))
	defer resp.Body.Close()
	 body, err := io.ReadAll(resp.Body)
	 check(err)
	 fmt.Println(string(body))
	if resp.StatusCode == 200 {
		log.Println("Indexing")
	} else if resp.StatusCode != 200 {
		body, err := io.ReadAll(resp.Body)
		check(err)
		log.Println(string(body))
	}
	// indexingDirectory("enron_mail_20110402/maildir", dir)
	// _,err:= os.Stat(name)
	// if os.IsNotExist(err) {
	// 	file, err:= os.Create(name)
	// 	check(err)
	// 	defer file.Close()
	// }
	// file, err:= os.OpenFile(name, os.O_APPEND|os.O_WRONLY, 0644)
	// check(err)
	// defer file.Close()
	// nuevo:=&index
	// b, err := json.Marshal(nuevo)
	// check(err)
	// _, err = file.WriteString(string(b))
	// check(err)
	// err=file.Sync()
	// check(err)
}

//requests
//create index function with chi router
func createIndex(w http.ResponseWriter, r *http.Request) {
	configurarCorsOptions(&w,r)
	index:="maildir"
	configIndex:= `{
		"name":"`+index+`",
		"storage_type":"disk",
		"shard_num":1,
		"mappings":{
			"properties":{
				"username":{
					"type":"text",
					"index":true,
					"store":true,
					"highlightable":true
				},
				"sub_folder":{
					"type":"text",
					"index":true,
					"store":true,
					"highlightable":true
				},
				"file_name":{
					"type":"text",
					"index":true,
					"store":true,
					"highlightable":true
				},
				"content":{
					"type":"text",
					"index":true,
					"store":true,
					"highlightable":true
				}
			}
		}
	}`
	rootFiles:="maildir"
	resp, err := postAPI("/api/index",configIndex)
	check(err)
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		w.Write([]byte("Index created"))
		w.Write([]byte("Indexing"))
		indexingDirectory(rootFiles)
	} else if resp.StatusCode != 200 {
		w.Write([]byte("Index already exists"))
	}
	// body, err := io.ReadAll(resp.Body)
	// check(err)
	// fmt.Println(string(body))
	// var index= []*IndexDirectory{}
	// var dir []*Directory
	// indexingDirectory("enron_mail_20110402/maildir", dir)
	// index= append(index, &IndexDirectory{
	// 	Name: "maildir", Directory: dir})
	// createJson(index)
	// index=nil
	// fmt.Println("Indexing")
	// w.Write([]byte("Indexing"))
}

func postAPI(endpoint string, body string) (*http.Response, error) {
	user:="admin"
	password:="Complexpass#123"
	zinc_host := "http://localhost:4080"
	req, err := http.NewRequest("POST", zinc_host+endpoint, strings.NewReader(body))
	check(err)
	req.SetBasicAuth(user,password)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	check(err)
	return resp, err
}
func configurarCorsOptions(w *http.ResponseWriter, request *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func searchMaildir(w http.ResponseWriter, r *http.Request) {
	// configurarCorsOptions(&w,r)
	req:=readingBody(r.Body)
	fmt.Println(r.Body)
	resp, err:= postAPI("/api/maildir/_search",string(req))
	check(err)

	defer resp.Body.Close()
	w.WriteHeader(resp.StatusCode)
	if resp.StatusCode == 200 {
		//return search results
		data:= readingBody(resp.Body)
		// dataJson,_:= json.Marshal(data)
		w.Header().Set("Content-Type", "application/json")
		
		w.Write([]byte(data))
		return
	} else if resp.StatusCode != 200 {
		type  Message struct {
			Message string `json:"message"`
		}
		//return error
		msjInvalido:= Message{Message: "Invalid query"}
		msjInvalidoJson,_:= json.Marshal(msjInvalido)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(msjInvalidoJson))
		return
	}
	// var q Query
	// var search Search
	// json.NewDecoder(r.Body).Decode(&q)
	// resp, err:= postAPI("/api/maildir/_search",r.Body)
}

//main
func main() {
	PORT:= "8080"

	log.Printf("Serving on port %s", PORT)
	r:= chi.NewRouter()
	//change to no require auth for options
	
	
	// accept cors access control origins
	r.Use(middleware.AllowContentType("application/json"))
	//header for cors
	r.Use(middleware.SetHeader("Access-Control-Allow-Origin", "*"))
	r.Use(middleware.SetHeader("Access-Control-Allow-Methods", "OPTIONS,POST,GET"))
	r.Use(middleware.SetHeader("Access-Control-Allow-Headers", "*"))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:     []string{"http://localhost:5500"},
		AllowedMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowedHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposedHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 300,
	}))
	r.Get("/api/index", createIndex)
	r.Post("/hola", func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("hola")
		w.Write([]byte("Hola"))
	})
	r.Options("/hola", func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("hola")
		w.Write([]byte("Hola"))
	})

	r.Options("/api/search", searchMaildir)
	r.Post("/api/search", searchMaildir)
	log.Fatal(http.ListenAndServe(":" + PORT, r))
}