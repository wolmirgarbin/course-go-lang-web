package produto

import (
	"log"
	"net/http"
	"strconv"
)

func ProductReceiver(r *http.Request) Produto {
	nome := r.FormValue("nome")
	descricao := r.FormValue("descricao")
	preco := r.FormValue("preco")
	quantidade := r.FormValue("quantidade")

	precoConvertido, err := strconv.ParseFloat(preco, 64)
	if err != nil {
		log.Println("Erro ao converter preço:", err)
	}

	quantidadeConvertida, err := strconv.Atoi(quantidade)
	if err != nil {
		log.Println("Erro ao converter quantidade:", err)
	}

	return Produto{
		Nome:       nome,
		Descricao:  descricao,
		Preco:      precoConvertido,
		Quantidade: quantidadeConvertida,
	}
}
