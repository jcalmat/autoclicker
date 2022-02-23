package main

import "errors"

type ClickType string

const (
	Left  ClickType = "left"
	Right ClickType = "right"
)

func (c ClickType) String() string {
	return string(c)
}

type Config struct {
	Positions []Position `json:"positions"`
	Smooth    bool       `json:"smooth"`
	ClickType ClickType  `json:"clickType"`
	Frequency Frequency  `json:"frequency"`
	Debug     bool       `json:"debug"`
}

type Frequency struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (c Config) Check() error {
	if len(c.Positions) == 0 {
		return errors.New("positions field is missing")
	}

	if c.ClickType != Left && c.ClickType != Right {
		return errors.New("clickType field is not properly set")
	}

	if c.Frequency.Min == 0 || c.Frequency.Max == 0 || c.Frequency.Min >= c.Frequency.Max {
		return errors.New("frequency field is not properly set")
	}

	return nil
}
