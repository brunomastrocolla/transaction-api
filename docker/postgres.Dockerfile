FROM postgres:12-alpine
COPY ../db/init/init.sql /docker-entrypoint-initdb.d/
