package middlewares

import (
	"hasherapi/app/event"
	"net/http"
)

func RegisterCallEvents(next http.Handler, eventStream event.Stream) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/send":
			eventStream.Send(event.NewSendCallEvent())
		case "/check":
			eventStream.Send(event.NewCheckCallEvent())
		}

		next.ServeHTTP(w, r)
	})
}
