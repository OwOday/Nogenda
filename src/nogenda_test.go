package nogenda

import (
	"fmt"
	"testing"

	"gioui.org/app"
	"gioui.org/unit"
	silk "github.com/OwOday/Silk-Go/src"
)

func TestGui(t *testing.T) {
	//make a gui module like a consumer would
	go func() {
		// create new window
		w := new(app.Window)
		w.Option(app.Title("Test Window"))
		w.Option(app.Size(unit.Dp(400), unit.Dp(600)))

		//use the nogenda
		//silkdb := silk.New()
		var db_push chan silk.RelationalNode
		var db_pull chan string
		nogenda := New(db_push, db_pull)

		// listen for events in the window
		for {
			//in a real application setup tabs here and switch to other eventhandlers, also handle close event
			nogenda.HandleEvent(w.Event())
		}
	}()
	app.Main()
}

func dump(a any) {
	fmt.Printf("%+v\n", a)
}
