package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"go_modules/models"
)

var tempSapato = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsSapato := models.ListagemDosSapatos()
	tempSapato.ExecuteTemplate(w, "Index", todosOsSapato)
}

func New(w http.ResponseWriter, r *http.Request) {
	tempSapato.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		cidade := r.FormValue("cidade")
		municipio := r.FormValue("municipio")
		qtdEsperada := r.FormValue("qtde")
		qtdEstoque := r.FormValue("qtdest")

		qtdEsperadaConvertendoparaInt, err := strconv.Atoi(qtdEsperada)
		if err != nil {
			log.Println("Erro na quantidade do sapato Esperado:", err)
		}

		qtdEstoqueConvertendoparaInt, err := strconv.Atoi(qtdEstoque)
		if err != nil {
			log.Println("Erro na quantidade do sapato do estoque:", err)
		}

		models.CriaNovoSapato(cidade, municipio, qtdEsperadaConvertendoparaInt, qtdEstoqueConvertendoparaInt)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoSapato := r.URL.Query().Get("id")
	models.DeletaSapato(idDoSapato)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idDoSapato := r.URL.Query().Get("id")
	sapato := models.EditaSapato(idDoSapato)
	tempSapato.ExecuteTemplate(w, "Edit", sapato)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		cidade := r.FormValue("cidade")
		municipio := r.FormValue("municipio")
		qtdEsperada := r.FormValue("qtde")
		qtdEstoque := r.FormValue("qtdest")

		idConvertidaParaInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na convesão do ID para int:", err)
		}

		qtdEsperadaConvertendoParaInt, err := strconv.Atoi(qtdEsperada)
		if err != nil {
			log.Println("Erro na quantidade do sapato Esperado para int:", err)
		}

		qtdEstoqueConvertendoParaInt, err := strconv.Atoi(qtdEstoque)
		if err != nil {
			log.Println("Erro na quantidade do sapato do estoque para int:", err)
		}

		models.AtualizaSapato(idConvertidaParaInt, cidade, municipio, qtdEsperadaConvertendoParaInt, qtdEstoqueConvertendoParaInt)
	}
	http.Redirect(w, r, "/", 301)
}
