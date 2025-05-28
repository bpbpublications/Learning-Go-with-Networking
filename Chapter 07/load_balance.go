package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

func main() {
	backendURLs := []string{"http://server1:8080", "http://server2:8080", "http://server3:8080"}

	// Create a load balancer
	lb := NewLoadBalancer(backendURLs)

	// Start a web server using the load balancer
	http.HandleFunc("/", lb.ServeHTTP)
	fmt.Println("Load balancer running on :8080")
	http.ListenAndServe(":8080", nil)
}

// LoadBalancer represents a simple round-robin load balancer
type LoadBalancer struct {
	backends []*httputil.ReverseProxy
	current  int
	mu       sync.Mutex
}

func NewLoadBalancer(backendURLs []string) *LoadBalancer {
	var backends []*httputil.ReverseProxy
	for _, url := range backendURLs {
		target, _ := url.Parse(url)
		backends = append(backends, httputil.NewSingleHostReverseProxy(target))
	}
	return &LoadBalancer{backends: backends}
}

func (lb *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	// Select the next backend in a round-robin fashion
	backend := lb.backends[lb.current]
	lb.current = (lb.current + 1) % len(lb.backends)

	// Proxy the request to the selected backend
	backend.ServeHTTP(w, r)
}
