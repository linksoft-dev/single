package api

import (
	log "github.com/sirupsen/logrus"
	"sync"
)

var (
	applicationName string
	adapters        []Adapter
	apps            []App
)

// App struct that define which method the app instances should have
type App interface {
	// BeforeStart execute before the application is loaded
	BeforeStart()
	// AfterStart execute when the servers are started, so, after all the applications is loaded
	AfterStart()
}

// Adapter struct that define the adapter interface at abstract way
type Adapter interface {
	Run() error
	GetName() string
	GetApps() []App
}

// Start function that starts all the primary adapters
func Start(appName string) {
	log.Infof("Starting API '%s'", appName)

	log.Info("Executing BeforeStart events")
	for _, app := range apps {
		app.BeforeStart()
	}
	log.Info("Completed all BeforeStart events")

	// execute all adapters and wait
	log.Info("Starting adapters")
	wg := &sync.WaitGroup{}
	for idx := range adapters {
		adapter := adapters[idx]
		if adapter == nil {
			continue
		}
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			if err := adapter.Run(); err == nil {
				log.Infof("Adapter '%s' has been started with '%d' apps", adapter.GetName(), len(adapter.GetApps()))
				wg.Done()
			}
		}(wg)
	}
	wg.Wait()
	log.Info("Completed all adapters")

	//run all AfterStart of all application after startup all adapters,
	//ch is used to block main thread
	log.Info("Executing AfterStart events")
	for _, app := range apps {
		app.AfterStart()
	}
	log.Info("Completed all AfterStart events")

	// block main thread
	wg.Add(1)
	log.Infof("Application '%s' has been started", appName)
	wg.Wait()
}

// AddAdapter add the adapter to the list
func AddAdapter(adapter Adapter) {
	adapters = append(adapters, adapter)
}
func AddApp(app App) {
	apps = append(apps, app)
}
