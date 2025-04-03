# Requirements
- Go 1.23+
- PostgreSQL
## Installation
You may install this cli application using `go install www.github.com/dabates/gator`

##Config File
You should create a `.gatorconfig.json` in your home directory to setup the cli
```json
{
    "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable",
    "current_user_name": ""
}
```

** This project is a assignment from boot.dev **