package rotas

import (
	"net/http"
)

// estrutura das rotas que vao ser utilizadas pela api
type Rota struct {
	Uri                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}
