package config

type Config struct {
	JsonRpc *JsonRpcConfig
}

type JsonRpcConfig struct {
	Url string
}
