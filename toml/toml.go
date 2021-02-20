package toml

import (
	"os"
	"path/filepath"
	"reflect"

	"github.com/triamazikamno/goconfig"
	"github.com/pelletier/go-toml"
)

func init() {
	f := goconfig.Fileformat{
		Extension:   ".toml",
		Save:        SaveTOML,
		Load:        LoadTOML,
		PrepareHelp: PrepareHelp,
	}
	goconfig.Formats = append(goconfig.Formats, f)
}

// LoadTOML config file
func LoadTOML(config interface{}) (err error) {
	configFile := filepath.Join(goconfig.Path, goconfig.File)
	_, err = os.Stat(configFile)
	if os.IsNotExist(err) && !goconfig.FileRequired {
		err = nil
		return
	} else if err != nil {
		return
	}
	var tree *toml.Tree
	tree, err = toml.LoadFile(configFile)
	if err != nil {
		return
	}
	err = tree.Unmarshal(config)
	return
}

// SaveTOML config file
func SaveTOML(config interface{}) (err error) {
	_, err = os.Stat(goconfig.Path)
	if os.IsNotExist(err) {
		err = os.Mkdir(goconfig.Path, os.ModePerm)
		if err != nil {
			return
		}
	} else if err != nil {
		return
	}
	configFile := filepath.Join(goconfig.Path, goconfig.File)
	_, err = os.Stat(goconfig.Path)
	if err != nil {
		return
	}
	file, err := os.Create(configFile)
	if err != nil {
		return
	}
	defer file.Close()
	cfg := reflect.ValueOf(config).Elem()
	err = toml.NewEncoder(file).Encode(cfg)
	return
}

// PrepareHelp return help string for this file format.
func PrepareHelp(config interface{}) (help string, err error) {
	var byt []byte
	cfg := reflect.ValueOf(config).Elem()
	byt, err = toml.Marshal(cfg)
	if err != nil {
		return
	}
	help = string(byt)
	return
}
