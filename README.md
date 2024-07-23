### 1

Создать базу данных с помощью Docker:

```bash
docker-compose up -d
```

или вручную:

```bash
sudo apt update
sudo apt install postgresql postgresql-contrib

sudo systemctl start postgresql
sudo systemctl enable postgresql

sudo -i -u postgres

psql

CREATE DATABASE test;
CREATE USER test WITH PASSWORD 'test';
GRANT ALL PRIVILEGES ON DATABASE test TO test;
\q
```

### 2

```bash
make test-migration-up
```

### 3

```bash
go run cmd/http-app/main.go
```
