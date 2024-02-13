package middlewares

import "net/http"

func Cors(next *http.ServeMux) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Authorization, DNT,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Range")
    w.Header().Set("Access-Control-Expose-Headers", "Authorization, Content-Length, Content-Range")

    if r.Method == "OPTIONS" {
        w.Header().Set("Access-Control-Max-Age", "1728000")
        w.Header().Set("Content-Type", "text/plain charset=UTF-8")
        w.Header().Set("Content-Length", "0")
        w.WriteHeader(http.StatusNoContent)
        return
    }
		next.ServeHTTP(w, r)
	})
}
//add_header 'Access-Control-Allow-Private-Network' 'true';
