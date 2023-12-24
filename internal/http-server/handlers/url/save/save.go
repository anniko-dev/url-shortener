package save

import (
	"log/slog"
	"net/http"
)

type URLSaver interface {
	SaveURL(urlToSave, alias string) (int64, error)
}

func New(log *slog.Logger, urlSaver URLSaver) http.Handler {

}
