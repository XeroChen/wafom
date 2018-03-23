package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Emb_struct1 struct {
	Sfield1 string
	Sfield2 string
}

type Emb_struct2 struct {
	present bool
	Ifield1 int
	Ifield2 int
}

/* 妈的这儿有个坑啊 定义成func (st *Emb_struct2) IsZero() bool是不行的 */
func (st Emb_struct2) IsZero() bool {
	return !st.present
}

type TestYaml struct {
	Reverse_xfwd4 string      `yaml:"reverse_xfwd4"`
	Embb1         Emb_struct1 `yaml:",omitempty"`
	Embb2         Emb_struct2 `yaml:",omitempty"`
	Other         string
}

func yamlMarshal(st *TestYaml) {
	out, err := yaml.Marshal(st)
	if err != nil {
		fmt.Printf("yaml_marshal failed with error:\n%s", err.Error())
		return
	}
	fmt.Printf("yaml_marshal succeeded with:\n%s", string(out))
}

func yamlUnmashal() {
	content, err := ioutil.ReadFile("./test.yaml")
	if err != nil {
		return
	}
	var conf TestYaml
	err = yaml.Unmarshal(content, &conf)
	if err != nil {
		fmt.Printf("yaml_unmarshal failed with error:\n%s", err.Error())
		return
	}
	fmt.Printf("yaml_unmarshal succeeded with:\n%v", conf)
	return
}

func main() {
	var conf TestYaml
	conf.Reverse_xfwd4 = "on"
	conf.Embb1.Sfield1 = "xxx"
	conf.Embb1.Sfield2 = "x-o"
	conf.Embb2.present = true
	conf.Embb2.Ifield1 = 0
	conf.Embb2.Ifield2 = 0
	conf.Other = "yes"
	yamlMarshal(&conf)
	yamlUnmashal()
}