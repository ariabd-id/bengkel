package customer

import (
	"bengkel/domain"
	"context"
	"database/sql"
	"github.com/doug-martin/goqu/v9"
)

type repository struct {
	db *goqu.Database
}

func NewRepository(con *sql.DB) domain.CustomerRepository {
	return &repository{
		db: goqu.New("default", con),
	}
}

func (r repository) FindById(ctx context.Context, id int64) (customer domain.Customer, err error) {
	dataset := r.db.From("customers").Where(goqu.Ex{"id": id})

	if _, err := dataset.ScanStructContext(ctx, &customer); err != nil {
		return domain.Customer{}, err
	}

	return
}

func (r repository) FindByIds(ctx context.Context, ids []int64) (customer []domain.Customer, err error) {
	dataset := r.db.From("customers").Where(goqu.Ex{"id": ids})

	if err := dataset.ScanStructsContext(ctx, &customer); err != nil {
		return nil, err
	}
	return nil, nil
}

func (r repository) FindByPhone(ctx context.Context, phone string) (customer domain.Customer, err error) {
	dataset := r.db.From("customers").Where(goqu.Ex{"phone": phone})

	if _, err := dataset.ScanStructContext(ctx, &customer); err != nil {
		return domain.Customer{}, err
	}

	return
}

func (r repository) Insert(ctx context.Context, customer *domain.Customer) error {
	executor := r.db.Insert("customers").Rows(*customer).Executor()
	result, err := executor.Exec()

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	customer.ID = id

	return err
}
