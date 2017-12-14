package main

import (
	"net/http"
	"log"
	"encoding/json"
	"strconv"
	"fmt"
)

type MyHandler struct{
	pattern map[string]func(w http.ResponseWriter, r *http.Request)
}

func (h *MyHandler) Add(path string, f func(w http.ResponseWriter, r *http.Request)) {
	if h.pattern == nil {
		h.pattern = make(map[string]func(w http.ResponseWriter, r *http.Request), 10)
	}
	h.pattern[path] = f

	fmt.Println("Add:", path)
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Println(r.URL.Path)
	if r.URL.Path == "/" {
		GetPeople(w, r)
	} else if h.pattern[r.URL.Path] == nil {
		http.NotFound(w, r)
	} else {
		f := h.pattern[r.URL.Path]
		f(w, r)
	}

	return
}


type Person struct {
	Id int `json:"id, omitempty"`
	Name string `json:"name, omitempty"`
	Tel string `json:"tel, omitempty"`
	Age int `json: "age,omitempty"`
}

var people []Person

func CreatePerson(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.URL.Path[15:])
	fmt.Println("CreatePerson:", r.Method)
	if err != nil {
		w.Write([]byte(`<h1>Create: URL-ID非数值，出错</h1>`))
	} else {
		var person Person
		json.NewDecoder(r.Body).Decode(&person)
		for _, item := range people {
			if id == item.Id {
				w.Write([]byte(`<h1>Id重复</h1>`))
				return
			}
		}
		person.Id = id
		people = append(people, person)
		json.NewEncoder(w).Encode(people)
	}

	return
}

func GetPeople(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(people)

	fmt.Println("GetPeople:", r.Method)
}

func GetPerson(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.URL.Path[8:])
	fmt.Println("GetPerson:", r.Method)
	if err != nil {
		w.Write([]byte(`<h1>GetOne: URL-ID非数值，出错</h1>`))
	} else if id >= len(people) || id < 0 {
		w.Write([]byte(`<h1>URL-ID过大，出错</h1>`))
	} else {
		for _, item := range people {
			if id == item.Id {
				json.NewEncoder(w).Encode(item)
				break
			}
		}
	}

	return
}

func DeletePerson(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.URL.Path[15:])
	fmt.Println("DeletePeople:", r.Method)
	if err != nil {
		w.Write([]byte(`<h1>Create: URL-ID非数值，出错</h1>`))
	} else {
		for index, item := range people {
			if id == item.Id {
				people = append(people[:index], people[index+1:]...)
			}
		}
		json.NewEncoder(w).Encode(people)
	}

	return
}


func main(){
	myHandler := &MyHandler{}

	people = append(people, Person{Id:1, Name:"eternal", Tel:"18772580284", Age:21})
	people = append(people, Person{Id:2, Name:"peach", Tel:"13207134496", Age:20})
	myHandler.Add("/people", GetPeople)
	myHandler.Add("/people/:id", GetPerson)
	myHandler.Add("/people/create/:id", CreatePerson)
	myHandler.Add("/people/delete/:id", DeletePerson)

	err := http.ListenAndServe(":9090", myHandler)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}