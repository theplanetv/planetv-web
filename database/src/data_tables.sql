CREATE TABLE public.blogcategory (
    id   SERIAL PRIMARY KEY CHECK (id >= 1),
    name TEXT
);

CREATE TABLE public.blogtag (
    id              SERIAL PRIMARY KEY CHECK (id >= 1),
    blogcategory_id INT,
    name            TEXT,
    CONSTRAINT fk_blogcategory_for_blogtag
        FOREIGN KEY (blogcategory_id)
        REFERENCES public.blogcategory(id)
);

CREATE TABLE public.blogfile (
    id           BIGSERIAL PRIMARY KEY CHECK (id >= 1),
    filename     TEXT,
    created_date DATE DEFAULT NOW(),
    updated_date DATE DEFAULT NOW()
);

CREATE TABLE public.blogtagfile (
    id          BIGSERIAL PRIMARY KEY CHECK (id >= 1),
    blogtag_id  INT,
    blogfile_id BIGINT,
    CONSTRAINT fk_blogtag_for_blogtagfile
        FOREIGN KEY (blogtag_id)
        REFERENCES public.blogtag(id),
    CONSTRAINT fk_blogfile_for_blogtagfile
        FOREIGN KEY (blogfile_id)
        REFERENCES public.blogfile(id)
);
