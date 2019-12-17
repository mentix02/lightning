package main

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

func getDB() *sql.DB {
	db, err := sql.Open("mysql", "root:toor@/medialist")
	check(err)
	return db
}

var db = getDB()

// Returns a list containing all the ids of articles that
// the given username's Author has bookmarked. It uses a list
// since there's a lot of appends from generated from db.Next().
// This method should only be called once the request has been
// properly authenticated by some form of Token (header based)
// authenticated - this would make sure that the Author with
// the provided username exists.
func ArticleIdsSortedByAuthorBookmarks(username string) List {

	// Build query.
	q, err := db.Prepare("SELECT `bookmark_bookmark`.`article_id` FROM " +
								   "`bookmark_bookmark` INNER JOIN `author_author`" +
						  		   " ON (`bookmark_bookmark`.`author_id` = " +
						  		   "`author_author`.`id`) WHERE `author_author`.`username`" +
						  		   " = ? ORDER BY `bookmark_bookmark`.`id` DESC;")
	check(err)

	rows, err := q.Query(username)

	if err != nil {
		return List{}
	}

	columns, err := rows.Columns()
	check(err)

	// To hold empty bytes and scan the row
	// values into scanArgs.
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))

	// To store all the primary keys.
	pks := List{}

	// Point scanArgs to values' index.
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Loop over results.
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		check(err)
		pk, _ := strconv.Atoi(string(values[0]))
		pks.append(uint32(pk))
	}

	if err = rows.Err(); err != nil {
		panic(err.Error())
	}

	return pks
}

func getAuthorUsernameFromKey(key string) (string, error) {

	// Prepare query statement.
	q, err := db.Prepare("SELECT `author_author`.`username` FROM `authtoken_token` INNER" +
						   		" JOIN `author_author` ON (`authtoken_token`.`user_id` = `author_author`.`id`)" +
						   		"WHERE `authtoken_token`.`key` = ?")
	check(err)

	var username string

	// Execute query and load result in
	// the username variable above.
	err = q.QueryRow(key).Scan(&username)

	if err != nil {
		return "", errors.New("Invalid credentials.")
	} else {
		return username, nil
	}

}

func getTokenFromUsername(username string) (string, error) {
	q, err := db.Prepare("SELECT `authtoken_token`.`key` FROM `authtoken_token` INNER JOIN `author_author` " +
								" ON (`authtoken_token`.`user_id` = `author_author`.`id`) WHERE" +
								" `author_author`.`username` = ?")
	check(err)
	var token string

	err = q.QueryRow(username).Scan(&token)

	if err != nil {
		return "", errors.New("Invalid credentials.")
	} else {
		return token, nil
	}

}

func getHashedPasswordFromUsername(username string) (string, error) {
	q, err := db.Prepare("SELECT `author_author`.`password` FROM `author_author` WHERE `author_author`.`username`" +
								" = ? ORDER BY `author_author`.`date_joined` DESC, `author_author`.`id` DESC")
	check(err)
	var password string

	err = q.QueryRow(username).Scan(&password)

	if err != nil {
		return "", errors.New("Invalid credentials.")
	} else {
		return password, nil
	}

}

func authorUsernameExists(username string) bool {
	q, err := db.Prepare("SELECT `author_author`.`id` FROM `author_author` WHERE" +
								" `author_author`.`username` = ? ORDER BY `author_author`.`date_joined`" +
								"DESC, `author_author`.`id` DESC")
	check(err)

	var exists bool

	err = q.QueryRow(username).Scan(&exists)

	if err != nil {
		return false
	} else {
		return exists
	}
}
