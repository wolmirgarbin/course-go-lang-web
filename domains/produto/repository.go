package produto

import (
	"database/sql"

	"github.com/wolmirgarbin/go-web-alura/db"
)

const (
	query        = "select * from produtos"
	queryGetById = "select * from produtos where id = ?"
	queryInsert  = "insert into produtos(name, descricao, preco, quantidade) values (?, ?, ?, ?)"
	queryDelete  = "delete from produtos where id = ?"
	queryUpdate  = "update produtos set name = ?, descricao = ?, quantidade = ?, preco = ? where id = ?"
)

func ListProducts() []Produto {
	db := db.ConnectaComBancoDeDados()

	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	produtos := []Produto{}

	for rows.Next() {
		produtos = append(produtos, getProduct(rows))
	}
	defer db.Close()

	return produtos
}

func GetById(id string) Produto {
	db := db.ConnectaComBancoDeDados()

	rows, err := db.Query(queryGetById, id)
	if err != nil {
		panic(err.Error())
	}

	produto := Produto{}
	if rows.Next() {
		produto = getProduct(rows)
	}
	defer db.Close()

	return produto
}

func getProduct(rows *sql.Rows) Produto {
	produto := Produto{}

	var id, quantidade int
	var nome, descricao string
	var preco float64

	err := rows.Scan(&id, &nome, &descricao, &preco, &quantidade)
	if err != nil {
		panic(err.Error())
	}

	produto.Id = id
	produto.Nome = nome
	produto.Preco = preco
	produto.Descricao = descricao
	produto.Quantidade = quantidade

	return produto
}

func Save(produto Produto) {
	db := db.ConnectaComBancoDeDados()

	insert, err := db.Prepare(queryInsert)
	if err != nil {
		panic(err.Error())
	}

	insert.Exec(produto.Nome, produto.Descricao, produto.Preco, produto.Quantidade)
	defer db.Close()
}

func Delete(id string) {
	db := db.ConnectaComBancoDeDados()

	insert, err := db.Prepare(queryDelete)
	if err != nil {
		panic(err.Error())
	}

	insert.Exec(id)
	defer db.Close()
}

func Edit(produto Produto) {
	db := db.ConnectaComBancoDeDados()

	update, err := db.Prepare(queryUpdate)
	if err != nil {
		panic(err.Error())
	}

	update.Exec(
		produto.Nome,
		produto.Descricao,
		produto.Quantidade,
		produto.Preco,
		produto.Id)

	defer db.Close()
}
