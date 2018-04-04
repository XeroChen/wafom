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
	Intellect_defence_enabled string      `yaml:",omitempty"`
	intellect_defence_data    interface{} `yaml:",omitempty"` //* 已废弃
	Statslog                  string
	Statslog_querystring      string
	Statslog_filter           string
	Statslog_filetypes        string
	Statslog_defaultfile      string
	Virtualpatch              string `yaml:",omitempty"` //是否开启虚拟补丁  off/...
	Virtualpatchfile          string `yaml:",omitempty"` //虚拟补丁规则文件路径 string
	Linkage_mode              string `yaml:",omitempty"` //部署模式 enum transparent/reverse/bypass/bridge
	Https                     string
	Ssl_protocols             interface{} `yaml:",omitempty"` //* 已废弃
	Https_public              []string    `yaml:",omitempty"` //HTTPS公钥证书信息 array
	Https_private             []string    `yaml:",omitempty"` //HTTPS似钥证书信息 array
	Https_chain               []string    `yaml:",omitempty"` //HTTPS证书链信息 array
	Ddos                      string
	Ddos_total                interface{} `yaml:",omitempty"` //*已废弃
	Ddos_second               interface{} `yaml:",omitempty"` //*已废弃
	Cache                     string
	Compression_server        string
	Compression_client        string
	Tamper                    string
	Tamper_mode               []string `yaml:",omitempty"` //防篡改运行模式 enum LEARNING_MODE/PROTECTING_MODE/DETECTING_MODE/TIMING_MODE
	Tamper_text_suffix        string   `yaml:",omitempty"` //防篡改文件后缀名列表 string
	Tamper_indexing           string   `yaml:",omitempty"` //防篡改默认页面 string
	Tamper_mime               []string `yaml:",omitempty"` //防篡改MIME类型 array
	Tamper_timing_ranges      []string `yaml:",omitempty"` //防篡改定时模式的时间范围 array
	Tamper_timing_other       string   `yaml:",omitempty"` //防篡改定时模式之外的运行模式
	Transparent               string   `yaml:",omitempty"`
	Xfwd4_status              string   `yaml:",omitempty"`
	Xfwd4                     string   `yaml:",omitempty"`
	Frontend_port             string   `yaml:",omitempty"`
	Load_balancing            string   `yaml:",omitempty"` //负载均衡工作模式 enum disabled/round_robin/ip_hash/least_conn
	Load_balancing_nodes      []string `yaml:",omitempty"` //负载均衡节点列表 array
	Linkage                   interface{}
	Mode                      string
	Vrrp                      string
	Vrrp_position             string   `yaml:",omitempty"`
	Vrrp_virtual_route_id     string   `yaml:",omitempty"`
	Domain_list               []string `yaml:",flow"`
	Acls_check_native         string   `yaml:",omitempty"`
	Acls_check_xfwd4          string   `yaml:",omitempty"`
	Blacklist                 string   `yaml:",omitempty"`
	Header_del                string
	Header_dels               []string
	Vlan                      string `yaml:",omitempty"`
	Vlan_id                   string `yaml:",omitempty"` //到保护站点的VLANID int
	Dvlan_id                  string `yaml:",omitempty"` //到保护站点的VLANID int
	Svlan_id                  string `yaml:",omitempty"` //到客户端的VLANID int
	Entry                     int
	Deploy                    string
	Enabled                   string
	Acls                      []string `yaml:",flow"`
	Log                       string
	Firewall                  []string    `yaml:",flow"`
	Mask                      string      `yaml:"mask,omitempty"`
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
	//present bool
	Ip      string
	Netmask string
	Gateway string
	Linkage string
}

/*
func (fn *FrontEndFmt) SetPresent(present bool) {
	fn.present = present
}

func (fn FrontEndFmt) IsZero() bool {
	return (bn.Ip == "" && bn.Netmask == "" && bn.Gateway == "" && bn.Linkage == "")
}*/

type BackendFmt struct {
	//present bool
	Ip      string
	Netmask string
	Gateway string
	Linkage string
}

/*
func (bn *BackendFmt) SetPresent(present bool) {
	bn.present = present
}

func (bn BackendFmt) IsZero() bool {
	return (bn.Ip == "" && bn.Netmask == "" && bn.Gateway == "" && bn.Linkage == "")
}*/
