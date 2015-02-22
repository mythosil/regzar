package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"

	"github.com/codegangsta/cli"
	"github.com/mattbaird/http-digest-auth-client"
)

var Commands = []cli.Command{
	commandChannel,
	commandVolume,
	commandPower,
	commandStatus,
}

var commandChannel = cli.Command{
	Name:   "channel",
	Usage:  "change channel",
	Action: doChannel,
}

var commandVolume = cli.Command{
	Name:   "volume",
	Usage:  "up/down/mute volume",
	Action: doVolume,
}

var commandPower = cli.Command{
	Name:   "power",
	Usage:  "up/down power",
	Action: doPower,
}

var commandStatus = cli.Command{
	Name:   "status",
	Usage:  "show status",
	Action: doStatus,
}

func doChannel(c *cli.Context) {
	argChannel := c.Args().First()
	if argChannel == "" {
		cli.ShowCommandHelp(c, "channel")
		os.Exit(1)
	}

	if argChannel == "prev" {
		sendRemoteRequest("40bf1f")
		return
	} else if argChannel == "next" {
		sendRemoteRequest("40bf1b")
		return
	}

	ch, err := strconv.Atoi(argChannel)
	if err != nil {
		panic(err)
	}

	key := fmt.Sprintf("40bf%02x", ch)

	sendRemoteRequest(key)
}

func doVolume(c *cli.Context) {
	argAction := c.Args().First()
	if argAction == "" {
		cli.ShowCommandHelp(c, "volume")
		os.Exit(1)
	}

	switch argAction {
	case "up":
		sendRemoteRequest("40bf1a")
	case "down":
		sendRemoteRequest("40bf1e")
	case "mute":
		sendRemoteRequest("40bf10")
	}
}

func doPower(c *cli.Context) {
	sendRemoteRequest("40bf12")
}

func doStatus(c *cli.Context) {
	sendStatusRequest()
}

func sendRemoteRequest(key string) {
	ip, user, pass := ReadConfig()

	values := url.Values{}
	values.Add("key", key)

	uri := "http://" + ip + "/remote/remote.htm?" + values.Encode()
	_, err := http_digest_auth.Auth(user, pass, uri)

	if err != nil {
		panic(err)
	}
}

func sendStatusRequest() {
	ip, user, pass := ReadConfig()

	uri := "http://" + ip + "/remote/status.htm"
	resp, err := http_digest_auth.Auth(user, pass, uri)

	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
