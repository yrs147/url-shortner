package routes

import (
	"time"

)

type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

type response struct {
	URL         	string         `json:"url"`
	CustomShort     string         `json:"short"`
	Expiry      	time.Duration  `json:"expiry"`
	RateLimit		int			   `json:"rate_limit"`
	RateLimitReset  time.Duration  `json:"rate_limit_reset"`
}


func main() {

}
