package statistic

import (
	"encoding/json"
	"net/http"
	"time"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.RequestURI != "/stat" {
			statistics = append(statistics, StatisticType{
				RemoteAddr: r.RemoteAddr,
				Date:       time.Now().Format("02.01.06 15:04:05"),
				Method:     r.Method,
				User:       r.Header.Get("User-Agent"),
				RequestURI: r.RequestURI,
			})
		} else {
			json.NewEncoder(w).Encode(statistics)
			return
		}
		//fmt.Printf("%+v\n", r)

		next.ServeHTTP(w, r)
	})
}
