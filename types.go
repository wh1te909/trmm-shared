package trmm

type CheckInNats struct {
	Agentid string `json:"agent_id"`
	Version string `json:"version"`
}

type AgentInfoNats struct {
	Agentid      string  `json:"agent_id"`
	Username     string  `json:"logged_in_username"`
	Hostname     string  `json:"hostname"`
	OS           string  `json:"operating_system"`
	Platform     string  `json:"plat"`
	TotalRAM     float64 `json:"total_ram"`
	BootTime     int64   `json:"boot_time"`
	RebootNeeded bool    `json:"needs_reboot"`
}

type WinSvcNats struct {
	Agentid string           `json:"agent_id"`
	WinSvcs []WindowsService `json:"services"`
}

type WindowsService struct {
	Name             string `json:"name"`
	Status           string `json:"status"`
	DisplayName      string `json:"display_name"`
	BinPath          string `json:"binpath"`
	Description      string `json:"description"`
	Username         string `json:"username"`
	PID              uint32 `json:"pid"`
	StartType        string `json:"start_type"`
	DelayedAutoStart bool   `json:"autodelay"`
}

type WinWMINats struct {
	Agentid string      `json:"agent_id"`
	WMI     interface{} `json:"wmi"`
}

type WinDisksNats struct {
	Agentid string      `json:"agent_id"`
	Disks   interface{} `json:"disks"`
}
