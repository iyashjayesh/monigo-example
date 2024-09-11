package main

import (
	"net/http"

	"github.com/iyashjayesh/monigo"
	"github.com/iyashjayesh/monigo/models"
	"golang.org/x/exp/rand"
)

func main() {

	monigoInstance := &monigo.Monigo{
		ServiceName:             "data-api", // Compulsory
		PurgeMonigoStorage:      true,       // Default is false
		DashboardPort:           8080,       // Default is 8080
		DataPointsSyncFrequency: "1m",       // Default is 5 Minutes
		DataRetentionPeriod:     "4d",       // Default is 14 days (2 weeks)
	}

	// **Thresholds Explanation:**

	// The `Thresholds` structure defines the performance and resource usage thresholds used to evaluate system health:

	// - **Low**: The percentage value below which the system is considered to be in optimal health.
	// - **Medium**: The percentage value indicating moderate health; usage above this threshold but below the High threshold may be acceptable but should be monitored.
	// - **High**: The percentage value indicating high usage, suggesting potential performance issues or resource constraints.
	// - **Critical**: The percentage value where the system is critically stressed and immediate attention is needed.
	// - **GoroutinesLow**: The lower bound for the number of goroutines; fewer goroutines are considered better.
	// - **GoroutinesHigh**: The upper bound for the number of goroutines; more goroutines may indicate high load or potential inefficiencies.

	// Example values:
	// - `Low: 20.0` - The system is healthy if usage is below 20%.
	// - `Medium: 50.0` - Usage between 20% and 50% is moderate.
	// - `High: 80.0` - Usage between 50% and 80% is high.
	// - `Critical: 100.0` - Usage at or above 100% is critical.

	// For goroutines:
	// - `GoroutinesLow: 100` - Ideal number of goroutines is below 100.
	// - `GoroutinesHigh: 500` - The system may experience performance issues if the number of goroutines exceeds 500.

	monigoInstance.ConfigureServiceThresholds(&models.ServiceHealthThresholds{
		Low:            20.0,  // Default is 20.0
		Medium:         50.0,  // Default is 50.0
		High:           80.0,  // Default is 80.0
		Critical:       100.0, // Default is 100.0
		GoRoutinesLow:  100,   // Default is 100
		GoRoutinesHigh: 500,   // Default is 500
	})

	monigoInstance.Start()

	// monigoInstance.DeleteMonigoStorage() // Delete monigo storage
	// monigoInstance.SetDbSyncFrequency("1m")
	// routinesStats := monigoInstance.PrintGoRoutinesStats() // Print go routines stats
	// log.Println(routinesStats)

	http.HandleFunc("/api", apiHandler)
	http.HandleFunc("/api2", apiHandler2)
	http.ListenAndServe(":8000", nil)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	// start := time.Now()

	// monigo.MeasureExecutionTime("MemExpensiveFunc", memexpensiveFunc)
	// monigo.MeasureExecutionTime("CpuExpensiveFunc", cpuexpensiveFunc)
	// monigo.RecordRequestDuration(time.Since(start))
	w.Write([]byte("API response"))
}

func apiHandler2(w http.ResponseWriter, r *http.Request) {
	// start := time.Now()
	// monigo.MeasureExecutionTime("MemExpensiveFunc", memexpensiveFunc)
	// monigo.MeasureExecutionTime("CpuExpensiveFunc", cpuexpensiveFunc)
	// monigo.RecordRequestDuration(time.Since(start))
	w.Write([]byte("API response"))
}

//go:noinline
func memexpensiveFunc() {
	m := make([]int, 10_000_000)
	for i := range m {
		m[i] = rand.Intn(127)
	}

	memanotherExpensiveFunc()
}

//go:noinline
func memanotherExpensiveFunc() {
	m := make(map[int]float64, 1_000_000)

	for key := range m {
		m[key] = rand.Float64()
	}
}

//go:noinline
func cpuexpensiveFunc() {
	var sum float64
	for i := 0; i < 10_000_000; i++ {
		sum += rand.Float64()
	}

	anotherExpensiveFunc()
}

//go:noinline
func anotherExpensiveFunc() {
	var sum int
	for i := 0; i < 1_000_000; i++ {
		sum += rand.Intn(10)
	}
}
