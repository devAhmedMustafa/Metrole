package registry

type AppsRegistry struct {
	Apps []AppEntry
}

func NewAppsRegistry() *AppsRegistry {
	return &AppsRegistry{
		Apps: []AppEntry{},
	}
}

func (r *AppsRegistry) AddApp(app AppEntry) {
	r.Apps = append(r.Apps, app)
}

func (r *AppsRegistry) GetAppPort(appName string) (int, bool) {
	for _, app := range r.Apps {
		if app.Name == appName {
			return app.PORT, true
		}
	}
	return 0, false
}