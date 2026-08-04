package config

type AppConfig struct {
	Base_Url       string
	Version_number string
	Get_action     string
	Get_All_action string
	Create_action  string
	Update_action  string
	Delete_action  string
}

func (app_config *AppConfig) CreateConfig() *AppConfig {
	var newConfig AppConfig

	newConfig.Base_Url = "http://localhost:9090/api"
	newConfig.Version_number = "/v1"
	newConfig.Get_action = "/item/%v"
	newConfig.Get_All_action = "/item"
	newConfig.Create_action = "/item"
	newConfig.Update_action = "/item/%v"
	newConfig.Delete_action = "/item/%v"

	app_config = &newConfig
	return app_config
}
