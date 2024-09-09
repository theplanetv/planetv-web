package services

import (
	"api-fiber/config"
	"api-fiber/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BlogTagService struct {
	Conn *pgxpool.Pool
}

func (s *BlogTagService) Open() error {
	conn, err := pgxpool.New(config.CTX, config.URL_DATABASE)
	if err != nil {
		return err
	}

	// Assign connection
	s.Conn = conn
	return nil
}

func (s *BlogTagService) Close() {
	s.Conn.Close()
}

func (s *BlogTagService) Create(input *models.BlogTag) error {
	// Execute SQL
	sql := "SELECT create_blogtag(@blogcategory_id, @name);"
	args := pgx.NamedArgs{
		"blogcategory_id": input.BlogcategoryId,
		"name":            input.Name,
	}
	_, err := s.Conn.Exec(config.CTX, sql, args)
	if err != nil {
		return err
	}

	// If success return nil
	return nil
}

func (s *BlogTagService) Update(input *models.BlogTag) error {
	// Execute SQL
	sql := "SELECT update_blogtag("
	args := pgx.NamedArgs{}

	// Add attribute and change value
	if input.Id > 0 {
		sql += "@id,"
		args["id"] = input.Id
	} else {
		sql += "NULL,"
	}

	if input.BlogcategoryId > 0 {
		sql += "@blogcategory_id,"
		args["blogcategory_id"] = input.BlogcategoryId
	} else {
		sql += "NULL,"
	}

	if input.Name != "" {
		sql += "@name,"
		args["name"] = input.Name
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

func (s *BlogTagService) Remove(id *int) error {
	// Execute SQL
	sql := "SELECT remove_blogtag(@id);"
	args := pgx.NamedArgs{"id": *id}
	_, err := s.Conn.Exec(config.CTX, sql, args)
	if err != nil {
		return err
	}

	// If success return nil
	return nil
}

func (s *BlogTagService) Count() (int, error) {
	// Execute SQL
	sql := "SELECT * FROM count_blogtag();"
	value := 0
	err := s.Conn.QueryRow(config.CTX, sql).Scan(&value)
	if err != nil {
		return 0, err
	}

	// If success return nil
	return value, nil
}

func (s *BlogTagService) GetFirst() (models.BlogTag, error) {
	// Execute SQL
	item := models.BlogTag{}
	sql := "SELECT bt.id, bt.blogcategory_id, bt.name, bc.name FROM blogtag bt INNER JOIN blogcategory bc ON bc.id = bt.blogcategory_id ORDER BY id ASC LIMIT 1;"
	err := s.Conn.QueryRow(config.CTX, sql).Scan(
		&item.Id,
		&item.BlogcategoryId,
		&item.Name,
		&item.BlogcategoryName,
	)
	if err != nil {
		return item, err
	}

	// If success return item
	return item, nil
}

func (s *BlogTagService) GetLast() (models.BlogTag, error) {
	// Execute SQL
	item := models.BlogTag{}
	sql := "SELECT bt.id, bt.blogcategory_id, bt.name, bc.name FROM blogtag bt INNER JOIN blogcategory bc ON bc.id = bt.blogcategory_id ORDER BY id DESC LIMIT 1;"
	err := s.Conn.QueryRow(config.CTX, sql).Scan(
		&item.Id,
		&item.BlogcategoryId,
		&item.Name,
		&item.BlogcategoryName,
	)
	if err != nil {
		return item, err
	}

	// If success return item
	return item, nil
}

func (s *BlogTagService) GetAll(limit *int, page *int) ([]models.BlogTag, error) {
	// Execute SQL
	sql := "SELECT * FROM get_all_blogtag(@limit, @page);"
	args := pgx.NamedArgs{
		"limit": *limit,
		"page":  *page,
	}
	rows, err := s.Conn.Query(config.CTX, sql, args)
	if err != nil {
		return nil, err
	}

	// Convert SQL rows to data
	data := []models.BlogTag{}
	for rows.Next() {
		model := models.BlogTag{}

		if err := rows.Scan(
			&model.Id,
			&model.BlogcategoryId,
			&model.Name,
			&model.BlogcategoryName,
		); err != nil {
			return nil, err
		}

		data = append(data, model)
	}

	// If success return nil
	return data, nil
}
