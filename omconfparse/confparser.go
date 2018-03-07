package omconfparse

import (
	//"github.com/go-yaml/yaml"
	"io/ioutil"

	"gopkg.in/yaml.v2"
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
			ParseSite(v)
		}
	}
}

func ParseSite(sitedata map[interface{}]interface{}) {

}
