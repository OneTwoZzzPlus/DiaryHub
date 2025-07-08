package postgresql

import (
	"context"
	"database/sql"
	"diaryhub/sso-service/internal/domain/models"
	"diaryhub/sso-service/internal/storage"
	"errors"
	"fmt"
	"time"

	pq "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func New(connStr string) (*Storage, error) {
	const op = "storage.postgresql.New"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// TODO: replace with migrations
	query := `
        CREATE TABLE IF NOT EXISTS users (
            id          SERIAL PRIMARY KEY,
            email       TEXT NOT NULL UNIQUE,
            pass_hash   BYTEA NOT NULL,
            is_admin    BOOLEAN NOT NULL DEFAULT FALSE
        );
        
        CREATE INDEX IF NOT EXISTS idx_email ON users(email);
        
        CREATE TABLE IF NOT EXISTS apps (
            id      SERIAL PRIMARY KEY,
            name    TEXT NOT NULL UNIQUE,
            secret  TEXT NOT NULL UNIQUE
        );
    `

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if _, err := db.ExecContext(ctx, query); err != nil {
		db.Close()
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) Close() {
	s.db.Close()
}

func (s *Storage) SaveUser(ctx context.Context, email string, passHash []byte) (int64, error) {
	const op = "storage.postgresql.SaveUser"

	stmt, err := s.db.Prepare("INSERT INTO users(email, pass_hash) VALUES($1, $2) RETURNING id")
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	defer stmt.Close()

	var userid int64
	err = stmt.QueryRowContext(ctx, email, passHash).Scan(&userid)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" { // email is not UNIQUE
			return 0, fmt.Errorf("%s: %w", op, storage.ErrUserExists)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	fmt.Println("SaveUser sand Query")

	return userid, nil
}

func (s *Storage) User(ctx context.Context, email string) (models.User, error) {
	return models.User{ID: 52, Email: "lol@lol.com", PassHash: []byte("##secret##")}, nil
}

func (s *Storage) IsAdmin(ctx context.Context, userID int64) (bool, error) {
	return true, nil
}

func (s *Storage) App(ctx context.Context, appID int) (models.App, error) {
	return models.App{ID: 35, Name: "BASE APP", Secret: "##secret##"}, nil
}
