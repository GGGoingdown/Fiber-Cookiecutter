package pkg

import (
	"log"

	"github.com/getsentry/sentry-go"
)

func NewSentryOptions(
	dsn string,
	traceSampleRate float64,
	environment string,
	serverName string,
	release string,
) *sentry.ClientOptions {
	return &sentry.ClientOptions{
		Dsn:              dsn,
		TracesSampleRate: traceSampleRate,
		Environment:      environment,
		ServerName:       serverName,
		Release:          release,
	}
}

func InitSentry(options *sentry.ClientOptions) {
	err := sentry.Init(*options)
	if err != nil {
		log.Fatalln("sentry initialization failed: ", err.Error())
	}
}
