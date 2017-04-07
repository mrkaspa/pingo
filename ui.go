package main

import (
	"github.com/gizak/termui"
)

type ui struct {
	sl  sline
	mon monitor
}

func createUI() ui {
	mon := newMonitor("google", "http://localhost:4000/v1/health")
	return ui{
		mon: mon,
		sl:  createSline(mon.url, []int{}),
	}
}

func (u ui) render() {
	termui.Render(u.sl.sparkLines)
}

func (u *ui) run() {
	err := termui.Init()
	if err != nil {
		panic(err)
	}
	defer termui.Close()

	u.handleEvents(u.mon)
	u.render()
	termui.Loop()
}

func (u *ui) handleEvents(m monitor) {
	termui.Merge(m.evtName(), m.notifyRequestSec)

	termui.Handle(m.evtName(), func(e termui.Event) {
		evt := e.Data.(evtHTTP)
		elem := 10
		if evt.err != nil {
			elem = 1
		}
		u.sl.updateData(elem)
		u.render()
	})

	termui.Handle("/sys/kbd/C-c", func(termui.Event) {
		termui.StopLoop()
	})
}
