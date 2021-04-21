package environments

type Configuration struct {
	ExchangeratesapiRoot	string
	AccessKey 						string 	`env:"EXCHANGERATESAPI_ACCESS_KEY"`
}