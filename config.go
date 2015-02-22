package main

import (
	"io/ioutil"
	"os/user"

	"gopkg.in/yaml.v2"
)

type RC struct {
	IP   string
	User string
	Pass string
}

func ReadConfig() (string, string, string) {
	usr, _ := user.Current()
	rcpath := usr.HomeDir + "/.regzarrc"

	buf, err := ioutil.ReadFile(rcpath)
	if err != nil {
		panic(err)
	}

	rc := new(RC)
	err = yaml.Unmarshal(buf, &rc)
	if err != nil {
		panic(err)
	}

	return rc.IP, rc.User, rc.Pass
}
