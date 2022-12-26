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
	fmt.Println("onemovie")
	// defines when to timeout
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `select id, title, release_date, runtime, mpaa_rating, description, coalesce(image, ''), created_at, updated_at from movies where id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var movie models.Movie
	err := row.Scan(&movie.Id, &movie.Title, &movie.ReleaseDate,
		&movie.RunTime, &movie.MPAARating, &movie.Description, &movie.Image, &movie.CreatedAt, &movie.UpdatedAt)

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
			// copy the first element of the row to the destination (g.ID pointer)
			&g.ID,
			// copy the 2nd element of the row to the destination (g.Genre pointer)
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

func (m *PostgresDBRepo) OneMovieForEdit(id int) (*models.Movie, []*models.Genre, error) {
	// defines when to timeout
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `select id, title, release_date, runtime, mpaa_rating, description, coalesce(image, ''), created_at, updated_at from movies where id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var movie models.Movie
	err := row.Scan(&movie.Id, &movie.Title, &movie.ReleaseDate,
		&movie.RunTime, &movie.MPAARating, &movie.Description, &movie.Image, &movie.CreatedAt, &movie.UpdatedAt)

	if err != nil {
		return nil, nil, err
	}

	// get genres. genre + movie_genres
	query = `select g.id, g.genre from movies_genres mg left join genres g on (mg.genre_id = g.id) where mg.movie_id = $1 order by g.genre`

	rows, err := m.DB.QueryContext(ctx, query, id)
	if err != nil && err != sql.ErrNoRows {
		return nil, nil, err
	}
	defer rows.Close()

	var genres []*models.Genre
	var genresArray []int

	for rows.Next() {
		var g models.Genre
		err := rows.Scan(
			// copy the first element of the row to the destination (g.ID)
			&g.ID,
			// copy the 2nd element of the row to the destination (g.Genre)
			&g.Genre,
		)
		if err != nil {
			return nil, nil, err
		}
		genres = append(genres, &g)
		genresArray = append(genresArray, g.ID)
	}

	movie.Genres = genres
	movie.GenresArray = genresArray

	var allGenres []*models.Genre

	query = "select id, genre from genres order by genre"
	gRows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, nil, err
	}
	defer gRows.Close()

	for gRows.Next() {
		var g models.Genre
		err := gRows.Scan(
			&g.ID,
			&g.Genre,
		)
		if err != nil {
			return nil, nil, err
		}
	}

	return &movie, allGenres, err
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

func (m *PostgresDBRepo) AllGenres() ([]*models.Genre, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, genre, created_at, updated_at from genres order by genre`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var genres []*models.Genre

	for rows.Next() {
		var g models.Genre
		err := rows.Scan(&g.ID, &g.Genre, &g.CreatedAt, &g.UpdatedAt)
		if err != nil {
			return nil, err
		}
		// & is to pointer, * is to variable
		genres = append(genres, &g)
	}
	return genres, nil
}

func (m *PostgresDBRepo) InsertMovie(movie models.Movie) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	stmt :=
		`insert into movies (title, description, release_date, runtime, mpaa_rating, created_at, updated_at, image) values ($1, $2, $3, $4, $5, $6, $7, $8) returning id`
	var newID int

	err := m.DB.QueryRowContext(ctx, stmt,
		movie.Title,
		movie.Description,
		movie.ReleaseDate,
		movie.RunTime,
		movie.MPAARating,
		movie.CreatedAt,
		movie.UpdatedAt,
		movie.Image,
	).Scan(&newID)
	fmt.Println("newID")
	fmt.Println(newID)

	if err != nil {
		return 0, err
	}

	return newID, nil

}

func (m *PostgresDBRepo) UpdateMovie(movie models.Movie) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `update movies set title`
}

func (m *PostgresDBRepo) UpdateMovieGenres(id int, genreIDs []int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from movies_genres where movie_id = $1`
	_, err := m.DB.ExecContext(ctx, stmt, id)

	if err != nil {
		return err
	}

	for _, n := range genreIDs {
		stmt := `insert into movies_genres (movie_id, genre_id) values ($1, $2)`
		_, err := m.DB.ExecContext(ctx, stmt, id, n)
		if err != nil {
			return err
		}
	}
	return nil
}
