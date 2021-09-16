package produto

import "github.com/wolmirgarbin/go-web-alura/db"

const (
	query       = "select * from produtos"
	queryInsert = "insert into produtos(name, descricao, preco, quantidade) values (?, ?, ?, ?)"
)

func ListProducts() []Produto {
	db := db.ConnectaComBancoDeDados()

	selecionaTodosOsProdutos, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selecionaTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err := selecionaTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Preco = preco
		p.Descricao = descricao
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()

	return produtos
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

	insert, err := db.Prepare("delete from produtos where id = ?")
	if err != nil {
		panic(err.Error())
	}

	insert.Exec(id)
	defer db.Close()
}
