package services

import (
	"api-fiber/config"
	"api-fiber/models"

	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BlogFileService struct {
	Conn *pgxpool.Pool
}

func (s *BlogFileService) Open() error {
	conn, err := pgxpool.New(config.CTX, config.URL_DATABASE)
	if err != nil {
		return err
	}

	// Assign connection
	s.Conn = conn
	return nil
}

func (s *BlogFileService) Close() {
	s.Conn.Close()
}

func (s *BlogFileService) Create(input *models.BlogFile) error {
	// Execute SQL
	sql := "SELECT create_blogfile(@filename);"
	args := pgx.NamedArgs{
		"filename": input.Filename,
	}
	_, err := s.Conn.Exec(config.CTX, sql, args)
	if err != nil {
		return err
	}

	// If success return nil
	return nil
}

func (s *BlogFileService) Update(input *models.BlogFile) error {
	// Execute SQL
	sql := "SELECT update_blogfile("
	args := pgx.NamedArgs{}

	// Add attribute and change value
	if input.Id > 0 {
		sql += "@id,"
		args["id"] = input.Id
	} else {
		sql += "NULL,"
	}

	if input.Filename != "" {
		sql += "@filename,"
		args["filename"] = input.Filename
	} else {
		sql += "NULL,"
	}

	if input.CreatedDate != (time.Time{}) {
		sql += "@created_date,"
		args["created_date"] = input.CreatedDate
	} else {
		sql += "NULL,"
	}

	if input.UpdatedDate != (time.Time{}) {
		sql += "@updated_date,"
		args["updated_date"] = input.UpdatedDate
	} else {
		sql += "NULL,"
	}

	// Remove last ,
	sql = sql[:len(sql)-1]
	sql += ");"

	// Execute SQL
	_, err := s.Conn.Exec(config.CTX, sql, args)
	if err != nil {
		return err
	}

	// If success return nil
	return nil
}

func (s *BlogFileService) Remove(id *int) error {
	// Execute SQL
	sql := "SELECT remove_blogfile(@id);"
	args := pgx.NamedArgs{"id": *id}
	_, err := s.Conn.Exec(config.CTX, sql, args)
	if err != nil {
		return err
	}

	// If success return nil
	return nil
}

func (s *BlogFileService) Count() (int, error) {
	// Execute SQL
	sql := "SELECT * FROM count_blogfile();"
	value := 0
	err := s.Conn.QueryRow(config.CTX, sql).Scan(&value)
	if err != nil {
		return 0, err
	}

	// If success return nil
	return value, nil
}

func (s *BlogFileService) GetFirst() (models.BlogFile, error) {
	// Execute SQL
	item := models.BlogFile{}
	sql := "SELECT id, filename, created_date, updated_date FROM blogfile ORDER BY id ASC LIMIT 1;"
	err := s.Conn.QueryRow(config.CTX, sql).Scan(
		&item.Id,
		&item.Filename,
		&item.CreatedDate,
		&item.UpdatedDate,
	)
	if err != nil {
		return item, err
	}

	// If success return item
	return item, nil
}

func (s *BlogFileService) GetLast() (models.BlogFile, error) {
	// Execute SQL
	item := models.BlogFile{}
	sql := "SELECT id, filename, created_date, updated_date FROM blogfile ORDER BY id DESC LIMIT 1;"
	err := s.Conn.QueryRow(config.CTX, sql).Scan(
		&item.Id,
		&item.Filename,
		&item.CreatedDate,
		&item.UpdatedDate,
	)
	if err != nil {
		return item, err
	}

	// If success return item
	return item, nil
}

func (s *BlogFileService) GetAll(limit *int, page *int) ([]models.BlogFile, error) {
	// Execute SQL
	sql := "SELECT * FROM get_all_blogfile(@limit, @page);"
	args := pgx.NamedArgs{
		"limit": *limit,
		"page":  *page,
	}
	rows, err := s.Conn.Query(config.CTX, sql, args)
	if err != nil {
		return nil, err
	}

	// Convert SQL rows to data
	data := []models.BlogFile{}
	for rows.Next() {
		model := models.BlogFile{}

		if err := rows.Scan(
			&model.Id,
			&model.Filename,
			&model.CreatedDate,
			&model.UpdatedDate,
		); err != nil {
			return nil, err
		}

		data = append(data, model)
	}

	// If success return nil
	return data, nil
}
