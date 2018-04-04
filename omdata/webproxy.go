package omdata

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
	Neg         string   // legacy match: match 1; unmatch 0
}

func (newdata *AppACRuleFmt) Upgrade(olddata *AppACRuleFmt) {

	// handle Fetch
	if olddata.Object == "ext" {
		newdata.upgradeExt(olddata)
	} else if olddata.Object == "url" {
		newdata.upgradeUrl(olddata)

	} else {
		newdata.Fetch = olddata.Object
	}

	olddata.Object = ""
	if s, ok := olddata.Match.(string); ok == true {
		if s == "match" {
			newdata.Neg = "1"
		} else {
			newdata.Neg = "0"
		}
	}
}

func (newdata *AppACRuleFmt) upgradeExt(olddata *AppACRuleFmt) {
	newdata.Fetch = "path"
	newdata.Match_param = "end"
	newdata.Match = olddata.Content

}

func (newdata *AppACRuleFmt) upgradeUrl(olddata *AppACRuleFmt) {
	newdata.Fetch = "url"
	newdata.Match_param = "reg" 
	newdata.Match = olddata.Content
	if s, ok := olddata.Match.(string); ok == true {
		if s == "match" {
			newdata.Neg = "1"
		} else {
			newdata.Neg = "0"
		}
	}
}
