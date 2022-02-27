package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"go.uber.org/multierr"
	"golang.org/x/sync/errgroup"

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
	serv := service.New(repo, conf.Storage)

	log.Info("initialize wait err group")
	wg := errgroup.Group{}

	log.Info("initialize http API")
	wg.Go(func() error {
		return http.New(serv).Listen(":" + conf.HTTP.Port)
	})

	return wg.Wait()
}
