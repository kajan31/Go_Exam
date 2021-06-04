package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	 
	http.HandleFunc("/", list)
	//http.HandleFunc("/done", done)
	http.HandleFunc("/add", add)

	http.ListenAndServe(":8800", nil)
}

type Task struct {
	Description string
	Done bool
}
var tasks []Task



func list(rw http.ResponseWriter, _ *http.Request){
	for i,val := range tasks{
		if val.Done == false {
			chaine := fmt.Sprintf("ID: %d, task:%s",i , val.Description)
			bt := []byte (chaine)
			rw.WriteHeader(http.StatusOK)
			rw.Write(bt)
		}
	}

}
func done(rw http.ResponseWriter, _ *http.Request){


}
func add(rw http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Errorreadingbody:%v", err)
		http.Error(rw,"can'treadbody", http.StatusBadRequest,)
		return
	}
	descrip := string(body)

	fmt.Println(descrip)
	tasks = append(tasks,Task{descrip, false})// pas obligatoire de preciser false (Valeur par defaut)


}