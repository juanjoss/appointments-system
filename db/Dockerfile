FROM postgres:14

COPY /db/init.sql /docker-entrypoint-initdb.d/1.sql

CMD ["postgres"]