package middleware

import (
    "net/http"
    "project-layout-go/internal/infrastructure/common/config"
    "strconv"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/prometheus/client_golang/prometheus"
)

var labels = []string{"method", "uri", "code"}

var (
    requests = prometheus.NewCounterVec(prometheus.CounterOpts{
        Namespace: config.AppName,
        Subsystem: "http",
        Name:      "requests_count",
        Help:      "Requests count by method/path/status.",
    }, labels)

    durations = prometheus.NewSummaryVec(prometheus.SummaryOpts{
        Namespace:  config.AppName,
        Subsystem:  "http",
        Name:       "responses_duration_seconds",
        Help:       "Response time by method/path/status.",
        Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
    }, labels)
)

func init() {
    prometheus.MustRegister(requests, durations)
}

// HTTPMetrics is the middleware function that logs duration of responses.
func HTTPMetrics() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        c.Next()
        if isNeedRecordStatus(c.Writer.Status()) {
            metrics := []string{c.Request.Method, c.Request.URL.Path, strconv.Itoa(c.Writer.Status())}
            durations.WithLabelValues(metrics...).Observe(time.Since(start).Seconds())
            requests.WithLabelValues(metrics...).Inc()
        }
    }
}

func isNeedRecordStatus(status int) bool {
    return status == http.StatusOK || status == http.StatusInternalServerError || status == http.StatusGatewayTimeout ||
        status == http.StatusBadGateway || status == http.StatusServiceUnavailable
}
