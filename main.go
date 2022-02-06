package main

import (
	"awesomeProject3/main/blockchain"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const (
	port        string = ":4000"
	templateDir string = "templates/"
)

var templates *template.Template

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	data := homeData{"빵코인 홈페이지", blockchain.GetBlockchain().AllBlocks()}
	templates.ExecuteTemplate(rw, "home", data)
}

func add(rw http.ResponseWriter, r *http.Request) {
	//
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(rw, "add", nil)
	case "POST":
		r.ParseForm()
		data := r.Form.Get("blockData")
		blockchain.GetBlockchain().AddBlock(data)
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)
	}
}

func main() {
	/*
		chain := blockchain.GetBlockchain()
		chain.AddBlock("Second Block")
		chain.AddBlock("Third Block")
		chain.AddBlock("Fourth Block")
		for _, block := range chain.AllBlocks() {
			fmt.Printf("data: %s\n", block.Data)
			fmt.Printf("hash: %s\n", block.Hash)
			fmt.Printf("prev: %s\n", block.PrevHash)
		}
	*/
	templates = template.Must(templates.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	http.HandleFunc("/", home)
	http.HandleFunc("/add", add)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatalln(http.ListenAndServe(port, nil))
}
