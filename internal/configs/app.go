package configs

type AppConfig struct {
	SystemSummaryInterval int `yaml:"systemSummaryInterval"`
	SystemSummaryTimeout  int `yaml:"systemSummaryTimeout"`
}
