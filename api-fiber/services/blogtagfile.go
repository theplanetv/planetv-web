package services

import (
	"api-fiber/config"
	"api-fiber/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BlogTagFileService struct {
	Conn *pgxpool.Pool
}

func (s *BlogTagFileService) Open() error {
	conn, err := pgxpool.New(config.CTX, config.URL_DATABASE)
	if err != nil {
		return err
	}

	// Assign connection
	s.Conn = conn
	return nil
}

func (s *BlogTagFileService) Close() {
	s.Conn.Close()
}

func (s *BlogTagFileService) Create(input *models.BlogTagFile) error {
	// Execute SQL
	sql := "SELECT create_blogtagfile(@blogtag_id, @blogfile_id);"
	args := pgx.NamedArgs{
		"blogtag_id":  input.BlogtagId,
		"blogfile_id": input.BlogfileId,
	}
	_, err := s.Conn.Exec(config.CTX, sql, args)
	if err != nil {
		return err
	}

	// If success return nil
	return nil
}

func (s *BlogTagFileService) Update(input *models.BlogTagFile) error {
	// Execute SQL
	sql := "SELECT update_blogtagfile("
	args := pgx.NamedArgs{}

	// Add attribute and change value
	if input.Id > 0 {
		sql += "@id,"
		args["id"] = input.Id
	} else {
		sql += "NULL,"
	}

	if input.BlogtagId > 0 {
		sql += "@blogtag_id,"
		args["blogtag_id"] = input.BlogtagId
	} else {
		sql += "NULL,"
	}

	if input.BlogfileId > 0 {
		sql += "@blogfile_id,"
		args["blogfile_id"] = input.BlogfileId
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

func (s *BlogTagFileService) Remove(id *int) error {
	// Execute SQL
	sql := "SELECT remove_blogtagfile(@id);"
	args := pgx.NamedArgs{"id": *id}
	_, err := s.Conn.Exec(config.CTX, sql, args)
	if err != nil {
		return err
	}

	// If success return nil
	return nil
}

func (s *BlogTagFileService) Count() (int, error) {
	// Execute SQL
	sql := "SELECT * FROM count_blogtagfile();"
	value := 0
	err := s.Conn.QueryRow(config.CTX, sql).Scan(&value)
	if err != nil {
		return 0, err
	}

	// If success return nil
	return value, nil
}

func (s *BlogTagFileService) GetFirst() (models.BlogTagFile, error) {
	// Execute SQL
	item := models.BlogTagFile{}
	sql := "SELECT btf.id, btf.blogtag_id, btf.blogfile_id, bt.name, bf.filename FROM blogtagfile btf " +
		"INNER JOIN blogtag bt ON bt.id = btf.blogtag_id INNER JOIN blogfile bf ON bf.id = btf.blogfile_id ORDER BY id ASC LIMIT 1;"
	err := s.Conn.QueryRow(config.CTX, sql).Scan(
		&item.Id,
		&item.BlogtagId,
		&item.BlogfileId,
		&item.BlogtagName,
		&item.BlogfileFilename,
	)
	if err != nil {
		return item, err
	}

	// If success return item
	return item, nil
}

func (s *BlogTagFileService) GetLast() (models.BlogTagFile, error) {
	// Execute SQL
	item := models.BlogTagFile{}
	sql := "SELECT btf.id, btf.blogtag_id, btf.blogfile_id, bt.name, bf.filename FROM blogtagfile btf " +
		"INNER JOIN blogtag bt ON bt.id = btf.blogtag_id INNER JOIN blogfile bf ON bf.id = btf.blogfile_id ORDER BY id DESC LIMIT 1;"
	err := s.Conn.QueryRow(config.CTX, sql).Scan(
		&item.Id,
		&item.BlogtagId,
		&item.BlogfileId,
		&item.BlogtagName,
		&item.BlogfileFilename,
	)
	if err != nil {
		return item, err
	}

	// If success return item
	return item, nil
}

func (s *BlogTagFileService) GetAll(limit *int, page *int) ([]models.BlogTagFile, error) {
	// Execute SQL
	sql := "SELECT * FROM get_all_blogtagfile(@limit, @page);"
	args := pgx.NamedArgs{
		"limit": *limit,
		"page":  *page,
	}
	rows, err := s.Conn.Query(config.CTX, sql, args)
	if err != nil {
		return nil, err
	}

	// Convert SQL rows to data
	data := []models.BlogTagFile{}
	for rows.Next() {
		model := models.BlogTagFile{}

		if err := rows.Scan(
			&model.Id,
			&model.BlogtagId,
			&model.BlogfileId,
			&model.BlogtagName,
			&model.BlogfileFilename,
		); err != nil {
			return nil, err
		}

		data = append(data, model)
	}

	// If success return nil
	return data, nil
}
