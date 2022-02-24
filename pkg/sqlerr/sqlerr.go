package sqlerr

import "github.com/pkg/errors"


func WithSql(err error, query string, args ...interface{}) error {
	if err == nil {
		return nil
	}

	return errors.Wrapf(
		err,
		"query: %s args: %+v",
		query,
		args,
	)
}
