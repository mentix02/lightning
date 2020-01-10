# Lightning Documentation

Lightning is a blazing fast API server written for The Medialist. It's exists because the main Django server just isn't
fast enough for some calls that occur too frequently. Since most of the calls that Lightning handles are **simple**, no
ORM is used for them. Only a MySQL connector is used to perform raw SQL queries to fetch results.

Lightning isn't designed to handle complex (read: complicated) views such as user sign up or log in. It's only used where
raw performance can be extracted while handling **simple** queries. It is designed to be used in conjunction with the Django
backend and thus the frontend should not be able to discern any difference between the two services - Django and Lightning.

The task is simplified since the Django backend API uses JSON as its data format and Go has good support for handling
JSON data and other web related services such as routing.

## Table of Contents

1. [Installation](#installation)
2. [Dependencies](#dependencies)
3. [Testing](#testing)
    + [Go Tests](#go-tests)
    + [Python Tests](#python-tests)
4. [Lightning Components](#lightning-components)
    + [db.md](#dbmd)
    + [structs.md](#structsmd)

## Installation

There are no special instructions to get Lightning up and running. To get the code, simply run -

```sh
$ go get -u github.com/mentix02/lightning
```

To setup your database, install MySQL from your package manager and run migrations from the
[backend code](https://github.com/mentix02/medialist-backend) to the database named `medialist`.

To build (read: compile) the code, run - 

```sh
$ go build -i
```

For further info, refer [`INSTALLATION.md`](INSTALLATION.md).

## Dependencies

Lightning was designed from scratch to be **simple**. That is one of the reasons why it doesn't have a lot of dependencies.
A short compiled list can be found below - 

1. [mux](https://github.com/gorilla/mux) - a **simple** to use yet robust router and URL matcher for Go.
2. [go-mysql-driver](https://github.com/go-sql-driver/mysql) - a light weight MySQL driver for Go's `database/sql` package.
3. [unchained](https://github.com/alexandrevicenzi/unchained) - password hashing and validation from Django for Go.

So far, that's it and we hope it remains that way. Lightning will never use a _proper_ web framework since that would
destroy the purpose of keeping the codebase **simple**. Since the dependency list is so small, no package managers have been used.

## Testing

Writing and running unit tests is an integral part for any software project - big or small. And while lightning is certainly
not a giant code base, tests are written for proper functioning. When starting out the project, Python was picked to be
a _good enough_ testing suite with it's, admittedly, superior support for testing with its [`unittest`](https://docs.python.org/3/library/unittest.html)
package. Later, it was decided that Go's `testing` package could also be used for testing **simpler** functions natively 
implemented in Go. All the tests written in Go are of the format `*_test.go` and all the tests in Python live inside the
[`tests`](../tests) directory. So how to run and test the code? Well, the answer's a little more complicated than it needs
to be.

### Go Tests

Go tests are **simple**... "usage examples" for the lack of a better word. Since there's no inbuilt assertions library,
all the files that match the pattern `*_test.go` are simply calling the functions that live in the file that the `*` prefix
references. To run the tests, no database setup is required and is just - 

```sh
$ go test
```

### Python Tests

Initially, only Python based tests existed and they were only for checking HTTP based results. Since Lightning is an API server,
we make use of Python's [`requests`](https://2.python-requests.org/en/master/) library for live server testing. It's all
powered by the vanilla [`unittest`](https://docs.python.org/3/library/unittest.html) module. Running the tests is as simple as
calling the command below but you need to have the [lightning server running with a prepared database](INSTALLATION.md).

```sh
$ cd tests && python -m unittest
```

**Note - make sure you have `requests` installed (`pip install requests`).**

## Lightning Components

As you might've observed by the highlighting of every occurrence of the word **simple**, Lightning was made with the
[KISS](https://en.wikipedia.org/wiki/KISS_principle) principle in mind. Go is a **simple** language. The project was made
to tackle **simple** problems. Thus, the source code should also be **simple**. The code has been written in an extremely
easy to read way with proper documentation (including inline comments) where it's required by following the official Go
style guide (enforced by [`gofmt`](https://golang.org/cmd/gofmt/)).

The [docs](.) directory contains markdown files with the same name as the source code files (without the `.go` extension).
Here's a definitive listing - 

### [db.md](db.md)

Describes the query operations that are performed on the MySQL database - akin to the `models.py` in Django.

### [structs.md](structs.md)

Commonly used data structures for storing data retrieved from the database and for returning JSON "detail" type responses.
