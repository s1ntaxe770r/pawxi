package utils

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/s1ntaxe770r/pawxi/proxy"
)

func Vizualize(routes []proxy.Route) {
	t := table.NewWriter()
	tTemp := table.Table{}
	tTemp.Render()
	for _, route := range routes {
		t.AppendRow([]interface{}{color.YellowString(route.Path), color.MagentaString("=>"), color.GreenString(route.Destination)})
	}
	t.SetCaption("current routes setup.\n")
	fmt.Println(t.Render())
}
