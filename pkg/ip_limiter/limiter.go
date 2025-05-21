package ipLimiter

import (
	"sync"

	"golang.org/x/time/rate"
)

type IPRateLimiter struct {
	ips     map[string]*rate.Limiter
	mu      *sync.RWMutex
	rps     rate.Limit
	burst   int
}
func NewIPRateLimiter(rps rate.Limit, burst int) *IPRateLimiter {
	return &IPRateLimiter{
		ips:     make(map[string]*rate.Limiter),
		mu:      &sync.RWMutex{},
		rps:     rps,
		burst:   burst,
	}
}
func (l *IPRateLimiter) AddIP(ip string) *rate.Limiter {
	l.mu.Lock()
	defer l.mu.Unlock()
	limiter := rate.NewLimiter(l.rps, l.burst)
	l.ips[ip] = limiter

	return limiter
}


func (l *IPRateLimiter) GetLimiter(ip string) *rate.Limiter {
	l.mu.Lock()
	limiter,exists := l.ips[ip]

	if !exists {
		l.mu.Unlock()
		return nil
	}
	l.mu.Unlock()
	return limiter
}


