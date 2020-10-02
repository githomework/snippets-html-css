import (
		"html/template"
)


func mainPage(w http.ResponseWriter, r *http.Request) {

	p := struct{}{}
	t := template.New("main")                             // Create a template.
	t, err := t.ParseFiles(global.Folder + "\\main.html") // Parse template file.
	if err != nil {
		log.Println(err.Error())
	}

	t.Execute(w, p) // merge.
}


