# Lightning Documentation

Lightning is a blazing fast API server written for The Medialist. It's exists because the main Django server just isn't
fast enough for some calls that occur too frequently. Since most of the calls that Lightning handles are **simple**, no
ORM is used for them. Only a MySQL connector is used to perform raw SQL queries to fetch results.

Lightning isn't designed to handle complex (read: complicated) views such as user sign up or log in. It's only used where
raw performance can be extracted while handling **simple** queries. It is designed to be used in conjunction with the Django
backend and thus the frontend should not be able to discern any difference between the two services - Django and Lightning.

The task is simplified since the Django backend API uses JSON as its data format and Go has good support for handling
JSON data and other web related services such as routing.

## Installation

There are no special instructions to get Lightning up and running. To get the code, simply run -

```sh
$ go get -u github.com/mentix02/lightning
```

To setup your database, install MySQL from your package manager and run migrations from the
[backend code](https://github.com/mentix02/medialist-backend) to the database named `medialist`.

To build (read: compile) the code, run - 

```sh
$ go -i build
```

## Dependencies

Lightning was designed from scratch to be **simple**. That is one of the reasons why it doesn't have a lot of dependencies.
A short compiled list can be found below - 

1. [mux](https://github.com/gorilla/mux) - a **simple** to use yet robust router and URL matcher for Go.
2. [go-mysql-driver](https://github.com/go-sql-driver/mysql) - a light weight MySQL driver for Go's `database/sql` package.
3. [unchained](https://github.com/alexandrevicenzi/unchained) - password hashing and validation from Django for Go.

So far, that's it and we hope it remains that way. Lightning will never use a _proper_ web framework since that would
destroy the purpose of keeping the codebase **simple**. Since the dependency list is so small, no package managers have been used.
