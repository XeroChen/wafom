package main

import "../omconfparse"
import "fmt"
import "flag"
import "runtime"

func main() {
	var defaultFile string

	if runtime.GOOS == "windows" {
		defaultFile = "./webapp.yaml"
	} else {
		defaultFile = "/waf/config/misc/webapp.yaml"
	}

	oldfile := flag.String("old", defaultFile, "The pathname of the webapp.yaml file to be upgraded.")
	newfile := flag.String("new", defaultFile, "The pathname of the webapp.yaml file to be generated after upgrading.")

	flag.Parse()

	webappdata, err := omconfparse.ParseWebAppFile(*oldfile)

	if err != nil {
		fmt.Println("parse file error.")
	}

	result := omconfparse.ParseWebAppData(webappdata)
	omconfparse.GenWebAppFile(*newfile, result)
	return
}
