package omconfparse

import (
	"../omdata"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// ParseYamlFile parse yaml file to map[interface{}]interface{}
//func ParseWebAppFile(filepath string) (result map[interface{}]interface{}, err error) {
func ParseWebAppFile(filepath string) (result omdata.WebAppFmt, err error) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return omdata.WebAppFmt{}, err
	}
	//conf := make(map[interface{}]interface{})
	var conf omdata.WebAppFmt
	err = yaml.Unmarshal(content, &conf)
	/*
		for k, v := range conf.Webapps {
			fmt.Printf("\nSite %v: \n%v\n", k, v)
		}
	*/
	return conf, err
}

func ParseConf(confdata map[interface{}]interface{}) {
	for k, v := range confdata {
		if k == "webapps" {
			ParseWebapps(v.(map[interface{}]interface{}))
		}
	}
}

func ParseWebapps(webapp map[interface{}]interface{}) {
	for k, v := range webapp {
		if _, ok := k.(int); ok == true {
			/*test for site 1*/
			if k.(int) == 1 {
				ParseSite(k.(int), v.(map[interface{}]interface{}))
			}
		}
	}
}

func ParseSite(siteindex int, sitedata map[interface{}]interface{}) {
	for k, v := range sitedata {
		if k == "entry" {
			fmt.Println("entry:", v)
		}
		if k == "header_dels" {
			fmt.Println("header_dels:", v)
		}
		if k == "firewall" {
			fmt.Println("firewall:", v)
		}
	}
}
