package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

type Story map[string]Arc

type Arc struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

var defaultHandlerTmpl = `
<!DOCTYPE html>
<head>
    <meta charset="utf-8">
    <title>Choose Your Own Adventure</title>
</head>
<body>
    <h1>{{.Title}}</h1>
        {{range .Paragraphs}}
            <p>{{.}}</p>
        {{end}}
    <ul>
        {{range .Options}}
            <li> <a href="/{{.Arc}}">{{.Text}}</a></li>
        {{end}}
    </ul>
</body>`

type handler struct {
	s Story
}

func (h handler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.New("").Parse(defaultHandlerTmpl))

	urlPath := req.URL.Path[1:]
	arc, ok := h.s[urlPath]
	if !ok {
		arc = h.s["intro"]
	}

	err := tpl.Execute(writer, arc)
	if err != nil {
		panic(fmt.Sprintf("Failed to execute the path %s\n", urlPath))
	}
}

func NewHandler(story Story) http.Handler {
	return handler{
		s: story,
	}
}

func main() {
	jsonFile := flag.String("jsonFile", "gopher.json", "the json file to get the story")
	flag.Parse()

	file, err := os.Open(*jsonFile)
	if err != nil {
		panic(fmt.Sprintf("Couldn't open the file %s\n", *jsonFile))
	}

	d := json.NewDecoder(file)
	var story Story
	if err := d.Decode(&story); err != nil {
		panic(err)
	}

	httpHandler := NewHandler(story)

	fmt.Println("Starting the server on :8080")
	log.Fatal(http.ListenAndServe(":8080", httpHandler))
}
