# Expensify-API (WIP)

![Go](https://github.com/dsouzadyn/expensify-go-api/workflows/Go/badge.svg)

The backend api for Expensify(an expense management app I am working on).
This is also an exercise for me to flex my go-lang skills ;)

# DB Setup

Just run the sql file in MYSQL. I use MYSQL workbench.

Add the ```MYSQL_DB_DSN``` to ```.env``` file. You can take a look at the ```.env.example``` example file.

# Go Setup

1. Install the modules
```sh
$ go mod tidy
```

2. Build the backages
```sh
$ go build ./... # the 3 dots are not a typo
```

3. Run
```sh
$ go run main.go
```
