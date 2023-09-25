# AKO
Ausbildungskontrolle


docker run -d --name ako-db -p 5432:5432 -e POSTGRES_DB=ako -e POSTGRES_PASSWORD=secret -e POSTGRES_USER=postgres postgres

psql -h 0.0.0.0 -p 5432 -U postgres -d ako
