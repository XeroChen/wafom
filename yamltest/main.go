package main

import (
	//"../omconfparse"
	//"../omdata"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type ItfUpgrade interface {
	Upgrade() interface{}
}

type Emb_struct1 struct {
	Sfield1 string
	Sfield2 string
}

func (old Emb_struct1) Upgrade() (new Emb_struct1) {
	new = old
	fmt.Println("Emb_struct1 upgrading.")
	return
}

type Emb_struct2 struct {
	present bool
	Ifield1 int
	Ifield2 int
}

func (old Emb_struct2) Upgrade() (new Emb_struct2) {
	new = old
	fmt.Println("Emb_struct2 upgrading.")
	return
}

type Emb_struct3 struct {
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
	Embb3         Emb_struct3 `yaml:",omitempty"`
	Other         interface{} `yaml:",flow"`
	Other2        string
}

func DoUpgrade() {
	st1 := Emb_struct1{"yes", "no"}
	st2 := Emb_struct2{true, 1, 2}
	newSt1 := st1.Upgrade()
	newSt2 := st2.Upgrade()
	fmt.Println(newSt1)
	fmt.Println(newSt2)
}

func yamlMarshal(st *TestYaml) {
	out, err := yaml.Marshal(st)
	if err != nil {
		fmt.Printf("yaml_marshal failed with error:\n%s", err.Error())
		return
	}
	fmt.Printf("yaml_marshal succeeded with:\n%s", string(out))
}

func yamlUnmarshal() *TestYaml {
	content, err := ioutil.ReadFile("./test.yaml")
	if err != nil {
		return nil
	}
	var conf TestYaml
	err = yaml.Unmarshal(content, &conf)
	if err != nil {
		fmt.Printf("\nyaml_unmarshal failed with error:\n%s", err.Error())
		return nil
	}
	fmt.Printf("\nyaml_unmarshal succeeded with:\n%v", conf)
	/*for i := range conf.Other {
		fmt.Printf("\nslice within conf.Other[%v]:\n%v", i, conf.Other[i])
	}*/

	return &conf
}

func main() {
	DoUpgrade()
}

/*
func main() {
	var conf TestYaml
	conf.Reverse_xfwd4 = "on"
	conf.Embb1.Sfield1 = "xxx"
	conf.Embb1.Sfield2 = "x-o"
	conf.Embb2.present = true
	conf.Embb2.Ifield1 = 0
	conf.Embb2.Ifield2 = 0
	conf.Other = []string{"", "yes"}
	yamlMarshal(&conf)
	yamlUnmarshal()
}*/

/*
func main() {
	var result omdata.WebAppFmt
	result, _ = omconfparse.ParseWebAppFile("E:\\code\\wafom\\omconfparse\\webapp\\4.0\\webapp.yaml")
	result, _ = omconfparse.ParseWebAppFile("E:\\code\\wafom\\omconfparse\\webapp\\4.1\\webapp.yaml")
	result, _ = omconfparse.ParseWebAppFile("E:\\code\\wafom\\omconfparse\\webapp\\4.2\\webapp.yaml")
	for k, v := range result.Webapps {
		if d, ok := v.Linkage.([]interface{}); ok == true {
			fmt.Printf("site %v: %v\n", k, d)
		}
	}
	result, _ = omconfparse.ParseWebAppFile("E:\\code\\wafom\\omconfparse\\webapp\\4.3\\webapp.yaml")
	result, _ = omconfparse.ParseWebAppFile("E:\\code\\wafom\\omconfparse\\webapp\\4.3.1\\webapp.yaml")
	result, _ = omconfparse.ParseWebAppFile("E:\\code\\wafom\\omconfparse\\webapp\\4.3.2\\webapp.yaml")
	fmt.Printf("\n%v\n", result)
	return
}*/
