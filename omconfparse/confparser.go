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
			ParseWebapps(v.(map[interface{}]interface{}))
		}
	}
}

func ParseWebapps(webapp map[interface{}]interface{}) {
	for k, v := range webapp {
		if _, ok := k.(int); ok == true {
			ParseSite(v.(map[interface{}]interface{}))
		}
		/*switch v := k.(type) {
		case int:
			fmt.Println("k.(type):", "int", v)
		case uint16:
			fmt.Println("k.(type):", "uint16", v)
		case string:
			fmt.Println("k.(type):", "string", v)
		default:
			fmt.Println("k.(type):", "other", v)
		}*/
	}
}

func ParseSite(sitedata map[interface{}]interface{}) {
	for k, v := range sitedata {
		if k == "entry" {
			fmt.Println("entry:", v)
		}
	}
}
