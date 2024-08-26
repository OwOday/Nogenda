package nogenda

import (
	"time"

	"gioui.org/app"
	"gioui.org/io/event"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	silk "github.com/OwOday/Silk-Go/src"
)

type C = layout.Context
type D = layout.Dimensions

type Nogenda struct {
	timeZone      string
	selectedMonth time.Month
	selectedDay   int
	selectedYear  int
	silkdb_push   chan silk.RelationalNode
	silkdb_pull   chan string
	theme         *material.Theme
	ops           op.Ops
	button        widget.Clickable
}

func New(db_push chan silk.RelationalNode, db_pull chan string) *Nogenda {
	now := time.Now()
	zone, _ := now.Zone()
	year, month, day := now.Date()
	var ops op.Ops
	return &Nogenda{
		timeZone:      zone,
		selectedDay:   day,
		selectedMonth: month,
		selectedYear:  year,
		silkdb_push:   db_push,
		silkdb_pull:   db_pull,
		theme:         material.NewTheme(),
		ops:           ops,
		//really, what we need here is some way of managing a gio frame(?), passing us to a gio manager isn't going to work, we need a callback.
	}
}

func (nogenda *Nogenda) HandleEvent(gui_event event.Event) {
	switch event_type := gui_event.(type) {
	case app.FrameEvent:
		nogenda.frameUpdate(event_type)
	}
}

func (nogenda *Nogenda) frameUpdate(frame_event app.FrameEvent) {
	gtx := app.NewContext(&nogenda.ops, frame_event)
	nogenda.timeZone = nogenda.timeZone + "a"
	if nogenda.button.Clicked(gtx) {
		nogenda.timeZone = "uhoh"
	}
	// Let's try out the flexbox layout concept
	layout.Flex{
		// Vertical alignment, from top to bottom
		Axis: layout.Vertical,
		// Empty space is left at the start, i.e. at the top
		Spacing: layout.SpaceStart,
	}.Layout(gtx,
		layout.Rigid(
			func(gtx C) D {
				// ONE: First define margins around the button using layout.Inset ...
				margins := layout.Inset{
					Top:    unit.Dp(25),
					Bottom: unit.Dp(25),
					Right:  unit.Dp(35),
					Left:   unit.Dp(35),
				}
				// TWO: ... then we lay out those margins ...
				return margins.Layout(gtx,
					// THREE: ... and finally within the margins, we ddefine and lay out the button
					func(gtx C) D {
						btn := material.Button(nogenda.theme, &nogenda.button, nogenda.timeZone)
						return btn.Layout(gtx)
					},
				)
			},
		),
	)
	frame_event.Frame(gtx.Ops)
	//var startButton widget.Clickable
	//gtx := app.NewContext(&nogenda.ops, frame_event)
	//btn := material.Button(nogenda.theme, &startButton, "Start")
	//btn.Layout(gtx)
	//frame_event.Frame(gtx.Ops)
}

func (nogenda *Nogenda) SetDBPull(db_pull chan string) {

}
