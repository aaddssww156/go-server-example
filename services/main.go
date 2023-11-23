package services

import (
	"database/sql"
	"time"
)

var db *sql.DB

const dbTimeout = 3 * time.Second

type Models struct {
	Coffe        Coffe
	JsonResponse JsonResponse
}

func New(dbPool *sql.DB) Models {
	db = dbPool
	return Models{}
}
