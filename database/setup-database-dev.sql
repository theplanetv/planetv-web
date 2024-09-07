CREATE DATABASE planetv;

\c planetv;

-- Create tables
\i /docker-entrypoint-initdb.d/src/data_tables.sql

-- Create functions
\i /docker-entrypoint-initdb.d/src/functions/public/blogcategory.sql
\i /docker-entrypoint-initdb.d/src/functions/public/blogfile.sql
\i /docker-entrypoint-initdb.d/src/functions/public/blogtag.sql
\i /docker-entrypoint-initdb.d/src/functions/public/blogtagfile.sql

-- Create triggers
\i /docker-entrypoint-initdb.d/src/triggers/public/blogtagfile.sql

\i /docker-entrypoint-initdb.d/src/insert-test-data.sql
