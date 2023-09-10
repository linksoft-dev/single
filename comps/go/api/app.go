package api

import (
	log "github.com/sirupsen/logrus"
)

var (
	applicationName string
	adapters        []Adapter
	apps            []App
)

// App struct that define which method the app instances should have
type App interface {
	// AfterLoad execute after the app is loaded
	AfterLoad()
	// AfterStart execute when the servers are started, so, after all the applications is loaded
	AfterStart()
}

// Adapter struct that define the adapter interface at abstract way
type Adapter interface {
	Run()
}

// Start function that starts all the prymary adapters
func Start(appName string) {
	applicationName = appName
	defer func() {
		log.Info("Application '%s' has been started", applicationName)
	}()

	ch := make(chan bool, 1)
	defer func() {
		for _, app := range apps {
			app.AfterStart()
		}
		<-ch
	}()

	for _, adapter := range adapters {
		if adapter != nil {
			go func() {
				adapter.Run()
			}()
		}
	}
}

// AddAdapter add the adapter to the list
func AddAdapter(adapter Adapter) {
	adapters = append(adapters, adapter)
}
func AddApp(app App) {
	apps = append(apps, app)
}
