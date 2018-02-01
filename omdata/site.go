package wafsite

type site struct {
	tablename string /* tbl_waf_site */
	index     uint32
	sitename  string
}

type serverpool struct {
	tablename string /* tbl_waf_serverpool */
	index     uint32
	ip        uint64
	port      uint32
}

type site_serverpool struct {
	tablename string /* tbl_waf_site_serverpool */
	siteindex uint32
	spindex   uint32
}

type site_vhost struct {
	tablename string /* tbl_waf_site_vhost */
	siteindex uint32
	vhindex   uint32
}

type vhost struct {
	tablename string /* tbl_waf_vhost */
	index     uint32
}

type vhmatchtable struct {
	tablename string /* tbl_waf_vhmatchtable */
	vhindex   uint32 /* key1 */
	ip1       uint64
	ip2       uint64
	port      uint16
	domain    string
}

type vhost_gateway struct {
	tablename string /* tbl_waf_vhost_gateway */
	vhindex   uint32
	gwindex   uint32
}

type gateway struct {
	tablename string /* tbl_waf_gateway */
	index     uint32
	gwtype    uint32
	gwip      uint64
	gwbr      string
}
