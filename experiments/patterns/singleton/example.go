package main

func main() {
	logger := getLoggerInstance()

	logger.SetLogLevel(1)
	logger.Log("Hi mum!")
	logger.Log("Guacamole")
	logger.SetLogLevel(3)
	logger.Log("Hi dad!")
}
