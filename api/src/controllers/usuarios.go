package controllers

import (
	"net/http"
)

//funçoes para usar as rotas crud

// criar novo  user create
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando user"))
}

// buscar todos users read all
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("buscando todos os users"))
}

// buscar um usuario read one user
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("buscando user"))
}

// atualizar user update user
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando user"))
}

// deletar  user delete user
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando user"))
}
