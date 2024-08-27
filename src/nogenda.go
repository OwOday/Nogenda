package nogenda

import (
	"fmt"
	"image"
	"image/color"
	"time"

	"gioui.org/app"
	"gioui.org/io/event"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
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
		//days represented on screen should be 6 weeks sun-sat with non-existant days blurred out
		//when you first open a month, use calendar module to populate federal holidays
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

func inner_square(gtx layout.Context) []layout.FlexChild {
	var squares []layout.FlexChild
	for weeks := 0; weeks < 6; weeks++ {
		squares = append(squares, layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			fmt.Println(gtx.Constraints)
			const r = 10
			bounds := image.Rect(0, 0, gtx.Constraints.Max.X, 100)
			rrect := clip.RRect{Rect: bounds, SE: r, SW: r, NW: r, NE: r}
			var rect_color color.NRGBA
			paint.FillShape(gtx.Ops, color.NRGBA{R: 0, G: 0, B: 250, A: 150},
				clip.Stroke{
					Path:  rrect.Path(gtx.Ops),
					Width: 6,
				}.Op(),
			)
			return ColorBox(gtx, gtx.Constraints.Max, rect_color)
		}))
	}
	return squares
}

func inner_flex(gtx layout.Context) []layout.FlexChild {
	var columns []layout.FlexChild
	for days := 0; days < 7; days++ {
		columns = append(columns, layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceAround,
			}.Layout(gtx, inner_square(gtx)...)
		}))
	}
	return columns

}

func (nogenda *Nogenda) frameUpdate(frame_event app.FrameEvent) {
	gtx := app.NewContext(&nogenda.ops, frame_event)
	//days := 7
	//weeks := 6
	//var button widget.Clickable

	layout.Flex{
		Axis:    layout.Horizontal,
		Spacing: layout.SpaceSides,
	}.Layout(gtx, inner_flex(gtx)...)
	//inner_flex()...

	//make calendar
	// cal := newCalendar(*nogenda.theme)
	// component.Grid(nogenda.theme, &cal.calendar_state).Layout(gtx, 6, 7,
	// 	func(axis layout.Axis, index, constraint int) int {
	// 		smaller_side := min(frame_event.Size.X, frame_event.Size.Y)
	// 		var units int
	// 		if smaller_side == frame_event.Size.X {
	// 			units = 7
	// 		} else {
	// 			units = 6
	// 		}
	// 		return smaller_side / units
	// 	},
	// 	func(gtx C, row, col int) D {
	// 		c := color.NRGBA{R: uint8(15 * row), G: uint8(25 * col), B: uint8(row * col), A: 255}
	// 		paint.FillShape(gtx.Ops, c, clip.Rect{Max: gtx.Constraints.Max}.Op())
	// 		return D{Size: gtx.Constraints.Max}

	// // 	})
	// offset := 10
	// gtx.Reset()
	// ncols := 7
	// weight := 1 / float32(ncols)
	// var columns []layout.FlexChild
	// for i := 0; i < ncols; i++ {
	// 	columns = append(columns,
	// 		layout.Flexed(weight, func(gtx C) D {
	// 			gtx.Constraints.Max.X = gtx.Constraints.Max.X - offset
	// 			gtx.Constraints.Max.Y = gtx.Constraints.Max.Y - offset
	// 			gtx.Constraints.Min.X = gtx.Constraints.Min.X + offset
	// 			gtx.Constraints.Min.Y = gtx.Constraints.Min.Y + offset
	// 			layoutColumn(&gtx, i)
	// 			fmt.Println(gtx.Constraints)
	// 			return D{Size: gtx.Constraints.Max}
	// 		}),
	// 	)
	// }
	// layout.Flex{}.Layout(gtx, columns...)

	//Let's try out the flexbox layout concept
	// layout.Flex{
	// 	// Vertical alignment, from top to bottom
	// 	Axis: layout.Vertical,
	// 	// Empty space is left at the start, i.e. at the top
	// 	Spacing: layout.SpaceStart,
	// }.Layout(gtx,
	// 	cal.calendar_state,
	// layout.Rigid(
	// 	func(gtx C) D {
	// 		// ONE: First define margins around the button using layout.Inset ...
	// 		margins := layout.Inset{
	// 			Top:    unit.Dp(25),
	// 			Bottom: unit.Dp(25),
	// 			Right:  unit.Dp(35),
	// 			Left:   unit.Dp(35),
	// 		}
	// 		// TWO: ... then we lay out those margins ...
	// 		return margins.Layout(gtx,
	// 			// THREE: ... and finally within the margins, we ddefine and lay out the button
	// 			func(gtx C) D {
	// 				btn := material.Button(nogenda.theme, &nogenda.button, nogenda.timeZone)
	// 				return btn.Layout(gtx)
	// 			},
	// 		)
	// 	},
	// ),
	// )
	frame_event.Frame(gtx.Ops)
	//var startButton widget.Clickable
	//gtx := app.NewContext(&nogenda.ops, frame_event)
	//btn := material.Button(nogenda.theme, &startButton, "Start")
	//btn.Layout(gtx)
	//frame_event.Frame(gtx.Ops)
}
