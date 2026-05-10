package presets

type Preset struct {
	Name        string
	Description string
	AlwaysOn    []string
	Manual      []string
	HooksOn     []string
	SpecTemplate string
}

var All = []Preset{
	{
		Name:         "api",
		Description:  "REST API — net/http or chi router, middleware, handlers, JSON responses",
		AlwaysOn:     []string{"01", "02", "03", "04", "05", "06", "07", "08", "10"},
		Manual:       []string{"09"},
		HooksOn:      []string{"01", "02", "03", "04", "05"},
		SpecTemplate: "_API-ENDPOINT-TEMPLATE",
	},
	{
		Name:         "cli",
		Description:  "CLI tool — cobra or flag, subcommands, config file, stdout/stderr hygiene",
		AlwaysOn:     []string{"01", "02", "03", "04", "05", "06", "07", "09"},
		Manual:       []string{"08", "10"},
		HooksOn:      []string{"01", "02", "03", "05"},
		SpecTemplate: "_FEATURE-TEMPLATE",
	},
	{
		Name:         "svc",
		Description:  "Microservice — gRPC or HTTP, health checks, graceful shutdown, telemetry",
		AlwaysOn:     []string{"01", "02", "03", "04", "05", "06", "07", "08", "10"},
		Manual:       []string{"09"},
		HooksOn:      []string{"01", "02", "03", "04", "05"},
		SpecTemplate: "_API-ENDPOINT-TEMPLATE",
	},
}

func Find(name string) (Preset, bool) {
	for _, p := range All {
		if p.Name == name {
			return p, true
		}
	}
	return Preset{}, false
}
