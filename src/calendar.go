package nogenda

import (
	"fmt"
	"image"
	"image/color"
	"math/rand/v2"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"gioui.org/x/component"
)

type calendar struct {
	calendar_state component.GridState
	calendar_style component.GridStyle
}

func newCalendar(theme material.Theme) *calendar {
	var grid component.GridState
	return &calendar{
		calendar_state: grid,
		calendar_style: component.Grid(&theme, &grid),
	}
}

func layoutColumn(gtx *layout.Context, day int) {
	var children []layout.FlexChild
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		fmt.Println(day)
		children = append(children,
			layout.Rigid(func(layout.Context) D {
				layoutRect(gtx, i, day)
				return D{Size: gtx.Constraints.Max}
			}),
		)
	}
	layout.Flex{Axis: layout.Vertical}.Layout(*gtx, children...)
}

func layoutRect(gtx *layout.Context, week int, day int) {
	min_y := (gtx.Constraints.Min.Y / 5) * (week)
	min_x := (gtx.Constraints.Min.X / 5) * (week)
	max_y := (gtx.Constraints.Max.Y / 5) * (week)
	max_x := (gtx.Constraints.Max.X / 5) * (week)
	fmt.Printf("%d %d %d %d\n", min_x, min_y, max_x, max_y)
	layout.UniformInset(unit.Dp(0)).Layout(*gtx, func(C) D {
		//colorMultiple := (number + 1) % 2
		//fmt.Println(gtx.Constraints)
		const r = 10
		bounds := image.Rect(min_x, min_y, max_x, max_y)
		rrect := clip.RRect{Rect: bounds, SE: r, SW: r, NW: r, NE: r}
		var rect_color color.NRGBA
		//if colorMultiple == 1 {
		//	rect_color = color.NRGBA{R: 250, G: 250, B: 250, A: 150}
		//} else {
		//	rect_color = color.NRGBA{R: uint8(randRange(0, 255) * colorMultiple), G: uint8(randRange(0, 255) * colorMultiple), B: uint8(randRange(0, 255) * colorMultiple), A: 150}
		//}
		paint.FillShape(gtx.Ops, color.NRGBA{R: 250, G: 150, B: 250, A: 150},
			clip.Stroke{
				Path:  rrect.Path(gtx.Ops),
				Width: 6,
			}.Op(),
		)
		return ColorBox(*gtx, gtx.Constraints.Max, rect_color)
	})
}

func ColorBox(gtx layout.Context, size image.Point, color color.NRGBA) layout.Dimensions {
	defer clip.Rect{Max: size}.Push(gtx.Ops).Pop()
	paint.ColorOp{Color: color}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	return layout.Dimensions{Size: size}
}

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}
