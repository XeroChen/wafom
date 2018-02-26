package wafsite

type site struct {
	index    uint32
	sitename string /* name */
	vhindex  uint32
	spindex  uint32
}

type serverpool struct {
	index uint32
	ip    uint64 /* ip */
	port  uint32 /* port */
}

type vhost struct {
	index   uint32
	gwindex uint32
}

type vhmatchtable struct {
	vhindex uint32 /* key1 */
	port    uint16 /* frontend_port */
	domain  string /* domain_list[] */
}

type gateway struct {
	index  uint32
	gwtype uint32   /* 0:ip 1:bridge REF:deploy:transparent */
	gwip   uint64   /* frontend */
	gwbr   []string /* linkage */
}
