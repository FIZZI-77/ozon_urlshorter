package repository

import (
	"context"
	"database/sql"
	"ozon/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestUrlRepo_Create(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewUrlRepo(db)

	link := &models.Link{
		URL:   "https://google.com",
		Short: "abc1234567",
	}

	mock.ExpectExec("INSERT INTO urls").
		WithArgs(link.URL, link.Short).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Create(context.Background(), link)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestUrlRepo_GetByShort(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewUrlRepo(db)

	rows := sqlmock.NewRows([]string{"id", "base_url", "short_url"}).
		AddRow(1, "https://google.com", "abc1234567")

	mock.ExpectQuery("SELECT id, base_url, short_url FROM urls WHERE short_url").
		WithArgs("abc1234567").
		WillReturnRows(rows)

	link, err := repo.GetByShort(context.Background(), "abc1234567")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if link.URL != "https://google.com" {
		t.Fatalf("wrong url")
	}
}

func TestUrlRepo_GetByShort_NotFound(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewUrlRepo(db)

	mock.ExpectQuery("SELECT id, base_url, short_url FROM urls WHERE short_url").
		WithArgs("not_exist").
		WillReturnError(sql.ErrNoRows)

	_, err := repo.GetByShort(context.Background(), "not_exist")

	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestUrlRepo_GetByURL(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewUrlRepo(db)

	rows := sqlmock.NewRows([]string{"id", "base_url", "short_url"}).
		AddRow(1, "https://google.com", "abc1234567")

	mock.ExpectQuery("SELECT id, base_url, short_url FROM urls WHERE base_url").
		WithArgs("https://google.com").
		WillReturnRows(rows)

	link, err := repo.GetByURL(context.Background(), "https://google.com")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if link.Short != "abc1234567" {
		t.Fatalf("wrong short")
	}
}
