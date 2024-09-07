CREATE OR REPLACE FUNCTION create_blogtagfile(
        input_blogtag_id  INT,
        input_blogfile_id BIGINT
    )
    RETURNS VOID
    AS $$
    BEGIN
        INSERT INTO blogtagfile (blogtag_id, blogfile_id)
        VALUES (input_blogtag_id, input_blogfile_id);
    END;
    $$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION update_blogtagfile(
        input_id          BIGINT,
        input_blogtag_id  INT,
        input_blogfile_id BIGINT
    )
    RETURNS VOID
    AS $$
    DECLARE
        sql_query        TEXT := 'UPDATE blogtagfile SET ';
        flag_blogtag_id  BOOLEAN := false;
        flag_blogfile_id  BOOLEAN := false;
    BEGIN
        -- Check id
        IF input_id IS NULL THEN
            RAISE EXCEPTION 'Must have id!';
        END IF;

        -- Build the SET clause based on non-NULL input values
        IF input_blogtag_id IS NOT NULL OR input_blogtag_id > 0 THEN
            sql_query := sql_query || 'blogtag_id = ' || input_blogtag_id || ',';
            flag_blogtag_id = true;
        END IF;

        IF input_blogfile_id IS NOT NULL OR input_blogfile_id > 0 THEN
            sql_query := sql_query || 'blogfile_id = ' || quote_literal(input_blogfile_id) || ',';
            flag_blogfile_id = true;
        END IF;

        -- Remove trailing comma and space from the SET clause
        IF length(sql_query) > 0 THEN
            sql_query := left(sql_query, length(sql_query) - 1);
        END IF;

        -- Only append WHERE clause if there are changes
        IF flag_blogtag_id = true OR flag_blogfile_id = true THEN
            sql_query := sql_query || ' WHERE id = ' || input_id || ';';
            EXECUTE sql_query;
        ELSE
            RAISE EXCEPTION 'No fields to update!';
        END IF;
    END;
    $$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION remove_blogtagfile(
        input_id BIGINT
    )
    RETURNS VOID
    AS $$
    BEGIN
        DELETE FROM blogtagfile
        WHERE id = input_id;
    END;
    $$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION count_blogtagfile()
    RETURNS BIGINT
    AS $$
    DECLARE
        value_count BIGINT;
    BEGIN
        SELECT COUNT(id) INTO value_count FROM blogtagfile;

        RETURN value_count;
    END;
    $$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION get_all_blogtagfile(
        input_limit INT,
        input_page  INT
    )
    RETURNS TABLE (
        id                BIGINT,
        blogtag_id        INT,
        blogfile_id       BIGINT,
        blogtag_name      TEXT,
        blogfile_filename TEXT
    )
    AS $$
    DECLARE
        value_count  BIGINT;
        max_page     INT;
        value_offset INT;
    BEGIN
        SELECT * INTO value_count FROM count_blogtagfile();

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
            SELECT
                btf.id, btf.blogtag_id, btf.blogfile_id,
                bt.name, bf.filename
            FROM blogtagfile btf
            INNER JOIN blogtag bt ON bt.id = btf.blogtag_id
            INNER JOIN blogfile bf ON bf.id = btf.blogfile_id
            ORDER BY bt.id
            LIMIT input_limit
            OFFSET value_offset;
    END;
    $$ LANGUAGE plpgsql;

-- Trigger functions
CREATE OR REPLACE FUNCTION prevent_if_same_blogtagfile()
    RETURNS TRIGGER
    AS $$
    DECLARE
        value_count SMALLINT;
    BEGIN
        -- Get count where blogtag_id = NEW.blogtag_id and blogfile_id = NEW.blogfile_id
        SELECT COUNT(b.id) INTO value_count
        FROM blogtagfile b
        WHERE b.blogtag_id = NEW.blogtag_id
        AND b.blogfile_id = NEW.blogfile_id;
   
        IF (value_count > 0) THEN
            RAISE EXCEPTION 'There is another value exist in blogtagfile!';
        END IF;

        RETURN NEW;
    END;
    $$ LANGUAGE plpgsql;
