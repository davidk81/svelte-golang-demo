# 
# docker-compose up -d

#
docker run -v `pwd`/migrations:/migrations --network host migrate/migrate:v4.12.1 \
    -path=/migrations/ -database "postgres://docker:docker@localhost:5432/patientdb?sslmode=disable" $*
