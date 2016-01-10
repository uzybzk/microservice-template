package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
    "time"
)

type Config struct {
    Port        string `json:"port"`
    ServiceName string `json:"service_name"`
    Version     string `json:"version"`
}

type HealthResponse struct {
    Status    string    `json:"status"`
    Service   string    `json:"service"`
    Version   string    `json:"version"`
    Timestamp time.Time `json:"timestamp"`
    Uptime    string    `json:"uptime"`
}

var (
    config    Config
    startTime time.Time
)

func init() {
    startTime = time.Now()
    
    config = Config{
        Port:        getEnv("PORT", "8080"),
        ServiceName: getEnv("SERVICE_NAME", "microservice-template"),
        Version:     getEnv("VERSION", "1.0.0"),
    }
}

func main() {
    setupRoutes()
    
    fmt.Printf("Starting %s v%s on port %s\n", 
        config.ServiceName, config.Version, config.Port)
    
    log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}

func setupRoutes() {
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/health", healthHandler)
    http.HandleFunc("/ready", readyHandler)
    http.HandleFunc("/metrics", metricsHandler)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    
    response := map[string]interface{}{
        "service": config.ServiceName,
        "version": config.Version,
        "message": "Microservice Template",
        "endpoints": []string{
            "/health - Health check",
            "/ready - Readiness check", 
            "/metrics - Service metrics",
        },
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
    uptime := time.Since(startTime).String()
    
    health := HealthResponse{
        Status:    "healthy",
        Service:   config.ServiceName,
        Version:   config.Version,
        Timestamp: time.Now(),
        Uptime:    uptime,
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(health)
}

func readyHandler(w http.ResponseWriter, r *http.Request) {
    // In a real microservice, this would check dependencies
    // like database connections, external services, etc.
    
    ready := map[string]interface{}{
        "status": "ready",
        "checks": map[string]string{
            "database": "ok",
            "cache":    "ok",
        },
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(ready)
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
    // Simple Prometheus-style metrics
    metrics := fmt.Sprintf(`# HELP service_uptime_seconds Service uptime in seconds
# TYPE service_uptime_seconds counter
service_uptime_seconds %.0f

# HELP service_requests_total Total number of requests
# TYPE service_requests_total counter
service_requests_total{method="GET",endpoint="/"} 42

# HELP service_info Service information
# TYPE service_info gauge
service_info{service="%s",version="%s"} 1
`, 
        time.Since(startTime).Seconds(),
        config.ServiceName,
        config.Version)
    
    w.Header().Set("Content-Type", "text/plain")
    fmt.Fprint(w, metrics)
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}