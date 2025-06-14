package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
)

var (
	//go:embed templates
	aaa embed.FS

	//go:embed static
	bbb embed.FS
)

func main() {
	t := template.Must(template.ParseFS(aaa, "templates/index.gohtml"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_ = t.Execute(w, map[string]any{
			"BG": fmt.Sprintf("/static/bgs/%d.webp", rand.Intn(5)+1),
			"Deployments": []struct {
				Title       string
				Link        string
				Description string
			}{

				{Title: "DankMuzikk", Link: "https://dankmuzikk.com", Description: "Create, Share and Play Music Playlists."},
				{Title: "DankLyrics", Link: "https://danklyrics.com", Description: "Find lyrics for songs or something."},
				{Title: "DankScreen", Link: "https://screen.dankstuff.net", Description: "Display capture card's output into your browser."},
				{Title: "DankTodo", Link: "https://todo.dankstuff.net", Description: "The first htmx app with C (Ulfius)"},
				{Title: "DankNotes", Link: "https://notes.dankstuff.net", Description: "Bootleg version of Notion and Google Notes. (WIP)"},
				// {Title: "DankSim", Link: "https://danksim.com", Description: "Get an eSIM qucickly, ANYWHERE."},
				{Title: "DankDysk", Link: "https://dankdysk.com", Description: "Magical unlimited cloud drive. (WIP)"},
				{Title: "DankTorrent", Link: "https://dankdysk.com", Description: "Download anything as torrent. (WIP)"},
				{Title: "DankIP", Link: "https://dankip.com", Description: "Get a public IP for any server/computer. (WIP)"},
				{Title: "libdank", Link: "https://libdank.org", Description: "File compression extension but the result is a video. (WIP)"},
			},
		})
	})
	http.Handle("/static/", http.FileServer(http.FS(bbb)))
	log.Println("server running on port 8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
