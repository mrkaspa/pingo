package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gizak/termui"
)

type monitor struct {
	name             string
	url              string
	periodicity      int
	notifyRequestSec chan termui.Event
}

type evtHTTP struct {
	err error
}

func newMonitor(name, url string) monitor {
	m := monitor{
		name:             name,
		url:              url,
		periodicity:      1,
		notifyRequestSec: make(chan termui.Event),
	}
	go m.run()
	return m
}

func (m *monitor) run() {
	go func() {
		for {
			time.Sleep(time.Second * time.Duration(m.periodicity))
			err := m.callHTTP()
			e := termui.Event{}
			e.Type = "monitor"
			e.Path = m.evtName()
			e.Time = time.Now().Unix()
			evt := evtHTTP{}
			if err != nil {
				evt.err = err
			}
			e.Data = evt
			m.notifyRequestSec <- e
		}
	}()
}

func (m monitor) evtName() string {
	return "/monitor/" + m.name
}

func (m monitor) callHTTP() error {
	response, err := http.Get(m.url)
	if err != nil {
		return err //errors.WithStack(err)
	}
	if response.StatusCode >= 400 {
		return fmt.Errorf("Invalid state %s", response.Status)
	}
	return nil
}
