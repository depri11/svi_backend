package delivery

func (d *articleDelivery) Routes() {
	d.r.HandleFunc("/article", d.CreateArticle).Methods("POST")
}
