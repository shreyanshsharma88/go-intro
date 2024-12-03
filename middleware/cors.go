package middleware
import (
	"net/http"
)
func CORSMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		res.Header().Set("Access-Control-Allow-Headers", "Content-Type, token")
		next(res, req)
	}
} 