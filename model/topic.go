package model

import "errors"
import "database/sql"
import "strconv"
import "strings"

type Topic struct {
	Id          int
	Title       string
	Description string
	ForumId     int
	PostCount   int
}

func NewTopic() *Topic {
	return &Topic{-1, "", "", -1, -1}
}

func ValidateTopic(topic *Topic) (ok bool, errs []error) {
	errs = make([]error, 0)

	trimmedTitle := strings.TrimSpace(topic.Title)
	trimmedDescription := strings.TrimSpace(topic.Description)

	if trimmedTitle == "" {
		errs = append(errs, errors.New("Topic must have a title."))
	}

	if len(trimmedTitle) > 255 {
		errs = append(errs, errors.New("Topic title is too long."))
	}

	if len(trimmedDescription) > 255 {
		errs = append(errs, errors.New("Topic description is too long."))
	}

	// todo, check for valid forum id with database query
	if topic.ForumId == -1 {
		errs = append(errs, errors.New("Post must belong to a valid topic."))
	}

	return len(errs) == 0, errs
}

func SaveTopic(db *sql.DB, topic *Topic) error {
	_, err := db.Exec("INSERT INTO topics (id, title, description, forum_id) VALUES (NULL,?,?,?)", topic.Title, topic.Description, topic.ForumId)
	return err
}

func FindOneTopic(db *sql.DB, reqId string) (Topic, error) {
	var (
		id          int
		title       string
		description string
		forumId     int
	)

	row := db.QueryRow("SELECT * FROM topics WHERE id = ?", reqId)
	err := row.Scan(&id, &title, &description, &forumId)
	if err != nil {
		return Topic{}, errors.New("could not query for topic with id " + reqId)
	}

	return Topic{id, title, description, forumId, 0}, nil
}

func postCount(db *sql.DB, reqId string) (int, error) {
	var count int

	row := db.QueryRow("SELECT count(*) FROM posts WHERE topic_id = ?", reqId)
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func FindTopics(db *sql.DB, reqId string) ([]Topic, error) {
	rows, err := db.Query("SELECT * FROM topics WHERE forum_id = ?", reqId)
	defer rows.Close()
	if err != nil {
		return nil, errors.New("could not query for topics for fourm " + reqId)
	}

	topics := make([]Topic, 0)
	for rows.Next() {
		var (
			id          int
			title       string
			description string
			forumId     int
		)

		err := rows.Scan(&id, &title, &description, &forumId)
		if err != nil {
			return nil, errors.New("could not process row")
		}

		postCount, err := postCount(db, strconv.Itoa(id))
		if err != nil {
			return nil, errors.New("could not could topics for forum with id " + strconv.Itoa(id))
		}

		topics = append(topics, Topic{id, title, description, forumId, postCount})
	}

	return topics, nil
}
