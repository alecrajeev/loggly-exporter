package config

import (
	"fmt"
	"io/ioutil"
	yaml "gopkg.in/yaml.v2"
)

type Conf struct {
	LogglySubDomain string `yaml:"loggly_subdomain"`
	ListenerPort string `yaml:"port"`
	Query string `yaml:"query"`
}

func (c *Conf) Load(file *string) error {
	yamlFile, err := ioutil.ReadFile(*file)

	if err != nil {
		fmt.Println("Error parsing yaml")
		fmt.Println(err)
		return err
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println("Error parsing yaml")
		return err
	}

	fmt.Printf("SubDomain: %v\n", c.LogglySubDomain)
	return nil
}