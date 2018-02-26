package omconfparse

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
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
