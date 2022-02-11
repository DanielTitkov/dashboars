package configs

type (
	DataConfig struct {
		Presets PresetsConfig
	}
	PresetsConfig struct {
		TasksPresetsPath string `yaml:"tasksPresetsPath"`
	}
)
