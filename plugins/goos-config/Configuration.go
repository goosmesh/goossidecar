package goos_config

const (
	DEFAULT_GOOS_ADDRESS = "goos:8080"
)

type Configuration struct {
	goosAddress string
}

func CreateConfiguration() Configuration {
	return Configuration{
		goosAddress: DEFAULT_GOOS_ADDRESS,
	}
}