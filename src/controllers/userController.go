package controllers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Cria usuário"))
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Busca um usuário"))
}
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Busca todos os usuário"))
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualiza usuário"))
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleta usuário"))
}
