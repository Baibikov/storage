package dbutils

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.uber.org/multierr"
)

type TxFunc func(tx *sqlx.Tx) error

func WrapTx(db *sqlx.DB, txFunc TxFunc) (err error) {
	begin, err := db.Beginx()
	if err != nil {
		return errors.Wrap(err, "beginning transaction")
	}
	defer func(err error) {
		if err != nil {
			multierr.AppendInto(&err, begin.Rollback())
		}
	}(err)

	err = txFunc(begin)
	if err != nil {
		return errors.Wrap(err, "transaction func exec")
	}

	return begin.Commit()
}
