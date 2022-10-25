package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	route.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	route.HandleFunc("/", handlerIndex).Methods("GET")
	route.HandleFunc("/project", handlerProject).Methods("GET")
	route.HandleFunc("/contact", handlerContact).Methods("GET")
	route.HandleFunc("/project-detail/{id}", handlerProjectDetail).Methods("GET")
	route.HandleFunc("/new-project", newProject).Methods("POST")

	fmt.Println("Server is running on localhost:5000")
	err := http.ListenAndServe("localhost:5000", route)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	view, err := template.ParseFiles("views/index.html")

	fmt.Println("Index Page")

	if err != nil {
		w.Write([]byte("message : " + err.Error()))
		return
	}

	view.Execute(w, nil)
}

func handlerProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	view, err := template.ParseFiles("views/project.html")

	fmt.Println("Project Page")

	if err != nil {
		w.Write([]byte("message : " + err.Error()))
		return
	}

	view.Execute(w, nil)
}

func handlerContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	view, err := template.ParseFiles("views/contact.html")

	fmt.Println("Contact Page")

	if err != nil {
		w.Write([]byte("message : " + err.Error()))
		return
	}

	view.Execute(w, nil)
}

func newProject(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ProjectName : " + r.PostForm.Get("projectName"))
	fmt.Println("Description : " + r.PostForm.Get("description"))

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func handlerProjectDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	view, err := template.ParseFiles("views/project-detail.html")

	if err != nil {
		w.Write([]byte("message : " + err.Error()))
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	data := map[string]interface{}{
		"Title":   "Pasar Coding di Indonesia Dinilai Masih Menjanjikan",
		"Content": "REPUBLIKA.CO.ID, JAKARTA -- Ketimpangan sumber daya manusia (SDM) disektor digital masih menjadi isu yang belum terpecahkan. Berdasarkan penelitian ManpowerGroup.REPUBLIKA.CO.ID, JAKARTA -- Ketimpangan sumber daya manusia (SDM) disektor digital masih menjadi isu yang belum terpecahkan. Berdasarkan penelitian ManpowerGroup.REPUBLIKA.CO.ID, JAKARTA -- Ketimpangan sumber daya manusia (SDM) disektor digital masih menjadi isu yang belum terpecahkan. Berdasarkan penelitian ManpowerGroup.",
		"Id":      id,
	}

	view.Execute(w, data)
}
