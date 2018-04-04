package omdata

import "strings"

type AppACRuleFmt struct {
	Fetch       string // legacy object
	Fetch_param string
	Object      string `yaml:",omitempty"`
	Action      string
	Content     []string
	Webapps     interface{}
	Match       interface{} `yaml:",omitempty"` // legacy string, current []string
	Match_param string
	Disabled    string
	Ipgeo       []string `yaml:",omitempty"`
	Geo_data    [][2]string
	Neg         string // legacy match: match 1; unmatch 0
}

func (newdata *AppACRuleFmt) Upgrade(olddata AppACRuleFmt) {

	newdata.upgradeFetch(olddata)
	newdata.upgradeWebapps(olddata)
	newdata.upgradeNeg(olddata)
	newdata.upgradeAction(olddata)
}

func (newdata *AppACRuleFmt) upgradeNeg(olddata AppACRuleFmt) {
	if s, ok := olddata.Match.(string); ok == true {
		if s == "match" {
			newdata.Neg = "1"
		} else {
			newdata.Neg = "0"
		}
	}
}

func (newdata *AppACRuleFmt) upgradeFetch(olddata AppACRuleFmt) {
	// handle Fetch
	if olddata.Object == "ext" {
		newdata.upgradeExt(olddata)
	} else if olddata.Object == "url" {
		newdata.upgradeURL(olddata)
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

	if len(olddata.Content) > 0 {
		oldGeo := olddata.Content[0]
	}

}

func (newdata *AppACRuleFmt) upgradeWebapps(olddata AppACRuleFmt) {
	if s, ok := olddata.Webapps.(string); ok == true {
		if s == "all" {
			sitelist := make([]string, 1)
			sitelist[0] = s
			newdata.Webapps = sitelist
		}
	} else {
		newdata.Webapps = olddata.Webapps
	}
}
