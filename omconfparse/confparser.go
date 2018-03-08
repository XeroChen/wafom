package omconfparse

import (
	//"github.com/go-yaml/yaml"
	"fmt"
	"io/ioutil"
	//"../omdata"
	"gopkg.in/yaml.v2"
	//"reflect"
)

// ParseYamlFile parse yaml file to map[interface{}]interface{}
func ParseYamlFile(filepath string) (result map[interface{}]interface{}, err error) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	conf := make(map[interface{}]interface{})
	err = yaml.Unmarshal(content, &conf)
	return conf, err
}

func ParseConf(confdata map[interface{}]interface{}) {
	for k, v := range confdata {
		if k == "webapps" {
			val := v.(map[interface{}]interface{})
			ParseSite(val)
		}
	}
}

func ParseWebapps(webapp map[interface{}]interface{}) {
	/*for k, v := range webapp {
	        if
		}*/
}

func ParseSite(sitedata map[interface{}]interface{}) {
	//var siteinfo wafsite.Site
	//sd, ok := sitedata.(map[interface{}]interface{})
	//fmt.Printf("\n%v\n--- t:\n%v\n\n", ok, sd)
	fmt.Printf("--- t:\n%v\n\n", sitedata)
}
