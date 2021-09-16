package produto

import (
	"log"
	"net/http"
	"strconv"
)

func ProductReceiver(r *http.Request) Produto {
	id := getToInt(r.FormValue("id"))
	nome := r.FormValue("nome")
	descricao := r.FormValue("descricao")
	preco := r.FormValue("preco")
	quantidade := getToInt(r.FormValue("quantidade"))

	precoConvertido, err := strconv.ParseFloat(preco, 64)
	if err != nil {
		log.Println("Erro ao converter preço:", err)
	}

	product := Produto{
		Id:         id,
		Nome:       nome,
		Descricao:  descricao,
		Preco:      precoConvertido,
		Quantidade: quantidade,
	}

	return product
}

func getToInt(value string) int {
	idConvertido, err := strconv.Atoi(value)
	if err != nil {
		log.Println("Erro ao converter valor:", value, err)
		return -1
	}

	return idConvertido
}
