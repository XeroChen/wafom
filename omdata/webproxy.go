package omdata

type AppACRuleFmt struct {
	Object   string
	Action   string
	Content  []string
	Webapps  interface{}
	Match    string `yaml:",omitempty"`
	Disabled string `yaml:",omitempty"`
}
