package main

func main() {
	u := createUI()
	u.run()
	// err := termui.Init()
	// if err != nil {
	// 	panic(err)
	// }
	// defer termui.Close()
	//
	// data := []int{1, 2, 3}
	//
	// spl3 := termui.NewSparkline()
	// spl3.Data = data
	// spl3.Title = "Enlarged Sparkline"
	// spl3.Height = 8
	// spl3.LineColor = termui.ColorYellow
	//
	// spls2 := termui.NewSparklines(spl3)
	// spls2.Height = 11
	// spls2.Width = 30
	// spls2.BorderFg = termui.ColorCyan
	// spls2.BorderLabel = "Tweeked Sparkline"
	//
	// termui.Merge("/timer/2s", termui.NewTimerCh(time.Second*2))
	//
	// termui.Handle("/timer/2s", func(e termui.Event) {
	// 	cnt := e.Data.(termui.EvtTimer)
	// 	elem := int(cnt.Count) % 2
	// 	dataMut := &spls2.Lines[0].Data
	// 	*dataMut = append(*dataMut, elem)
	// 	idx := 0
	// 	lenData := len(*dataMut)
	// 	if lenData >= 10 {
	// 		idx = lenData - 10
	// 	}
	// 	*dataMut = (*dataMut)[idx:]
	// 	if elem == 0 {
	// 		go exec.Command("afplay", "/System/Library/Sounds/Basso.aiff").Run()
	// 	} else {
	// 		go exec.Command("afplay", "/System/Library/Sounds/Pop.aiff").Run()
	// 	}
	//
	// 	termui.Render(spls2)
	// })
	//
	// termui.Handle("/sys/kbd/C-c", func(termui.Event) {
	// 	termui.StopLoop()
	// })
	//
	// termui.Loop()

	// p.Handle("/timer/1s", func(e ui.Event) {
	// 	cnt := e.Data.(ui.EvtTimer)
	// 	if cnt.Count%2 == 0 {
	// 		p.TextFgColor = ui.ColorRed
	// 	} else {
	// 		p.TextFgColor = ui.ColorWhite
	// 	}
	// })

	// ui.Merge("/timer/2s", termui.NewTimerCh(time.Second*2))
	//
	// list.Handle("/timer/2s", func(e ui.Event) {
	// 	t := e.Data.(ui.EvtTimer)
	// 	i := int(t.Count)
	// 	list.Items = strs[i%9:]
	// 	ui.Render(p, list, g, sp, lc, bc, lc1, p1)
	// })

	// ui.Handle("/timer/1s", func(e ui.Event) {
	// 	t := e.Data.(ui.EvtTimer)
	// 	draw(int(t.Count))
	// })
}
