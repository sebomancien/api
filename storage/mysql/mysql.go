package mysql

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/sebomancien/api/logger"
)

type Store struct {
	db *sql.DB
}

func NewStore(cfg mysql.Config) *Store {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	logger.LogInfo("Connected to MySQL!")

	return &Store{
		db: db,
	}
}

func (s *Store) Init() error {
	// initialize the tables
	if err := s.createUsersTable(); err != nil {
		return err
	}

	return nil
}

func (s *Store) createUsersTable() error {
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INT UNSIGNED NOT NULL AUTO_INCREMENT,
			email VARCHAR(255) NOT NULL,
			firstName VARCHAR(255) NOT NULL,
			lastName VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

			PRIMARY KEY (id),
			UNIQUE KEY (email)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	`)

	return err
}
