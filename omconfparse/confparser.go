package omconfparse

import (
	"github.com/mattn/go-sqlite3"
	//"github.com/go-yaml/yaml"
	"../omdb"
	"fmt"
	"github.com/go-xorm/xorm"
	"io/ioutil"
	//"../omdata"
	"gopkg.in/yaml.v2"
	//"reflect"
)

type ConfEngine struct {
	gDBEng     *xorm.Engine
	gConfFiles []string
}

func (eng *ConfEngine) Init() int {
	if dbeng, err := omdbsqlite.CreateDBEngine(); err == nil {
		eng.gDBEng = dbeng
		eng.init_tables()
		return 0
	}
	fmt.Println("[ERR] Init() failure!")
	return -1
}

func (eng *ConfEngine) init_tables() {
	return
}

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
			ParseSite(k.(int), v.(map[interface{}]interface{}))
		}
	}
}

func ParseSite(siteindex int, sitedata map[interface{}]interface{}) {
	for k, v := range sitedata {
		if k == "entry" {
			fmt.Println("entry:", v)
		}
	}
}
