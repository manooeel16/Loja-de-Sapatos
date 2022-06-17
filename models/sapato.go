package models

import (
	"go_modules/db"
)

type Sapato struct {
	Id            int
	Cidade        string
	Municipio     string
	QtdEsperada   int
	QtdEstoque    int
}

func ListagemDosSapatos() []Sapato {
	db := db.ConectaAoBancoDeDados()

	listandoOsSapatos, err := db.Query("select * from sapatos")
	if err != nil {
		panic(err.Error())
	}

	c := Sapato{}
	sapatos := []Sapato{}

	for listandoOsSapatos.Next() {
		var id, qtdEsperada, qtdEstoque int
		var cidade, municipio string

		err = listandoOsSapatos.Scan(&id, &cidade, &municipio, &qtdEsperada, &qtdEstoque)
		if err != nil {
			panic(err.Error())
		}

		c.Id = id
		c.Cidade = cidade
		c.Municipio = municipio
		c.QtdEsperada = qtdEsperada
		c.QtdEstoque = qtdEstoque

		sapatos = append(sapatos, c)
	}
	defer db.Close()
	return sapatos
}
func CriaNovoSapato(cidade, municipio string, qtdEsperada, qtdEstoque int) {
	db := db.ConectaAoBancoDeDados()

	insereDadosNoBanco, err := db.Prepare("insert into sapatos(cidade, municipio, qtdEsperada, qtdEstoque) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(cidade, municipio, qtdEsperada, qtdEstoque)
	defer db.Close()

}

func DeletaSapato(id string) {
	db := db.ConectaAoBancoDeDados()

	deletarsapato, err := db.Prepare("delete from sapatos where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deletarsapato.Exec(id)
	defer db.Close()
}

func EditaSapato(id string) Sapato {
	db := db.ConectaAoBancoDeDados()

	editandoSapatoNoBanco, err := db.Query("select * from sapatos where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	sapatoParaAtualizar := Sapato{}

	for editandoSapatoNoBanco.Next() {
		var id, qtdEsperada, qtdEstoque int
		var cidade, municipio string

		err = editandoSapatoNoBanco.Scan(&id, &cidade, &municipio, &qtdEsperada, &qtdEstoque)
		if err != nil {
			panic(err.Error())
		}
		sapatoParaAtualizar.Id = id
		sapatoParaAtualizar.Cidade = cidade
		sapatoParaAtualizar.Municipio = municipio
		sapatoParaAtualizar.QtdEsperada = qtdEsperada
		sapatoParaAtualizar.QtdEstoque = qtdEstoque
	}
	defer db.Close()
	return sapatoParaAtualizar
}

func AtualizaSapato(id int, cidade, municipio string, qtdEsperada, qtdEstoque int) {
	db := db.ConectaAoBancoDeDados()

	AtualizaSapato, err := db.Prepare("update sapatos set cidade=$1, municipio=$2, qtdEsperada=$3, qtdEstoque=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	AtualizaSapato.Exec(cidade, municipio, qtdEsperada, qtdEstoque, id)
	defer db.Close()
}
