package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

type asciiStruct struct {
	Txt         string
	Police      string
	SSt         string
	SSh         string
	STh         string
	SM1         string
	SM2         string
	ResultAscii string
	Title       string
}

const selected = "selected"

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		fmt.Fprintf(w, "ERROR 404")
		return
	}
	r.ParseForm()
	a := asciiStruct{Title: "ASCIII_ART_WEB", Police: r.FormValue("police"), Txt: r.FormValue("inputArea"), SSt: "", SSh: "", STh: "", SM1: "", SM2: ""}
	switch a.Police {
	case "standard":
		a.SSt = selected
	case "shadow":
		a.SSh = selected
	case "thinkertoy":
		a.STh = selected
	case "my-police1":
		a.SM1 = selected
	case "my-police2":
		a.SM2 = selected
	}
	generateAscii(a.Txt, a.Police)
	result, _ := os.ReadFile("result.txt")
	a.ResultAscii = string(result)
	t, _ := template.ParseFiles("./static/index.html")
	t.Execute(w, a)
	//t.ExecuteTemplate(w, "/ascii-art", a)
}

func main() {
	os.WriteFile("result.txt", []byte{}, 0644)

	// fs := http.FileServer(http.Dir("./static"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))
	// http.HandleFunc("/", handleIndex)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/", handleIndex)
	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func generateAscii(txt, police string) {
	result := ""
	standard, _ := os.ReadFile(police + ".txt")
	s := strings.Split(string(standard), "\n")
	if len(txt) > 0 && (txt[0] == 34 && txt[len(txt)-1] == 34) {
		txt = txt[1 : len(txt)-1]
	}
	txtTable := strings.Split(txt, "\r")
	if strings.Join(txtTable, "") == "" {
		for i := 1; i <= len(txtTable)-1; i++ {
			result += "\n"
		}
		os.WriteFile("result.txt", []byte{}, 0644)
		return
	}
	for _, w := range txtTable {
		if w == "" {
			result += "\n"
			continue
		}
		for i := 0; i != 8; i++ {
			for _, each := range w {
				if isWritable(each) {
					result += string(s[(i + (int(each)-32)*9 + 1)])
				}
			}
			result += "\n"
		}
	}
	os.WriteFile("result.txt", []byte(result), 0644)
}

func isWritable(r rune) bool {
	if r < 32 || r > 126 {
		return false
	}
	return true
}
