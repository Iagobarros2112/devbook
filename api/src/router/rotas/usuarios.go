package rotas

import "net/http"

//primeira rota

var rotasUsuarios = []Rota{

	//primeira rota cadastro usuario
	{
		Uri:    "/usuarios",
		Metodo: http.MethodPost,
		Funcao: func(w http.ResponseWriter, r *http.Request) {

		},
		RequerAutenticacao: false,
	},

	//segunda busca todos usuarios

	{
		Uri:    "/usuarios",
		Metodo: http.MethodGet,
		Funcao: func(w http.ResponseWriter, r *http.Request) {

		},
		RequerAutenticacao: false,
	},

	//terceira busca um unico usuario

	{
		Uri:    "/usuarios/{usuarioId}",
		Metodo: http.MethodGet,
		Funcao: func(w http.ResponseWriter, r *http.Request) {

		},
		RequerAutenticacao: false,
	},

	//quarta

	{
		Uri:    "/usuarios/{usuarioId}",
		Metodo: http.MethodPut,
		Funcao: func(w http.ResponseWriter, r *http.Request) {

		},
		RequerAutenticacao: false,
	},

	//quinta

	{
		Uri:    "/usuarios/{usuarioId}",
		Metodo: http.MethodDelete,
		Funcao: func(w http.ResponseWriter, r *http.Request) {

		},
		RequerAutenticacao: false,
	},
}
