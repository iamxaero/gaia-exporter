package controller

import (
	"strconv"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	gaia_highest_block_number = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "gaia_highest_block_number",
			Help: "Highest block number the node has reached",
		},
		[]string{
			"gaia_version",
		},
	)
	gaia_current_block_time_drift_in_seconds = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "gaia_current_block_time_drift_in_seconds",
			Help: "Current block time drift in seconds: Current time minus block creation time",
		},
		[]string{
			"gaia_version",
		},
	)
	gaia_number_of_connected_peers = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "gaia_number_of_connected_peers",
			Help: "Number of peers grouped by their version",
		},
		[]string{
			"gaia_version",
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

func (c *Controller) ProcGaiaStatus (status GaiaStatus) {
	gaia_highest_block_number.WithLabelValues(status.NodeInfo.Version).Set(parseFloatOrDefault(status.SyncInfo.LatestBlockHeight))
	gaia_current_block_time_drift_in_seconds.WithLabelValues(status.NodeInfo.Version).Set(parseFloatOrDefault(status.SyncInfo.LatestBlockTime))
	gaia_number_of_connected_peers.WithLabelValues(status.NodeInfo.Version).Set(parseFloatOrDefault(status.NodeInfo.ProtocolVersion.Block))
	gaia_number_of_peers_grouped_by_their_version.WithLabelValues(status.NodeInfo.Version).Add(1)
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
