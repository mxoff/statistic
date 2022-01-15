package statistic

type StatisticType struct {
	RemoteAddr string
	Date       string
	Method     string
	User       string
	RequestURI string
}

var statistics []StatisticType
