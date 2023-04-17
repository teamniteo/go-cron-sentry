package cron

import (
	"fmt"
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
	return fmt.Sprintf("https://sentry.io/api/0/organizations/%s/monitors/%s/%s/", c.Team, c.Monitor, verb)
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
	_, err := req.Post(m.url("checkins"), m.header(), started.json())
	return err
}

func (m *Cron) Stop() error {

	// handle crash
	if err := recover(); err != nil {
		req.Put(m.url("latest"), m.header(), errored.json())
		return nil
	}

	_, err := req.Put(m.url("latest"), m.header(), finished.json())
	return err
}
