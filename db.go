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
func articleIdsSortedByAuthorBookmarks(username string) List {

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
	}
	return username, nil
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
	}
	return token, nil
}

func getHashedPasswordFromUsername(username string) (string, error) {
	q, err := db.Prepare("SELECT `author_author`.`password` FROM `author_author` WHERE `author_author`.`username`" +
		" = ? ORDER BY `author_author`.`date_joined` DESC, `author_author`.`id` DESC")
	check(err)
	var password string

	err = q.QueryRow(username).Scan(&password)

	if err != nil {
		return "", errors.New("Invalid credentials.")
	}
	return password, nil
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
	}
	return exists
}

func getAuthorUsernameFromId(id uint32) (string, error) {
	q, err := db.Prepare("SELECT `author_author`.`username` FROM `author_author` WHERE +" +
		"`author_author`.`id` = ?")
	check(err)

	var username string

	err = q.QueryRow(id).Scan(&username)

	if err != nil {
		return "", errors.New("Author not found.")
	}
	return username, nil
}

func getTagsFromArticleId(id uint32) []string {
	q, err := db.Prepare("SELECT DISTINCT `taggit_tag`.`name` FROM `taggit_tag`" +
		" INNER JOIN `taggit_taggeditem` ON (`taggit_tag`.`id` = " + "`taggit_taggeditem`.`tag_id`)" +
		" INNER JOIN `django_content_type` ON (`taggit_taggeditem`.`content_type_id` = `django_content_type`.`id`)" +
		" WHERE (`django_content_type`.`app_label` = \"article\" AND `django_content_type`.`model` = \"article\"" +
		" AND `taggit_taggeditem`.`object_id` = ?)")
	check(err)

	rows, err := q.Query(id)

	if err != nil {
		return []string{}
	}

	// To hold empty bytes and scan the row
	// values into scanArgs.
	values := make([]sql.RawBytes, 1)
	scanArgs := make([]interface{}, len(values))

	// To store all the tags.
	tags := List{}

	// Point scanArgs to values' index.
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Loop over results.
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		check(err)
		tags.append(string(values[0]))
	}

	if err = rows.Err(); err != nil {
		panic(err.Error())
	}

	return tags.toStringSlice()
}

func getTopicNameFromId(id uint32) (string, error) {
	q, err := db.Prepare("SELECT `topic_topic`.`name` FROM `topic_topic` WHERE +" +
		"`topic_topic`.`id` = ?")
	check(err)

	var name string

	err = q.QueryRow(id).Scan(&name)

	if err != nil {
		return "", errors.New("Topic not found.")
	}
	return name, nil
}

func getRecentArticles() List {
	rows, err := db.Query("SELECT `article_article`.`id`, " +
		"`article_article`.`title`, `article_article`.`content`, " +
		" `article_article`.`created_on`, `article_article`.`topic_id`, `article_article`.`author_id`, " +
		" `article_article`.`slug`, `article_article`.`objectivity`," +
		" `article_article`.`thumbnail_url` FROM `article_article` WHERE `article_article`.`draft` =" +
		" False ORDER BY `article_article`.`created_on` DESC, `article_article`.`updated_on` DESC, " +
		" `article_article`.`id` DESC LIMIT 12")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	articles := List{}

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// Serialize fields properly.
		id, _ := strconv.Atoi(string(values[0]))
		tags := getTagsFromArticleId(uint32(id))
		content := string(values[2])[:150] + "..."
		topicId, _ := strconv.Atoi(string(values[4]))
		authorID, _ := strconv.Atoi(string(values[5]))
		topic, _ := getTopicNameFromId(uint32(topicId))
		objectivity, _ := strconv.Atoi(string(values[7]))
		author, _ := getAuthorUsernameFromId(uint32(authorID))

		article := Article{
			Tags:        tags,
			Topic:       topic,
			Author:      author,
			Content:     content,
			ID:          uint32(id),
			Title:       string(values[1]),
			Slug:        string(values[6]),
			Timestamp:   string(values[3]),
			Thumbnail:   string(values[8]),
			Objectivity: int64(objectivity),
		}

		articles.append(article)

	}

	return articles
}
