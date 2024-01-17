package delivery

func (d *articleDelivery) Routes() {
	d.r.HandleFunc("/article", d.CreateArticle).Methods("POST")
	d.r.HandleFunc("/article/{limit}/{offset}", d.GetArticles).Methods("GET")
	d.r.HandleFunc("/article/{id}", d.GetArticle).Methods("GET")
	d.r.HandleFunc("/article/{id}", d.UpdateArticle).Methods("PUT")
	d.r.HandleFunc("/article/{id}", d.DeleteArticle).Methods("DELETE")
}
