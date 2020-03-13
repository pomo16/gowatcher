package model

type Config struct {
	ElasticSearch ESConfig `yaml:"elasticsearch"`
}

type ESConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}
