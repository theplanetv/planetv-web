CREATE OR REPLACE FUNCTION create_blogtag(
        input_blogcategory_id INT,
        input_name            TEXT
    )
    RETURNS VOID
    AS $$
    BEGIN
        INSERT INTO blogtag (blogcategory_id, name)
        VALUES (input_blogcategory_id, input_name);
    END;
    $$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION update_blogtag(
        input_id              INT,
        input_blogcategory_id INT,
        input_name            TEXT
    )
    RETURNS VOID
    AS $$
    DECLARE
        sql_query            TEXT := 'UPDATE blogtag SET ';
        flag_blogcategory_id BOOLEAN := false;
        flag_name            BOOLEAN := false;
    BEGIN
        -- Check id
        IF input_id IS NULL THEN
            RAISE EXCEPTION 'Must have id!';
        END IF;

        -- Build the SET clause based on non-NULL input values
        IF input_blogcategory_id IS NOT NULL OR input_blogcategory_id > 0 THEN
            sql_query := sql_query || 'blogcategory_id = ' || input_blogcategory_id || ',';
            flag_blogcategory_id = true;
        END IF;

        IF input_name IS NOT NULL OR input_name != '' THEN
            sql_query := sql_query || 'name = ' || quote_literal(input_name) || ',';
            flag_name = true;
        END IF;

        -- Remove trailing comma and space from the SET clause
        IF length(sql_query) > 0 THEN
            sql_query := left(sql_query, length(sql_query) - 1);
        END IF;

        -- Only append WHERE clause if there are changes
        IF flag_blogcategory_id = true OR flag_name = true THEN
            sql_query := sql_query || ' WHERE id = ' || input_id || ';';
            EXECUTE sql_query;
        ELSE
            RAISE EXCEPTION 'No fields to update!';
        END IF;
    END;
    $$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION remove_blogtag(
        input_id INT
    )
    RETURNS VOID
    AS $$
    BEGIN
        DELETE FROM blogtag
        WHERE id = input_id;
    END;
    $$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION count_blogtag()
    RETURNS INT
    AS $$
    DECLARE
        value_count INT;
    BEGIN
        SELECT COUNT(id) INTO value_count FROM blogtag;

        RETURN value_count;
    END;
    $$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION get_all_blogtag(
        input_limit INT,
        input_page  INT
    )
    RETURNS TABLE (
        id                INT,
        blogcategory_id   INT,
        name              TEXT,
        blogcategory_name TEXT
    )
    AS $$
    DECLARE
        value_count  INT;
        max_page     INT;
        value_offset INT;
    BEGIN
        SELECT * INTO value_count FROM count_blogtag();

        -- Assign default limit
        IF input_limit < 5 THEN
            input_limit := 5;
        ELSIF input_limit > 50 THEN
            input_limit := 50;
        END IF;

        -- Assign default page
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
            SELECT bt.id, bt.blogcategory_id, bt.name, bc.name
            FROM blogtag bt
            INNER JOIN blogcategory bc ON bc.id = bt.blogcategory_id
            ORDER BY bt.id
            LIMIT input_limit
            OFFSET value_offset;
    END;
    $$ LANGUAGE plpgsql;
