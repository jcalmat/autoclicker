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
                                  @%s
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
	if config.Debug {
		fmt.Printf("Config file loaded: %#v\n", config)
	}

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
	if config.Debug {
		fmt.Printf("Next tick in %dmn\n", randomFrequency)
	}
	ticker := time.NewTicker(time.Minute * time.Duration(randomFrequency))
	go func() {
		for {
			select {
			case <-ticker.C:
				// compute random
				r := rand.Intn(len(rollout))

				// go to current position and click
				pos := rollout[r]
				robotgo.Move(pos.X, pos.Y)
				robotgo.Click(config.ClickType.String())

				if config.Debug {
					fmt.Printf("Going to [%d, %d]\n", pos.X, pos.Y)
				}

				// pop value from rollout
				rollout = append(rollout[:r], rollout[r+1:]...)
				// eventually reset the rollout
				if len(rollout) == 0 {
					rollout = append(rollout, config.Positions...)
				}

				// reset the ticker
				ticker.Stop()
				randomFrequency = rand.Intn(config.Frequency.Max-config.Frequency.Min) + config.Frequency.Min
				if config.Debug {
					fmt.Printf("Next tick in %dmn\n", randomFrequency)
				}
				ticker = time.NewTicker(time.Minute * time.Duration(randomFrequency))
			}
		}
	}()

	// #Wait for signal
	select {}
}
