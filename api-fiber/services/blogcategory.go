package services

import (
	"api-fiber/config"
	"api-fiber/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BlogCategoryService struct {
	Conn *pgxpool.Pool
}

func (s *BlogCategoryService) Open() error {
	conn, err := pgxpool.New(config.CTX, config.URL_DATABASE)
	if err != nil {
		return err
	}

	// Assign connection
	s.Conn = conn
	return nil
}

func (s *BlogCategoryService) Close() {
	s.Conn.Close()
}

func (s *BlogCategoryService) Create(input *models.BlogCategory) error {
	// Execute SQL
	sql := "SELECT create_blogcategory(@name);"
	args := pgx.NamedArgs{"name": input.Name}
	_, err := s.Conn.Exec(config.CTX, sql, args)
	if err != nil {
		return err
	}

	// If success return nil
	return nil
}

func (s *BlogCategoryService) Update(input *models.BlogCategory) error {
	// Execute SQL
	sql := "SELECT update_blogcategory(@id, @name);"
	args := pgx.NamedArgs{"id": input.Id, "name": input.Name}
	_, err := s.Conn.Exec(config.CTX, sql, args)
	if err != nil {
		return err
	}

	// If success return nil
	return nil
}

func (s *BlogCategoryService) Remove(id *int) error {
	// Execute SQL
	sql := "SELECT remove_blogcategory(@id);"
	args := pgx.NamedArgs{"id": *id}
	_, err := s.Conn.Exec(config.CTX, sql, args)
	if err != nil {
		return err
	}

	// If success return nil
	return nil
}

func (s *BlogCategoryService) Count() (int, error) {
	// Execute SQL
	sql := "SELECT * FROM count_blogcategory();"
	value := 0
	err := s.Conn.QueryRow(config.CTX, sql).Scan(&value)
	if err != nil {
		return 0, err
	}

	// If success return nil
	return value, nil
}

func (s *BlogCategoryService) GetFirst() (models.BlogCategory, error) {
	// Execute SQL
	item := models.BlogCategory{}
	sql := "SELECT id, name FROM blogcategory ORDER BY id ASC LIMIT 1;"
	err := s.Conn.QueryRow(config.CTX, sql).Scan(&item.Id, &item.Name)
	if err != nil {
		return item, err
	}

	// If success return item
	return item, nil
}

func (s *BlogCategoryService) GetLast() (models.BlogCategory, error) {
	// Execute SQL
	item := models.BlogCategory{}
	sql := "SELECT id, name FROM blogcategory ORDER BY id DESC LIMIT 1;"
	err := s.Conn.QueryRow(config.CTX, sql).Scan(&item.Id, &item.Name)
	if err != nil {
		return item, err
	}

	// If success return item
	return item, nil
}

func (s *BlogCategoryService) GetAll(limit *int, page *int) ([]models.BlogCategory, error) {
	// Execute SQL
	sql := "SELECT * FROM get_all_blogcategory(@limit, @page);"
	args := pgx.NamedArgs{
		"limit": *limit,
		"page":  *page,
	}
	rows, err := s.Conn.Query(config.CTX, sql, args)
	if err != nil {
		return nil, err
	}

	// Convert SQL rows to data
	data := []models.BlogCategory{}
	for rows.Next() {
		model := models.BlogCategory{}

		if err := rows.Scan(&model.Id, &model.Name); err != nil {
			return nil, err
		}

		data = append(data, model)
	}

	// If success return nil
	return data, nil
}
