package omdata

type WebAppFmt struct {
	Deploy               string
	Reverse_transparent  string
	Reverse_client_ip    string
	Reverse_xfwd4        string
	Reverse_xfwd4_status string
	Web_server           string
	Webapps              map[int]Webapps
}

type Webapps struct {
	Name                      string
	Ip                        string
	Port                      string
	Rule                      string
	Keepalive                 string
	Intellect_defence_enabled string `yaml:",omitempty"`
	Statslog                  string
	Statslog_querystring      string
	Statslog_filter           string
	Statslog_filetypes        string
	Statslog_defaultfile      string
	Https                     string
	Ddos                      string
	Cache                     string
	Compression_server        string
	Compression_client        string
	Tamper                    string
	Transparent               string `yaml:",omitempty"`
	Xfwd4_status              string `yaml:",omitempty"`
	Xfwd4                     string `yaml:",omitempty"`
	Frontend_port             int    `yaml:",omitempty"`
	Linkage                   interface{}
	Mode                      string
	Vrrp                      string
	Vrrp_position             string   `yaml:",omitempty"`
	Vrrp_virtual_route_id     string   `yaml:",omitempty"`
	Domain_list               []string `yaml:"flow"`
	Acls_check_native         string
	Acls_check_xfwd4          string
	Blacklist                 string
	Header_del                string
	Header_dels               []string
	Vlan                      string `yaml:",omitempty"`
	Entry                     int
	Deploy                    string
	Enabled                   string
	Acls                      []string `yaml:"flow"`
	Log                       string
	Firewall                  []string `yaml:"flow"`
	Mask                      string
	Frontend                  FrontEndFmt `yaml:"frontend,omitempty"`
	Backend                   BackendFmt  `yaml:"backend,omitempty"`
	//
	Action_type      string `yaml:",omitempty"`
	Frontend_linkage string `yaml:",omitempty"`
	Backend_linkage  string `yaml:",omitempty"`
	Frontend_ip      string `yaml:",omitempty"`
	Backend_ip       string `yaml:",omitempty"`
	Frontend_netmask string `yaml:",omitempty"`
	Backend_netmask  string `yaml:",omitempty"`
	Frontend_gateway string `yaml:",omitempty"`
	Backend_gateway  string `yaml:",omitempty"`
}

type FrontEndFmt struct {
	present bool
	Ip      string
	Netmask string
	Gateway string
	Linkage string
}

func (fn *FrontEndFmt) IsZero() bool {
	return !fn.present
}

type BackendFmt struct {
	present bool
	Ip      string
	Netmask string
	Gateway string
	Linkage string
}

func (bn *BackendFmt) IsZero() bool {
	return !bn.present
}
