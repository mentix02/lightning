# Lightning
![CircleCI](https://img.shields.io/circleci/build/github/mentix02/lightning/master)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/mentix02/lightning)
[![GitHub license](https://img.shields.io/github/license/mentix02/lightning)](https://github.com/mentix02/lightning/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/mentix02/lightning)](https://goreportcard.com/report/github.com/mentix02/lightning)
[![Run on Repl.it](https://repl.it/badge/github/mentix02/lightning)](https://repl.it/github/mentix02/lightning)

A **simple** & blazingly fast API server for The Medialist.

Some highly critical function calls that require to be fast live over here. This project was started due to the slow
response times achieved from the views from the server running Django. Since the majority of the backend is written in
Python, there's no point in re-writing the entire codebase but for the views that perform **simple** retrievals (with proper 
authentication, of course), Go seems to be an apt fit to implement them in. The pattern is a **simple** "model controller"
pattern since the only models we're working with are predefined from Django's ORM and we only return
JSON responses.

The server used is provided by Go's built in `http` package and that lives inside [`main.go`](main.go). All the handlers, 
aka the views, are written in [`handlers.go`](handlers.go).
The router used is the more than capable [mux](https://github.com/gorilla/mux) router. All the model operations
(mostly retrievals) live inside [`db.go`](db.go). There's some special data types that Go does not provide a good interface
for and thus they are all stored inside [`structs.go`](structs.go) - examples include a singly linked list for storing
primary keys from database fetches.

All the handlers return JSON and thus there are no templates to work with. The frontend will most likely be served from
the Django server but it might be implemented in Go in the future.

Since the project is mostly bare-bones Go with very few dependencies, no package manager is used. Proper documentation
is in development and will be coming soon.

## Documentation

Documentation can be found [here](docs).

## License

Lightning is licensed under the MIT license.

## Contributing

Lightning is a highly critical project and is, as of now, not accepting any pull requests since there is no formal
contributing guides or a CI endpoint available. But they're both in development and will soon be available.

## Further Reading

If you're here, check out the main backend written in Python - [medialist-backend](https://github.com/mentix02/medialist-backend).
