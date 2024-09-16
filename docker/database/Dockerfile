FROM postgres:16-alpine

# Install additional tools (e.g., pg_ctl) using apk
RUN apk update && \
    apk add --no-cache postgresql-client go graphviz openjdk8

# Define build argument
ARG ENVIRONMENT

# Documentation schemaspy
WORKDIR /app/schemaspy
COPY schemaspy/postgresql-42.7.3.jar schemaspy/schemaspy-6.2.4.jar ./

# Copy the appropriate setup file based on the environment
COPY database/src /docker-entrypoint-initdb.d/src

# Setup database dir
WORKDIR /app

# Use a script to determine which setup file to copy
COPY database/copy-setup-file.sh ./
COPY database/setup-database-stable.sql ./
COPY database/setup-database-dev.sql ./

# Run the script to copy the correct setup file
RUN chmod +x ./copy-setup-file.sh
RUN ./copy-setup-file.sh ${ENVIRONMENT}

# Back to test dir
WORKDIR /app/test
