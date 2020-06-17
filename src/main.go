package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/getlantern/systray"
	"github.com/gobuffalo/packr/v2"
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

	box := packr.New("box", "./assets")

	icon, err := box.Find("icon.ico")
	if err != nil {
		fmt.Println("Couldn't find the icon")
		panic(err)
	}

	systray.SetIcon(icon)
	systray.SetTooltip("Manage fusion cache")

	restart := systray.AddMenuItem("Restart content cache", "Restart content cache")
	systray.AddSeparator()
	quit := systray.AddMenuItem("Quit", "Close")

	go func() {
		for {
			select {
			case <-restart.ClickedCh:
				if runsContainer(Container) {
					cli.ContainerRestart(context.Background(), Container, nil)
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
