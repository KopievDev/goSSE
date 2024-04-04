package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		// Устанавливаем заголовок Content-Type на text/event-stream
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		// Создаем таймер, который будет отправлять события каждые 10 секунд
		ticker := time.NewTicker(3 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-r.Context().Done():
				return
			case <-ticker.C:
				// Генерируем случайное число для выбора варианта
				randomNumber := rand.Intn(2)

				// Определяем, какой вариант отправить
				var jsonData []byte
				var event string
				if randomNumber == 0 {
					data := map[string]string{
						"hello": "world",
						"name":  "zalupa",
					}
					jsonData, _ = json.Marshal(data)
					event = "firstEvent"
				} else {
					data := map[string]string{
						"helloS": "zaebal",
						"name":   "zalupa",
					}
					jsonData, _ = json.Marshal(data)
					event = "secondEvent"
				}

				// Отправляем данные клиенту с указанием события
				fmt.Fprintf(w, "event: %s\ndata: %s\n\n", event, string(jsonData))
				fmt.Fprintln(os.Stdout, []any{"event: %s\ndata: %s\n\n", event, string(jsonData)}...)

				// Форсируем отправку данных клиенту
				w.(http.Flusher).Flush()
			}
		}
	})

	// Запускаем сервер на порту 8080
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}
