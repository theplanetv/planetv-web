CREATE OR REPLACE FUNCTION create_blogcategory(
        input_name TEXT
    )
    RETURNS VOID
    AS $$
    BEGIN
        INSERT INTO blogcategory (name)
        VALUES (input_name);
    END;
    $$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION update_blogcategory(
        input_id INT,
        input_name TEXT
    )
    RETURNS VOID
    AS $$
    BEGIN
        UPDATE blogcategory
        SET name = input_name
        WHERE id = input_id;
    END;
    $$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION remove_blogcategory(
        input_id INT
    )
    RETURNS VOID
    AS $$
    BEGIN
        DELETE FROM blogcategory
        WHERE id = input_id;
    END;
    $$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION count_blogcategory()
    RETURNS INT
    AS $$
    DECLARE
        value_count INT;
    BEGIN
        SELECT COUNT(id) INTO value_count FROM blogcategory;

        RETURN value_count;
    END;
    $$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION get_all_blogcategory(
        input_limit INT,
        input_page  INT
    )
    RETURNS TABLE (
        id   INT,
        name TEXT
    )
    AS $$
    DECLARE
        value_count  INT;
        max_page     INT;
        value_offset INT;
    BEGIN
        -- Get count
        SELECT * INTO value_count FROM count_blogcategory();

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
            SELECT b.id, b.name
            FROM blogcategory b
            ORDER BY b.id
            LIMIT input_limit
            OFFSET value_offset;
    END;
    $$ LANGUAGE plpgsql;
