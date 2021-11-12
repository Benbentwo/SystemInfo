package bootstrap

import (
	"encoding/json"
	"github.com/Benbentwo/Windows10BootStrapper/pkg/common/log"
	"github.com/Benbentwo/utils/util"
	"github.com/pkg/errors"
	"io/ioutil"
)

//const (
//	// TODO fix this
//	SUPPORTED_PACKAGEMANAGERS [...]string = []string{
//		"apk",
//		"brew",
//		"chocolatey",
//	}
//)
gol

func LoadBootstrapConfig(inputFile string) (*[]BootstrapConfig, error) {
	var bootstrapConfig *[]BootstrapConfig

	exists, err := util.FileExists(inputFile)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("File inputted does not exist or could not be found")
	}
	//f, err := os.Open(inputFile)
	//
	//if err != nil {
	//	return nil, errors.Errorf("Error Opening file: %s, Error: %s", inputFile, err)
	//}
	//defer f.Close()
	//


	filebytes, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(filebytes, bootstrapConfig)
	if err != nil {
		return nil, err
	}
	log.Logger().Infof("Loaded the following config: %v", bootstrapConfig)
	return bootstrapConfig, nil
}
//
//func (bsc *BootstrapConfig) LoadFile(inputFile string) (*BootstrapConfig, error) {
//
//	filebytes, err := os.ReadFile(inputFile)
//	if err != nil {
//		return nil, err
//	}
//
//	err = json.Unmarshal(filebytes, bsc)
//	if err != nil {
//		return nil, err
//	}
//	log.Logger().Infof("Loaded the following config: %v", bsc)
//	return bsc, nil
//}

//- operating_system: darwin
// name: brew
// options:
//  - --cask
// packages:
//  - spotify
//  - goland
//- operating_system: darwin
// package_manager:
// name: custom
// command: /bin/bash -c
// packages:
//  - curl my-bash.sh | sh
//  - echo "hello world"
