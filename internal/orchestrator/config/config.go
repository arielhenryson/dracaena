package config

import "time"

// Config is the engine run time configuration
type Config struct {
	// time to check if workflow is done
	LifeCycleInterval time.Duration

	// time to check if job dependence is finish
	JobWaitInterval time.Duration
}

func MainConfig() *Config {
	conf := Config{
		LifeCycleInterval: 500,
		JobWaitInterval: 1000,
	}

	return &conf
}
