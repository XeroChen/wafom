package omconfparse

import (
	"../omdata"
	"bytes"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func ParseWebAppFile(filepath string) (result omdata.WebAppFmt, err error) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return omdata.WebAppFmt{}, err
	}
	//conf := make(map[interface{}]interface{})
	var conf omdata.WebAppFmt
	err = yaml.Unmarshal(content, &conf)
	return conf, err
}

func GenWebAppFile(newfile string, newdata omdata.WebAppFmt) bool {

	content, err := yaml.Marshal(newdata)

	if err != nil {
		fmt.Printf("\nyaml.Marshal error: %s", err.Error())
		return false
	}

	fmt.Printf("\nioutil.WriteFile writes %s with %v bytes", newfile, bytes.Count(content, nil)-1)

	err = ioutil.WriteFile(newfile, content, 0644)

	if err != nil {
		fmt.Printf("\nioutil.WriteFile error: %s", err.Error())
		return false
	}

	return true
}

func ParseWebAppData(olddata omdata.WebAppFmt) (newdata omdata.WebAppFmt) {
	newdata.Deploy = olddata.Deploy
	newdata.Reverse_transparent = olddata.Reverse_transparent
	newdata.Reverse_client_ip = olddata.Reverse_client_ip
	newdata.Reverse_xfwd4 = olddata.Reverse_xfwd4
	newdata.Reverse_xfwd4_status = olddata.Reverse_xfwd4_status
	newdata.Web_server = olddata.Web_server
	newdata.Webapps = ParseWebapps(olddata.Webapps)
	return newdata
}

func ParseWebapps(olddata map[int]omdata.Webapps) (newdata map[int]omdata.Webapps) {

	newdata = make(map[int]omdata.Webapps)

	for k := range olddata {
		newdata[k] = ParseSite(olddata[k])
		fmt.Printf("\nSite %v parsed", k)
	}
	return newdata
}

func ParseSite(olddata omdata.Webapps) (newdata omdata.Webapps) {

	newdata = olddata
	/* Handle FrontEnd */
	if olddata.Frontend_linkage != "" || olddata.Frontend_ip != "" ||
		olddata.Frontend_netmask != "" || olddata.Frontend_gateway != "" {

		newdata.Frontend.Ip = olddata.Frontend_ip
		newdata.Frontend.Netmask = olddata.Frontend_netmask
		newdata.Frontend.Gateway = olddata.Frontend_gateway
		newdata.Frontend.Linkage = olddata.Frontend_linkage

		newdata.Frontend_ip = ""
		newdata.Frontend_netmask = ""
		newdata.Frontend_gateway = ""
		newdata.Frontend_linkage = ""

	}

	/* Handle BackEnd */
	if olddata.Backend_linkage != "" || olddata.Backend_ip != "" ||
		olddata.Backend_netmask != "" || olddata.Backend_gateway != "" {

		newdata.Backend.Ip = olddata.Backend_ip
		newdata.Backend.Netmask = olddata.Backend_netmask
		newdata.Backend.Gateway = olddata.Backend_gateway
		newdata.Backend.Linkage = olddata.Backend_linkage

		newdata.Backend_ip = ""
		newdata.Backend_netmask = ""
		newdata.Backend_gateway = ""
		newdata.Backend_linkage = ""

	}

	/* Handle Linkage */
	if s, ok := olddata.Linkage.(string); ok == true {
		if s == "FAKE" {
			newdata.Linkage = s
		} else {
			LinkageList := make([]string, 1)
			LinkageList[0] = s
			newdata.Linkage = LinkageList
		}
	} else {
		newdata.Linkage = olddata.Linkage
	}

	return newdata
}
