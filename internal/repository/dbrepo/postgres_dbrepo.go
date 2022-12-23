// この辺全部queryを投げて返ってきた結果を返す関数
package dbrepo

import (
	"backend/internal/models"
	"context"
	"database/sql"
	"fmt"
	"time"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

const dbTimeout = time.Second * 3 // 3 seconds

func (m *PostgresDBRepo) Connection() *sql.DB {
	return m.DB
}

func (m *PostgresDBRepo) AllMovies() ([]*models.Movie, error) {
	// defines when to timeout
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		select
			id, title, release_date, runtime,
			mpaa_rating, description, coalesce(image, ''),
			created_at, updated_at
		from
			movies
		order by
			title
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var movies []*models.Movie

	for rows.Next() {
		var movie models.Movie
		err := rows.Scan(
			&movie.Id,
			&movie.Title,
			&movie.ReleaseDate,
			&movie.RunTime,
			&movie.MPAARating,
			&movie.Description,
			&movie.Image,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		)

		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		movies = append(movies, &movie)
	}

	return movies, nil
}

func (m *PostgresDBRepo) OneMovie(id int) (*models.Movie, error) {
	fmt.Print("one_movie")
	fmt.Println("onemovie")
	// defines when to timeout
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `select id, title, release_date, runtime, mpaa_rating, description, coalesce(image, ''), created_at, updated_at from movies where id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var movie models.Movie
	err := row.Scan(&movie.Id, &movie.Title, &movie.ReleaseDate,
		&movie.RunTime, &movie.MPAARating, &movie.Description, &movie.Image, &movie.CreatedAt, &movie.UpdatedAt)

	fmt.Println("err")
	if err != nil {
		fmt.Println("err1")
		fmt.Println(err)
		return nil, err
	}

	// get genres. genre + movie_genres
	query = `select g.id, g.genre from movies_genres mg left join genres g on (mg.genre_id = g.id) where mg.movie_id = $1 order by g.genre`

	rows, err := m.DB.QueryContext(ctx, query, id)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("err2")
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var genres []*models.Genre

	for rows.Next() {
		var g models.Genre
		err := rows.Scan(
			// copy the first element of the row to the destination (g.ID)
			&g.ID,
			// copy the 2nd element of the row to the destination (g.Genre)
			&g.Genre,
		)
		if err != nil {
			fmt.Println("err3")
			fmt.Println(err)
			return nil, err
		}
		genres = append(genres, &g)
	}

	movie.Genres = genres

	return &movie, err
}

func (m *PostgresDBRepo) GetUserByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, email, first_name, last_name, password,
	created_at, updated_at from users where email = $1`
	var user models.User
	row := m.DB.QueryRowContext(ctx, query, email)
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil // &<var>でpointerアドレスを返す
}

func (m *PostgresDBRepo) GetUserByID(id int) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, email, first_name, last_name, password,
	created_at, updated_at from users where id = $1`
	var user models.User
	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil // &<var>でpointerアドレスを返す
}
