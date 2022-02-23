package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/go-vgo/robotgo"
)

var (
	version string
)

func main() {
	fmt.Printf(`
     _____ _ _      _             
    / ____| (_)    | |            
   | |    | |_  ___| | _____ _ __ 
   | |    | | |/ __| |/ / _ \ '__|
   | |____| | | (__|   <  __/ |   
    \_____|_|_|\___|_|\_\___|_|   
                                  %s
	Ctrl+C or close window to quit
	`, version)

	// config
	configFile, err := os.ReadFile("./config.json")
	if err != nil {
		fmt.Println("failed to read config file. Please quit and retry.")
		select {}
	}

	var config Config
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		fmt.Println("failed to unmarshal config file. Please quit and retry.")
		select {}
	}
	debuglog(config, fmt.Sprintf("Config file loaded: %#v\n", config))

	if err := config.Check(); err != nil {
		fmt.Printf("error in config file: %s\n", err.Error())
		select {}
	}

	// init random
	rand.Seed(time.Now().UnixNano())

	// rollout list
	rollout := make([]Position, 0)
	rollout = append(rollout, config.Positions...)

	randomFrequency := rand.Intn(config.Frequency.Max-config.Frequency.Min) + config.Frequency.Min
	debuglog(config, fmt.Sprintf("Next tick in %ds", randomFrequency))
	ticker := time.NewTicker(time.Second * time.Duration(randomFrequency))
	defer ticker.Stop()
	go func() {
		for range ticker.C {
			// compute random
			r := rand.Intn(len(rollout))

			// go to current position and click
			pos := rollout[r]
			move(pos, config.Smooth)
			robotgo.Click(config.ClickType.String())

			debuglog(config, fmt.Sprintf("Going to [%d, %d]", pos.X, pos.Y))

			// pop value from rollout
			rollout = append(rollout[:r], rollout[r+1:]...)
			// eventually reset the rollout
			if len(rollout) == 0 {
				rollout = append(rollout, config.Positions...)
			}

			// reset the ticker
			randomFrequency = rand.Intn(config.Frequency.Max-config.Frequency.Min) + config.Frequency.Min
			debuglog(config, fmt.Sprintf("Next tick in %ds", randomFrequency))
			ticker = time.NewTicker(time.Second * time.Duration(randomFrequency))
		}
	}()

	// #Wait for signal
	select {}
}

func debuglog(c Config, s string) {
	if c.Debug {
		fmt.Println(s)
	}
}

func move(pos Position, smooth bool) {
	if smooth {
		robotgo.MoveSmooth(pos.X, pos.Y, 1.0, 0.3)
	} else {
		robotgo.Move(pos.X, pos.Y)
	}
}
