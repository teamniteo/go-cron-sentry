package cron

import "strings"

type Report string

const (
	started  Report = `{"status": "in_progress"}`
	finished Report = `{"status": "ok"}`
	errored  Report = `{"status": "error"}`
)

func (r Report) json() *strings.Reader {
	return strings.NewReader(string(r))
}
