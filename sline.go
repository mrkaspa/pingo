package main

import (
	"os/exec"

	"github.com/gizak/termui"
)

const maxData = 27

type sline struct {
	sparkLines *termui.Sparklines
	data       []int
}

func createSline(url string, initData []int) sline {
	spl := termui.NewSparkline()
	spl.Data = initData
	spl.Title = url
	spl.Height = 8
	spl.LineColor = termui.ColorGreen

	spls := termui.NewSparklines(spl)
	spls.Height = 11
	spls.Width = 30
	spls.BorderFg = termui.ColorCyan
	spls.BorderLabel = "PINGO!"

	spls.Add(spl)
	return sline{
		sparkLines: spls,
		data:       initData,
	}
}

func (s *sline) updateData(elem int) {
	spl := &(s.sparkLines.Lines[0])
	dataMut := &((*spl).Data)
	*dataMut = append(*dataMut, elem)
	idx := 0
	lenData := len(*dataMut)
	if lenData >= maxData {
		idx = lenData - maxData
	}
	*dataMut = (*dataMut)[idx:]

	if elem == 1 {
		spl.LineColor = termui.ColorRed
		go exec.Command("afplay", "/System/Library/Sounds/Basso.aiff").Run()
	} else {
		spl.LineColor = termui.ColorGreen
		go exec.Command("afplay", "/System/Library/Sounds/Pop.aiff").Run()
	}
}
