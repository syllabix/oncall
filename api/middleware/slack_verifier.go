package middleware

import (
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/slack-go/slack"
	"github.com/syllabix/oncall/config"
	"go.uber.org/zap"
)

// SlackVerifier is used to ensure requests are from Slack
type SlackVerifier struct {
	secret string
	log    *zap.Logger
}

func NewSlackVerifier(
	settings config.SlackSettings,
	log *zap.Logger,
) *SlackVerifier {
	return &SlackVerifier{
		secret: settings.SigningSecret,
		log:    log,
	}
}

// Verify is a method that serves as a middle for requests handlers from Slack.
// Slack signs all requests with the application issued signing key, and this middleware
// verifies the request is legit, sending back an http error and returning before handling
// control to the next middleware
func (s *SlackVerifier) Verify(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		verifier, err := slack.NewSecretsVerifier(r.Header, s.secret)
		if err != nil {
			s.log.Error("slack secrets verifier failed to initialize", zap.Error(err))
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		body := io.TeeReader(r.Body, &verifier)

		err = verifier.Ensure()
		if err != nil {
			s.log.Error("slack verifier encountered an unverifiable payload", zap.Error(err))
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		r.Body = io.NopCloser(body)
		next(w, r, ps)
	}
}
