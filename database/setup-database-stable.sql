CREATE DATABASE planetv;

\c planetv;

-- Tạo các bảng trong schema public
\i /docker-entrypoint-initdb.d/src/data_tables.sql
\i /docker-entrypoint-initdb.d/src/functions/public/blogcategory.sql
\i /docker-entrypoint-initdb.d/src/functions/public/blogfile.sql
\i /docker-entrypoint-initdb.d/src/functions/public/blogtag.sql
\i /docker-entrypoint-initdb.d/src/functions/public/blogtagfile.sql
