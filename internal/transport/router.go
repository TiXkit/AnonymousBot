package transport

import (
	"AnonymusBot/internal/shared/config"
	"log"
	"log/slog"
	"net/http"
)

func RunRouter(certFile, keyFile string) {

	addr := config.GetHTTPAddress().Addr

	if err := http.ListenAndServeTLS(addr, certFile, keyFile, nil); err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}

	slog.Info("Сервер успешно запущен на порту %s.", addr)
}
