package main

import cfp "../omconfparse"
import "fmt"

func main() {
	result, err := cfp.ParseYamlFile("E:\\code\\wafom\\omconfparse\\webapp.yaml")
	if result == nil || err != nil {
		fmt.Println("parse file error.")
	}
	fmt.Printf("--- t:\n%v\n\n", result)

	cfp.ParseConf(result)
	return
}
