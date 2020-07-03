package main

func main() {

	server := NewServer(":3000")
	server.Handle("/", HandleRoot)
	server.Handle("/api", server.AddMiddleware(HandleHome, CheckAuth())) //CheckAuth retorna la funcion:  func(f http.HandlerFunc) http.HandlerFunc
	server.Listen()

}
