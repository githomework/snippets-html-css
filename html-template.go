import (
	"html/template"
)

func pageString(printer string) string {
	p := &struct {
		CreateUrl string
		Printer   string
	}{
		SubmitUrl: "/something/submit",
		Printer:   printer,
	}

	t := template.New("main")                             // Create a template.
	t, err := t.ParseFiles(global.Folder + "\\html.html") // Parse template file.
	if err != nil {
		log.Println(err.Error())
	}

	var b bytes.Buffer
	t.Execute(&b, p)

	return b.String()

}

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, pageString("printer name"))
}


/*

{{define "main"}}
<!doctype html>
<html>
<head>
<title></title>
</head>
<body style="background-color: #F1F1F1">






</body>
</html>
{{end}}

*/
