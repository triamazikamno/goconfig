package goconfig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const configPath = "./"
const configFile = "./config.json"

// Configuration struct for main sistem
type Configuration struct {
	Name  string
	Value interface{}
}

// Config instantiate the system settings, all settings should be read in system load.
var Config = Configuration{}

func loadConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal("loadConfig open config.json:", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		log.Fatal("loadConfig Decode:", err)
	}
}

func saveConfig() {
	fmt.Println("init")
	_, err := os.Stat(configPath)

	if os.IsNotExist(err) {
		os.Mkdir(configPath, 0700)
	}

	_, err = os.Stat(configFile)

	m := Configuration{
		Name:  "Local Name",
		Value: 8080}

	fmt.Println(m)

	b, err := json.MarshalIndent(m, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(configFile, b, 0644)
	if err != nil {
		log.Fatal(err)
	}

}
