package app

// DisableHealthMonitorForTesting disables the health monitor for testing.
func (app *App) DisableHealthMonitorForTesting() {
	app.DaemonHealthMonitor.DisableForTesting()
}
