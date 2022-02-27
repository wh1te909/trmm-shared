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
	GoArch       string  `json:"goarch"`
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
	Agentid string `json:"agent_id"`
	Disks   []Disk `json:"disks"`
}

type Disk struct {
	Device  string `json:"device"`
	Fstype  string `json:"fstype"`
	Total   string `json:"total"`
	Used    string `json:"used"`
	Free    string `json:"free"`
	Percent int    `json:"percent"`
}

type PublicIPNats struct {
	Agentid  string `json:"agent_id"`
	PublicIP string `json:"public_ip"`
}

type WinSoftwareList struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Publisher   string `json:"publisher"`
	InstallDate string `json:"install_date"`
	Size        string `json:"size"`
	Source      string `json:"source"`
	Location    string `json:"location"`
	Uninstall   string `json:"uninstall"`
}

type WinSoftwareNats struct {
	Agentid  string            `json:"agent_id"`
	Software []WinSoftwareList `json:"software"`
}
