package controller

import (
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	gaia_highest_block_number = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "gaia_highest_block_number",
			Help: "Highest block number the node has reached",
		},
	)
	gaia_current_block_time_drift_in_seconds = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "gaia_current_block_time_drift_in_seconds",
			Help: "Current block time drift in seconds: Current time minus block creation time",
		},
	)
	gaia_number_of_connected_peers = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "gaia_number_of_connected_peers",
			Help: "Number of peers grouped by their version",
		},
	)
	gaia_number_of_peers_grouped_by_their_version = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "gaia_number_of_peers_grouped_by_their_version",
			Help: "Pipeline Number of jobs",
		},
		[]string{
			"gaia_version",
		},
	)
)

func (c *Controller) PromRegister() {
	// register metrics
	prometheus.MustRegister(gaia_highest_block_number)
	prometheus.MustRegister(gaia_current_block_time_drift_in_seconds)
	prometheus.MustRegister(gaia_number_of_connected_peers)
	prometheus.MustRegister(gaia_number_of_peers_grouped_by_their_version)
}

func (c *Controller) ProcGaiaStatus(status GaiaStatus) {
	gaia_highest_block_number.Set(parseFloatOrDefault(status.Result.SyncInfo.LatestBlockHeight))
	gaia_current_block_time_drift_in_seconds.Set(float64(time.Now().Unix() - status.Result.SyncInfo.LatestBlockTime.Unix()))
	// It can be wrong, check after the last metric
	gaia_number_of_connected_peers.Set(parseFloatOrDefault(status.Result.NodeInfo.ProtocolVersion.P2P))
	// gaia_number_of_peers_grouped_by_their_version.WithLabelValues(status.Result.NodeInfo.Version).Add(1)
}

func (c *Controller) ProcGaiaNet(versions map[string]int) {
	for ver, count := range versions {
		gaia_number_of_peers_grouped_by_their_version.WithLabelValues(ver).Set(parseFloatOrDefault(count))
	}
}

func parseFloatOrDefault(value interface{}) float64 {
	var defaultValue float64 = 0
	switch value := value.(type) {
	case string:
		f, _ := strconv.ParseFloat(value, 64)
		return f
	case int:
		return (float64(value))
	default:
		return defaultValue
	}
}

// Find all versions in GAIA Net Info
func (c *Controller) FindVersions(data interface{}, versions map[string]int) {
	// Define type  of data
	switch value := data.(type) {
	// If dict
	case map[string]interface{}:
		for key, val := range value {
			if key == "version" {
				if versionStr, ok := val.(string); ok {
					versions[versionStr]++
				}
			} else {
				// Рекурсивно обрабатываем вложенные объекты
				c.FindVersions(val, versions)
			}
		}
	// if list
	case []interface{}:
		// Если текущий объект - это массив, обрабатываем каждый элемент
		for _, item := range value {
			c.FindVersions(item, versions)
		}
	}
}
