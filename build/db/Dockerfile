FROM postgres:11.5-alpine

RUN mkdir -p /docker-entrypoint-initdb.d
ADD init.sql /docker-entrypoint-initdb.d/init.sql

EXPOSE 5432

CMD ["postgres"]