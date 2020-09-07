package main

import (
	"encoding/json"

	"github.com/davecgh/go-spew/spew"
	"github.com/petuhovskiy/hse-ruz-tg/pkg/conf"
	log "github.com/sirupsen/logrus"

	"github.com/petuhovskiy/telegram"
	"github.com/petuhovskiy/telegram/updates"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	log.SetReportCaller(true)
	log.SetLevel(log.DebugLevel)

	cfg, err := conf.ParseEnv()
	if err != nil {
		log.WithError(err).Fatal("in conf.ParseEnv()")
	}

	log.SetFormatter(&log.JSONFormatter{PrettyPrint: cfg.Bot.PrettyPrint})

	bot := telegram.NewBotWithOpts(cfg.Bot.Token, &telegram.Opts{
		Middleware: func(handler telegram.RequestHandler) telegram.RequestHandler {
			return func(methodName string, req interface{}) (message json.RawMessage, err error) {
				res, err := handler(methodName, req)
				if err != nil {
					log.WithError(err).Error("telegram response error")
				}

				return res, err
			}
		},
	})

	ch, err := updates.StartPolling(bot, telegram.GetUpdatesRequest{
		Offset:  0,
		Limit:   50,
		Timeout: 10,
	})
	if err != nil {
		log.WithError(err).Fatal("in updates.StartPolling()")
	}

	for upd := range ch {
		upd := upd
		spew.Dump(upd)
	}
}
