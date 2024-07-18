package domain

import (
	"context"
	"time"
)

type History struct {
	ID       int64     `db:"id"`
	NoRangka string    `db:"no_rangka"`
	Merek    string    `db:"merek"`
	UpdateAt time.Time `db:"update_at"`
}

type HistoryDetail struct {
	ID         int64     `db:"id"`
	HistoryID  int64     `db:"history_id"`
	PIC        string    `db:"pic"`
	CustomerID int64     `db:"customer_id"`
	PlatNomor  string    `db:"plat_nomor"`
	Notes      string    `db:"notes"`
	CreateAt   time.Time `db:"create_at"`
}

type HistoryRepository interface {
	FindById(ctx context.Context, id int64) (History, error)
	FindByNoRangka(ctx context.Context, no string) (History, error)
	FindDetailByHistory(ctx context.Context, id int64) ([]HistoryDetail, error)
	Insert(ctx context.Context, history *History) (History, error)
	InsertDetail(ctx context.Context, historyDetail *HistoryDetail) (HistoryDetail, error)
}

type HistoryService interface{}
