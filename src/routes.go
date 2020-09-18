package main

//create routes
func (s *Server) routes() {
	s.router.HandleFunc("/assetregister", s.handleexportasset()).Methods("GET")
}
