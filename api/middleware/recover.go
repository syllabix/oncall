package middleware

import (
	"log"
	"net/http"
	"runtime"
)

// Recover is simple middle ware for recovering from a panic that may
// occur downstream in the application
func Recover(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				buf := make([]byte, 1<<16)
				stackSize := runtime.Stack(buf, true)
				log.Printf("[PANIC] \n%v\n%s\n",
					err, string(buf[0:stackSize]))
				http.Error(w,
					"An unexpected error occurred", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	}
}
