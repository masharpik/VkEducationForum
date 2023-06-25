package middleware

import (
	"net/http"
)

func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// start := time.Now()
		// defer func() {
		// 	fmt.Printf("Function %s %s execution took %s\n", r.Method, r.URL.Path, time.Since(start))
		// }()
		// defer func() {
		// 	if err := recover(); err != nil {
		// 		writer.WriteErrorMessageRespond(w, r, http.StatusInternalServerError, fmt.Sprintf(mainLiterals.LogPanicOccured, err))
		// 	}
		// }()

		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
