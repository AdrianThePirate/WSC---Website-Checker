package crawler

import (
	"log/slog"
)

func Start(url string) error {
	slog.Log("starting crawler", "url", url)
	return nil
}
