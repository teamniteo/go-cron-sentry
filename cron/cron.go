package cron

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/imroc/req"
)

type Cron struct {
	Team    string
	DSN     string
	Monitor string
}

func (c *Cron) header() req.Header {
	return req.Header{
		"Content-Type":  "application/json",
		"Authorization": fmt.Sprintf("DSN %s", c.DSN),
	}
}

func (c *Cron) url(verb string) string {
	s, _ := url.JoinPath("https://sentry.io/api/0/organizations/", c.Team, "monitors", c.Monitor, verb)
	return s
}

func NewMonitor(team, monitor string) Cron {
	sentryDSN := os.Getenv("SENTRY_DSN")

	return Cron{
		Team:    team,
		DSN:     sentryDSN,
		Monitor: monitor,
	}
}

func (m *Cron) Start() error {
	url := m.url("/checkins/")
	res, err := req.Post(url, m.header(), started.json())
	log.Println(string(res.Bytes()))
	if err != nil {
		log.Println(err)
	}
	return err
}

func (m *Cron) Stop() error {
	url := m.url("/checkins/latest/")

	// handle crash
	if err := recover(); err != nil {

		res, errReq := req.Put(url, m.header(), errored.json())
		log.Println(string(res.Bytes()))
		if errReq != nil {
			log.Println(errReq)
		}
		return errReq
	}

	res, errReq := req.Put(url, m.header(), finished.json())
	log.Println(string(res.Bytes()))
	if errReq != nil {
		log.Println(errReq)
	}

	return errReq
}
