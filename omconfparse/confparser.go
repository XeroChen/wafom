package omconfparse

import (
	"../omdata"
	"bytes"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
)

func fileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func ParseWebProxyFile(filepath string) (result omdata.WebProxyFileFmt, err error) {

	var newFile = []string{"webproxy_transparent.yaml", "webproxy_reverse.yaml"}
	var dirPath = path.Dir(filepath)

	if fileExist(path.Join(dirPath, newFile[0])) || fileExist(path.Join(dirPath, newFile[1])) {
		fmt.Println("webproxy_transparent.yaml or webproxy_reverse.yaml exists. Processing quit.")
		return omdata.WebProxyFileFmt{}, fmt.Errorf("new file aready exists")
	}

	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return omdata.WebProxyFileFmt{}, err
	}

	/* 由于规则依赖于顺序，通过yaml.MapSlice解析出Key的顺序 */
	var confTmp yaml.MapSlice
	err = yaml.Unmarshal(content, &confTmp)

	var keyOrder []int

	for k := range confTmp {
		if i, ok := confTmp[k].Key.(int); ok == true {
			keyOrder = append(keyOrder, i)
		}
	}

	/* 按照Key顺序升级数据和重构yaml.MapSlice */
	var conf map[int]omdata.AppACRuleFmt
	err = yaml.Unmarshal(content, &conf)

	for _, idkey := range keyOrder {
		newdata := conf[idkey].Upgrade()
		result = append(result, yaml.MapItem{Key: idkey, Value: newdata})
	}
	err = nil
	return
}

func GenWebProxyFile(newfile string, newdata omdata.WebProxyFileFmt) bool {

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

func ParseWebAppFile(filepath string) (result omdata.WebAppFmt, err error) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return omdata.WebAppFmt{}, err
	}

	var conf omdata.WebAppFmt
	err = yaml.Unmarshal(content, &conf)

	return conf.Upgrade(), err
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
