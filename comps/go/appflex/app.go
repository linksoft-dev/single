package appflex

import (
	log "github.com/sirupsen/logrus"
)

var applicationName string
var adapters []Adapter
var apps []App

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
	GetApps() []App
}

// Start function that starts all the prymary adapters
func Start(appName string) {
	applicationName = appName
	ch := make(chan bool, 1)
	defer func() {
		log.Infof("Application '%s' has been started", applicationName)
		for _, adapter := range adapters {
			for _, app := range adapter.GetApps() {
				app.AfterStart()
			}
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

// AddAdapters add the adapter to the list
func AddAdapters(adapter ...Adapter) {
	adapters = append(adapters, adapter...)
}
func AddApp(app App) {
	apps = append(apps, app)
}
