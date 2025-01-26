package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// RateLimiter struct to manage token bucket
type RateLimiter struct {
	mu          sync.Mutex
	tokens      int
	maxTokens   int
	tokenRate   time.Duration
	lastUpdated time.Time
}

// NewRateLimiter creates a new RateLimiter
func NewRateLimiter(maxTokens int, tokenRate time.Duration) *RateLimiter {
	return &RateLimiter{
		tokens:      maxTokens,
		maxTokens:   maxTokens,
		tokenRate:   tokenRate,
		lastUpdated: time.Now(),
	}
}

// Allow checks if a request is allowed
func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	// Add tokens based on elapsed time
	now := time.Now()
	elapsed := now.Sub(rl.lastUpdated)
	newTokens := int(elapsed / rl.tokenRate)
	if newTokens > 0 {
		rl.tokens = min(rl.tokens+newTokens, rl.maxTokens)
		rl.lastUpdated = now
	}

	// Check if there are tokens available
	if rl.tokens > 0 {
		rl.tokens--
		return true
	}
	return false
}

// min utility function
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Middleware to enforce rate limiting
func rateLimitMiddleware(rl *RateLimiter, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if rl.Allow() {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
		}
	})
}

func main() {
	// Create a rate limiter: max 5 requests per second
	rl := NewRateLimiter(5, time.Second/5)

	// Example handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Request allowed!")
	})

	// Wrap the handler with rate limiting middleware
	http.Handle("/", rateLimitMiddleware(rl, handler))

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
