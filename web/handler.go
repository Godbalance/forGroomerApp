package web

import (
	"forGroomerApp/database"
	"html/template"
	"net/http"
)

func HandleRequest() {
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("./styles/"))))
	http.Handle("/res/", http.StripPrefix("/res/", http.FileServer(http.Dir("./res"))))
	http.HandleFunc("/admin/", adminPage)
	//http.HandleFunc("/set/", setValue)
	http.HandleFunc("/webapp/", webAppPage)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
func pro() {

}

//func setValue(w http.ResponseWriter, r *http.Request) {
//	path := r.FormValue("path")
//}

func webAppPage(w http.ResponseWriter, r *http.Request) {
	d := database.GetRealtime()
	t, _ := template.ParseFiles("layouts/webapp.html", "layouts/card.html")
	err := t.ExecuteTemplate(w, "webapp", d)
	if err != nil {
		return
	}
}

func adminPage(w http.ResponseWriter, r *http.Request) {
	d := database.GetRealtime()
	t, _ := template.ParseFiles("layouts/admin.html")
	err := t.Execute(w, d)
	if err != nil {
		return
	}
}
