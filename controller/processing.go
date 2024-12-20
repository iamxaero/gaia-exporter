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
	gaia_number_of_peers_grouped_by_their_version = prometheus.NewCounterVec(
		prometheus.CounterOpts{
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

func parseFloatOrDefault(value string) float64 {
	var defaultValue float64 = 0
	if value == "" || value == "null" {
		return defaultValue
	}
	if f, err := strconv.ParseFloat(value, 64); err == nil {
		return f
	}
	return defaultValue
}
