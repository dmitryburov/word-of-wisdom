package repository

import "github.com/dmitryburov/word-of-wisdom/internal/repository/file"

type Quotes interface {
	GetQuote() (string, error)
}

type Repositories struct {
	Quotes
}

func NewRepositories() Repositories {
	return Repositories{
		Quotes: file.NewQuote(),
	}
}
