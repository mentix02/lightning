# Lightning Installation

A key factor for me to install an open source tool is that it should be easy to install. It plays a huge psychological
role in getting people to contribute and join a project. This is why the installation section in the [`docs/README.md`](README.md)
is so small. While the steps provided in the former will work flawlessly if you have a proper Go environment set up, it still
misses out on some key database set up instructions. Here's a definitive guide.

## Table of Contents

1. [Get The Code](#get-the-code)
2. [Set Up The Database](#set-up-the-database)
3. [Migrating Database](#migrating-database)
4. [Running Tests (optional)](#running-tests)

## Get The Code

All of Lightning's development was done on a mix of systems running Linux (mostly Ubuntu) and MacOS - in other words,
support for a Windows development environment can not be expected in the future. Feel free to use any *nix system of your
choice - the easiest to work with was [Ubuntu](https://ubuntu.com/).

Once you're ready with your *nix OS of choice, create a directory in your Desktop or home folder to host the 
[medialist-backend](https://github.com/mentix02/medialist-backend).

Why make an entire directory instead of just cloning the repo directly? Because the Medialist has a grand total of three 
parts to it (as of now) - the backend, the frontend, and lightning. So if you need to have the frontend as well, it would
be good practise to keep them the backend as well in the same directory. There's no compulsion to clone the frontend repo
- that's the beauty of a RESTful project. Once you have this directory, clone the backend repo into a directory named `backend`.

**Note for Go developers - don't make this directory in `$GOPATH` since lightning is an standalone Go project that does
not work in _conjunction_ with the backend code which is in Python.**

```sh
$ mkdir medialist
$ git clone https://github.com/mentix02/medialist-backend medialist/backend
```

**Note - skip this step if you don't want lightning on your machine and have been redirected from the medialist-backend
documentation and ergo are only interested in running the Django server on your machine OR already have a Go environment
set up and ready to go.**

Now you have the backend code on your machine. But before migrating or creating a database, let's set up your Go environment
since that's what lightning is written in. First steps first, [install Go](https://golang.org/doc/install). Now create an
empty directory in your home folder - 

```sh
$ mkdir ~/go
```

This is your `$GOPATH` aka the place where everything Go related will live. To get lightning in your `$GOPATH`, you **can**
use `go get` but that isn't ideal since it'd store lightning in a subdirectory with path - `$GOPATH/src/github.com/mentix02/lightning`.
That's just ugly. Instead, manually make a `src` directory inside `$GOPATH` and `git clone` lightning - 

```sh
$ mkdir ~/go/src && cd ~/go/src/
$ git clone https://github.com/mentix02/lightning
```

That's all that it takes to set up your backend and lightning codebase in a proper environment. 

## Set Up The Database

The Medialist uses good ol' MySQL as its database. Postgres was considered but it's much more difficult to set up as compared
to MySQL and hence was dropped. Installing MySQL is simple. For Ubuntu/Debian based OSes, simply use `apt` to install it
along with some Python libraries to work with MySQL.

```sh
$ sudo apt install mysql-server python3-dev default-libmysqlclient-dev libssl-dev mysql-client
```

Now install the requirements. It's good practise to keep all your requirements in a virtualenv environment but that's up
to you.

```sh
$ cd <path to medialist backend>
$ pip install -r requirements.txt
```

Good job! We're nearly there. The only step that remains is to actually configure the MySQL database and migrate our, well,
migrations. Enter the following command and when prompted for a password, use `toor`. For all other options, use
[this file](mysql_secure_installation.txt) to enter the correct options.

```sh
$ sudo mysql_secure_installation
...
```

## Migrating Database

Congratulations! By this point, you've got all the dependencies required for the Medialist and are ready to migrate your
local copy of the project's schemas to the actual MySQL database. Navigation to the `backend` directory of The Medialist
and simply run `./manage.py migrate`. But just ONE small step before that - creating the actual Medialist database in MySQL.
Log in to the MySQL server with the root credentials and enter the command `CREATE DATABASE medialist;`.

```sh
$ mysql -u root -p
Enter password:
Welcome to the MySQL monitor.  Commands end with ; or \g.

...

mysql> CREATE DATABASE medialist;
Query OK, 1 row affected (0.00 sec)

mysql> SHOW DATABASES;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| medialist          |
| mysql              |
| performance_schema |
| sys                |
+--------------------+
5 rows in set (0.01 sec)
```

That's it. Just run migrate on the database command and you're good to go.

```sh
$ ./manage.py migrate
```

## Running Tests

This step is not at all required but is highly recommended. Unit testing is an important part for any software project &
the same can be said for the Medialist. To run the Python tests, navigate to the backend directory and run -

```sh
$ ./manage.py test
```

To run the Go tests, go to the lightning directory, and run -

```sh
go test
```
