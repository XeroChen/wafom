package omdata

import (
	"gopkg.in/yaml.v2"
	"strings"
)

type AppACRuleFmt struct {
	Fetch           string // legacy object
	Fetch_param     string
	Object          string `yaml:",omitempty"`
	Action          string
	Content         []string `yaml:",omitempty"`
	Webapps         interface{}
	Match           interface{} `yaml:",omitempty"` // legacy string, current []string
	Match_param     string
	Disabled        string
	Ipgeo           []string   `yaml:",omitempty"`
	Geo_data        [][]string `yaml:",omitempty"`
	Neg             int        // legacy match: match 1; unmatch 0
	Content_display string     `yaml:",omitempty"`
}

type WebProxyFileFmt yaml.MapSlice

func (olddata AppACRuleFmt) Upgrade() (newdata AppACRuleFmt) {

	newdata.upgradeFetch(olddata)
	newdata.upgradeWebapps(olddata)
	newdata.upgradeNeg(olddata)
	newdata.upgradeAction(olddata)
	newdata.upgradeDisabled(olddata)
	return newdata
}

func (newdata *AppACRuleFmt) upgradeNeg(olddata AppACRuleFmt) {
	if s, ok := olddata.Match.(string); ok == true {
		if s == "match" {
			newdata.Neg = 1
			return
		}
	}
	newdata.Neg = 0
}

func (newdata *AppACRuleFmt) upgradeFetch(olddata AppACRuleFmt) {
	// handle Fetch
	if olddata.Object == "ext" {
		newdata.upgradeExt(olddata)
	} else if olddata.Object == "url" {
		newdata.upgradeURL(olddata)
	} else if olddata.Object == "location" {
		newdata.upgradeGeo(olddata)
	} else {
		newdata.Fetch = olddata.Object
	}
	newdata.Object = ""
}

func (newdata *AppACRuleFmt) upgradeDisabled(olddata AppACRuleFmt) {
	if olddata.Disabled == "1" {
		newdata.Disabled = "1"
	} else {
		newdata.Disabled = "0"
	}
}

func (newdata *AppACRuleFmt) upgradeAction(olddata AppACRuleFmt) {
	newdata.Action = olddata.Action
}

func (newdata *AppACRuleFmt) upgradeExt(olddata AppACRuleFmt) {
	newdata.Fetch = "path"
	newdata.Match_param = "end"
	newdata.Match = olddata.Content
}

func (newdata *AppACRuleFmt) upgradeURL(olddata AppACRuleFmt) {
	newdata.Fetch = "url"
	newdata.Match_param = "reg"
	newdata.Match = olddata.Content
}

func (newdata *AppACRuleFmt) upgradeGeo(olddata AppACRuleFmt) {
	newdata.Fetch = "location"
	newdata.Match_param = ""
	newdata.Match = olddata.Ipgeo

	var geoName string
	var shortName string

	if len(olddata.Content) > 0 {
		geoName = olddata.Content[0]
	}

	if len(olddata.Ipgeo) > 0 {
		shortNameSplit := strings.Split(olddata.Ipgeo[0], "geo.")
		if len(shortNameSplit) > 1 {
			shortName = strings.Replace(shortNameSplit[1], ".", ",", -1)
		}
	}

	if geoName != "" && shortName != "" {
		newdata.Geo_data = append(newdata.Geo_data, []string{shortName, geoName})
		newdata.Content_display = geoName
	}

}

func (newdata *AppACRuleFmt) upgradeWebapps(olddata AppACRuleFmt) {
	if s, ok := olddata.Webapps.(string); ok == true && s == "all" {
		sitelist := make([]string, 1)
		sitelist[0] = s
		newdata.Webapps = sitelist

	} else {
		newdata.Webapps = olddata.Webapps
	}
}
