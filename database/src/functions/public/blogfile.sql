CREATE OR REPLACE FUNCTION create_blogfile(
        input_filename TEXT
    )
    RETURNS VOID
    AS $$
    BEGIN
        INSERT INTO blogfile (filename)
        VALUES (input_filename);
    END;
    $$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION update_blogfile(
        input_id           BIGINT,
        input_filename     TEXT,
        input_created_date DATE,
        input_updated_date DATE
    )
    RETURNS VOID
    AS $$
    DECLARE
        sql_query         TEXT    := 'UPDATE blogfile SET ';
        flag_filename     BOOLEAN := false;
        flag_created_date BOOLEAN := false;
        flag_updated_date BOOLEAN := false;
    BEGIN
        -- Check id
        IF input_id IS NULL THEN
            RAISE EXCEPTION 'Must have id!';
        END IF;

        -- Build the SET clause based on non-NULL input values
        IF input_filename IS NOT NULL OR input_filename != '' THEN
            sql_query := sql_query || 'filename = ' || quote_literal(input_filename) || ',';
            flag_filename = true;
        END IF;

        IF input_created_date IS NOT NULL THEN
            sql_query := sql_query || 'created_date = ' || quote_literal(input_created_date) || ',';
            flag_created_date = true;
        END IF;

        IF input_updated_date IS NOT NULL THEN
            sql_query := sql_query || 'updated_date = ' || quote_literal(input_updated_date) || ',';
            flag_updated_date = true;
        END IF;

        -- Remove trailing comma and space from the SET clause
        IF length(sql_query) > 0 THEN
            sql_query := left(sql_query, length(sql_query) - 1);
        END IF;

        -- Only append WHERE clause if there are changes
        IF flag_filename = true OR flag_created_date = true OR flag_updated_date = true THEN
            sql_query := sql_query || ' WHERE id = ' || input_id || ';';
            EXECUTE sql_query;
        ELSE
            RAISE EXCEPTION 'No fields to update!';
        END IF;
    END;
    $$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION remove_blogfile(
        input_id BIGINT
    )
    RETURNS VOID
    AS $$
    BEGIN
        DELETE FROM blogfile
        WHERE id = input_id;
    END;
    $$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION count_blogfile()
    RETURNS BIGINT
    AS $$
    DECLARE
        value_count BIGINT;
    BEGIN
        SELECT COUNT(id) INTO value_count FROM blogfile;

        RETURN value_count;
    END;
    $$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION get_all_blogfile(
        input_limit INT,
        input_page  INT
    )
    RETURNS TABLE (
        id           BIGINT,
        filename     TEXT,
        created_date DATE,
        updated_date DATE
    )
    AS $$
    DECLARE
        value_count  BIGINT;
        max_page     INT;
        value_offset INT;
    BEGIN
        -- Get count
        SELECT * INTO value_count FROM count_blogfile();

        -- Assign default limit
        IF input_limit < 5 THEN
            input_limit := 5;
        ELSIF input_limit > 50 THEN
            input_limit := 50;
        END IF;

        -- Assign max_page to input_page
        max_page := CEIL(value_count::NUMERIC / input_limit);
        IF input_page > max_page THEN
            input_page := max_page;
        END IF;

        -- Caculate value_offset
        value_offset := input_limit * (input_page - 1);

        -- value_offset must not negative
        IF value_offset < 0 THEN
            value_offset := 0;
        END IF;

        -- Return query table
        RETURN QUERY
            SELECT b.id, b.filename, b.created_date, b.updated_date
            FROM blogfile b
            ORDER BY b.id
            LIMIT input_limit
            OFFSET value_offset;
    END;
    $$ LANGUAGE plpgsql;
