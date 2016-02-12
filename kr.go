package main


import (
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/samples/flags"
)

// Create a PanelHolder with a 3 panels
func panelHolder(name string, theme gxui.Theme) gxui.PanelHolder {
	label := func(text string) gxui.Label {
		label := theme.CreateLabel()
		label.SetText(text)
		return label
	}

	holder := theme.CreatePanelHolder()
	holder.AddPanel(label(name+" 0 content"), name+" 0 panel")
	holder.AddPanel(label(name+" 1 content"), name+" 1 panel")
	holder.AddPanel(label(name+" 2 content"), name+" 2 panel")
	return holder
}

func appMain(driver gxui.Driver) {
	theme := flags.CreateTheme(driver)

	hSplitter := theme.CreateSplitterLayout()
	hSplitter.SetOrientation(gxui.Horizontal)
	hSplitter.AddChild(panelHolder("L", theme))
	hSplitter.AddChild(panelHolder("R", theme))

	window := theme.CreateWindow(500, 600, "Panels")
	window.SetScale(flags.DefaultScaleFactor)
	window.AddChild(hSplitter)
	window.OnClose(driver.Terminate)
}

func main() {
	gl.StartDriver(appMain)
}


