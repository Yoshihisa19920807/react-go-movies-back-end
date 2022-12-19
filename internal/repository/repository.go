package repository

import "backend/internal/models"

type DatabaseRepo interface {
	AllMovies() ([]*models.Movie, error)
}