package main


import (
    "html/template"
	"net/http"
	"os"
	"strconv"
)



type Receipe struct {
    Title string
    Done  bool
}



type ReceipePageData struct {
    PageTitle string
	Receipes     []Receipe
	Color	bool
}




func main() {
	colorSet, err := strconv.ParseBool(os.Getenv("COLOR"))
	if err != nil {
		colorSet = false
	}
    tmpl := template.Must(template.ParseFiles("template.html"))
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data := ReceipePageData{
            PageTitle: "Confee Recipe",
            Receipes: []Receipe{
                {Title: "Beans", Done: false},
                {Title: "Milk", Done: true},
                {Title: "Cinnamon", Done: true},
			},
			Color: colorSet,
        }
        tmpl.Execute(w, data)
    })
    http.ListenAndServe(":8080", nil)
}