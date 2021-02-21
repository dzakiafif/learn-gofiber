package entity

type CreateBookRequest struct {
	ID int `json:"id"`
	BookName  string `json:"book_name" validate:"required"`
	BookAuthor string `json:"book_author" validate:"required"`
	BookYear string `json:"book_year"`
}

type UpdateBookRequest struct {
	BookName  string `json:"book_name"`
	BookAuthor string `json:"book_author"`
	BookYear string `json:"book_year"`
}