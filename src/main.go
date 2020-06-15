package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/getlantern/systray"
)

// Container ...
const Container = "fusion-content-cache"

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	cli, err := client.NewEnvClient()
	if err != nil {
		fmt.Println("Unable to create docker client")
		panic(err)
	}

	systray.SetIcon(getIcon("assets/icon.ico"))
	systray.SetTooltip("Manage arc cache")

	init := systray.AddMenuItem("Start content cache", "Start content cache")
	stop := systray.AddMenuItem("Stop content cache", "Stop content cache")
	systray.AddSeparator()
	quit := systray.AddMenuItem("Quit", "Close")

	go func() {
		for {
			select {
			case <-init.ClickedCh:
				if !(runsContainer(Container)) {
					cli.ContainerStart(context.Background(), Container, types.ContainerStartOptions{})
				}
			case <-stop.ClickedCh:
				if runsContainer(Container) {
					cli.ContainerStop(context.Background(), Container, nil)
				}
			case <-quit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}

func onExit() {}

func runsContainer(name string) bool {
	cli, err := client.NewEnvClient()
	if err != nil {
		fmt.Println("Unable to create docker client")
		panic(err)
	}

	filter := filters.NewArgs()
	filter.Add("name", name)

	containerListOptions := types.ContainerListOptions{
		All:     false,
		Filters: filter,
	}

	containers, err := cli.ContainerList(context.Background(), containerListOptions)
	if err != nil {
		fmt.Println("Unable to get containers list")
		panic(err)
	}

	/* for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	} */

	if len(containers) > 0 {
		return true
	}

	return false
}

func getIcon(src string) []byte {
	b, err := ioutil.ReadFile(src)
	if err != nil {
		fmt.Print(err)
	}
	return b
}
