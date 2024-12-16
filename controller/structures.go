package controller

type GaiaStatus struct {
	NodeInfo struct {
		ProtocolVersion struct {
			P2P   string `json:"p2p default:"0""`
			Block string `json:"block" default:"0"`
			App   string `json:"app" default:"0"`
		} `json:"protocol_version"`
		ID         string `json:"id" default:"0"`
		ListenAddr string `json:"listen_addr" default:"0"`
		Network    string `json:"network" default:"0"`
		Version    string `json:"version" default:"0"`
		Channels   string `json:"channels" default:"0"`
		Moniker    string `json:"moniker" default:"0"`
		Other      struct {
			TxIndex    string `json:"tx_index" default:"0"`
			RPCAddress string `json:"rpc_address" default:"0"`
		} `json:"other"`
	} `json:"node_info"`
	SyncInfo struct {
		LatestBlockHash   string `json:"latest_block_hash" default:"0"`
		LatestAppHash     string `json:"latest_app_hash" default:"0"`
		LatestBlockHeight string `json:"latest_block_height" default:"0"`
		LatestBlockTime   string `json:"latest_block_time" default:"0"`
		EarliestBlockHash string `json:"earliest_block_hash" default:"0"`
		EarliestAppHash   string `json:"earliest_app_hash" default:"0"`
		EarliestBlockHeight string `json:"earliest_block_height" default:"0"`
		EarliestBlockTime   string `json:"earliest_block_time" default:"0"`
		CatchingUp        bool   `json:"catching_up" default:"0"`
	} `json:"sync_info"`
	ValidatorInfo struct {
		Address string `json:"address"`
		PubKey  struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"pub_key"`
		VotingPower string `json:"voting_power"`
	} `json:"validator_info"`
}
