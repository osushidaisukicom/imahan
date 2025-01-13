# imahan

完成したら今半に行く

## Usage

### setup

```shellsession
$ cp .envrc.example .envrc
```

```shellsession
# fill secrets
$ vim .envrc
```

```shellsession
$ source .envrc
```

### init db

```shellsession
$ psql -h $DB_HOST -p $DB_PORT -d $DB_NAME -U $DB_USER -f ./initdb.d/00_create_task_table.sql
```

### serve

```shellsession
$ go run ./cmd/imahan-api/main.go
```

### use SQLBoiler

/models にコードが自動生成される。
fyi: https://github.com/volatiletech/sqlboiler

```shellsession
$ sqlboiler psql -o models -p models --no-tests --wipe
```

### API

- GET `/task`
  - get task list
- POST `/task`
  - create new task
