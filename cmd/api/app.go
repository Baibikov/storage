package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"go.uber.org/multierr"

	"storage/internal/app/delivery/http"
	"storage/internal/app/repository"
	"storage/internal/app/service"
)

func app() (err error) {
	log.Info("initialize database connection")
	db, err := sqlx.Open("postgres", conf.DB.Conn)
	if err != nil {
		return errors.Wrapf(
			err,
			"connect to postgresql database %s",
			conf.DB.Conn,
		)
	}
	defer func(err error) {
		multierr.AppendInto(&err, db.Close())
	}(err)

	log.Info("initialize repository")
	repo := repository.New(db)

	log.Info("initialize service")
	serv := service.New(repo)

	log.Info("initialize http API")
	err = http.New(serv).Listen(":" + conf.HTTP.Port)
	if err != nil {
		return errors.Wrap(err, "lister server")
	}

	return err
}
