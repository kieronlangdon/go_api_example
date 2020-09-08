package main

type BookResponse struct {
	ID     string `json:"id"`
	Isbn   string `json:"isbn"`
	Title  string `json:"title"`
	Author Author `json:"author"`
}

// Author Struct
type AuthorResponse struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
