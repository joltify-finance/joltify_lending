package app

// DisableHealthMonitorForTesting disables the health monitor for testing.
func (app *App) DisableHealthMonitorForTesting() {
	if app.DaemonHealthMonitor == nil {
		return
	}
	app.DaemonHealthMonitor.DisableForTesting()
}
