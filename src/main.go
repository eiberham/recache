package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/getlantern/systray"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	cli, err := client.NewEnvClient()
	if err != nil {
		fmt.Println("Unable to create docker client")
		panic(err)
	}

	if cli != nil {
		fmt.Println("Cool")
	}

	containerListOptions := types.ContainerListOptions{
		All: true,
	}

	containers, err := cli.ContainerList(context.Background(), containerListOptions)
	if err != nil {
		fmt.Println("Unable to get containers list")
		panic(err)
	}
	//fmt.Print(containers)
	// cli.ContainerStart(context.Background(), )

	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	}

	systray.SetIcon(getIcon("assets/icon.ico"))
	systray.SetTooltip("Turn off arc cache")

	systray.AddMenuItem("Restart content cache", "Restart content cache")
	systray.AddSeparator()
	quit := systray.AddMenuItem("Quit", "Close")

	go func() {
		for {
			select {
			case <-quit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}

func onExit() {
	// Nothing here for the time being
}

func getIcon(s string) []byte {
	b, err := ioutil.ReadFile(s)
	if err != nil {
		fmt.Print(err)
	}
	return b
}
