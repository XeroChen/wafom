package main

import (
	"../omconfparse"
	"flag"
	"fmt"
	"runtime"
	"strings"
)

func handleWebapp() {

	var defaultFile string

	if runtime.GOOS == "windows" {
		defaultFile = "./webapp.yaml"
	} else {
		defaultFile = "/waf/config/misc/webapp.yaml"
	}

	oldfile := flag.String("old", defaultFile, "The pathname of the file to be upgraded.")
	newfile := flag.String("new", defaultFile, "The pathname of the file to be generated after upgrading.")

	flag.Parse()

	webappdata, err := omconfparse.ParseWebAppFile(*oldfile)

	if err != nil {
		fmt.Println("parse file error.")
	}

	omconfparse.GenWebAppFile(*newfile, webappdata)
}

func handleWebProxy() {
	var defaultFile string

	if runtime.GOOS == "windows" {
		defaultFile = "./webproxy.yaml"
	} else {
		defaultFile = "/waf/config/misc/webproxy.yaml"
	}

	oldfile := flag.String("old", defaultFile, "The pathname of the file to be upgraded.")
	newfile := flag.String("new", strings.Replace(defaultFile, "webproxy.yaml", "webproxy_transparent.yaml", -1), "The pathname of the file to be generated after upgrading.")

	flag.Parse()

	webproxydata, err := omconfparse.ParseWebProxyFile(*oldfile)

	if err != nil {
		fmt.Printf("ParseWebProxyFile returns:%s", err.Error())
		return
	}

	omconfparse.GenWebProxyFile(*newfile, webproxydata)
}
func main() {

	typeFlag := flag.String("type", "", "webapp or webproxy to be upgrade.")
	flag.Parse()

	switch *typeFlag {
	case "webapp":
		handleWebapp()
	case "webproxy":
		handleWebProxy()
	default:
		fmt.Printf("\ntype %v not accepted.", *typeFlag)
	}

	return
}
