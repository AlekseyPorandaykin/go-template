package metrics

import (
	"net/http"
	"time"
)

func IncCountQueryHttp(name string) {

}

func DurationExecuteQueryHttp(name string, duration time.Duration) {

}

func IncErrorQueryHttp(name string) {

}

type Sender interface {
	Do(req *http.Request) (*http.Response, error)
}

func DoHttpWithMetric(s Sender, name string) func(req *http.Request) (*http.Response, error) {
	return func(req *http.Request) (*http.Response, error) {
		IncCountQueryHttp(name)
		defer func(start time.Time) {
			DurationExecuteQueryHttp(name, time.Since(start))
		}(time.Now())
		resp, err := s.Do(req)
		if err != nil {
			IncErrorQueryHttp(name)
		}
		return resp, err
	}
}
