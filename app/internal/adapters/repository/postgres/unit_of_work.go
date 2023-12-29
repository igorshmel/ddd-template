package postgres

import (
	"fmt"
	"runtime"

	"default_ddd/app/internal/adapters/port"
)

var _ port.Persister = (*SQLStore)(nil)

// UnitOfWork --
func (ths *SQLStore) UnitOfWork(fn func(persister port.Persister) error) (err error) {
	trx := ths.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = trx.Rollback()

			switch e := p.(type) {
			case runtime.Error:
				panic(e)
			case error:
				err = fmt.Errorf("panic err: %v", p)
				return
			default:
				panic(e)
			}
		}
		if err != nil {
			trx.Rollback()
		} else {
			trx.Commit()
		}
	}()

	newStore := &SQLStore{
		db: trx,
	}
	return fn(newStore)
}
