
package config

var Port int

func ReadConfig() bool {
	Port = 8080
	return true
}
