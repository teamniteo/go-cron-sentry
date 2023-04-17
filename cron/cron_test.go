package cron

import "testing"

func TestCron_url(t *testing.T) {

	tests := []struct {
		name   string
		fields Cron
		verb   string
		want   string
	}{
		{"started", Cron{"team", "dsn", "monitor"}, "checkins", "https://sentry.io/api/0/organizations/team/monitors/monitor/checkins/"},
		{"ok", Cron{"team", "dsn", "monitor"}, "latest", "https://sentry.io/api/0/organizations/team/monitors/monitor/latest/"},
		{"errored", Cron{"team", "dsn", "monitor"}, "latest", "https://sentry.io/api/0/organizations/team/monitors/monitor/latest/"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cron{
				Team:    tt.fields.Team,
				DSN:     tt.fields.DSN,
				Monitor: tt.fields.Monitor,
			}
			if got := c.url(tt.verb); got != tt.want {
				t.Errorf("Cron.url() = %v, want %v", got, tt.want)
			}
		})
	}
}
