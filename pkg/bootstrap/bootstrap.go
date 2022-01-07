package bootstrap

import (
	_ "encoding/json"
	"github.com/Benbentwo/SystemInfo/pkg/bootstrap/package_managers"
	"github.com/Benbentwo/SystemInfo/pkg/common/utils"
	"gopkg.in/yaml.v2"
	"sort"

	"errors"
	"github.com/Benbentwo/SystemInfo/pkg/common/log"
	"github.com/Benbentwo/utils/util"
	"io/ioutil"
)

type Profile struct {
	ProfileName string
	Installers  []Installer
}
type Installer struct {
	OperatingSystem string   `yaml:"operating_system,omitempty"`
	Name            string   `yaml:"name,omitempty"`
	Options         []string `yaml:"options,omitempty"`
	Packages        []string `yaml:"packages,omitempty"`
	PackageManager  string   `yaml:"package_manager,omitempty"`
	Command         string   `yaml:"command,omitempty"`
}

func LoadBootstrapConfig(inputFile string) ([]Profile, error) {

	exists, err := util.FileExists(inputFile)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("File inputted does not exist or could not be found")
	}

	filebytes, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return nil, err
	}

	data := map[string][]Installer{}
	err = yaml.Unmarshal(filebytes, &data)
	if err != nil {
		log.Logger().Errorf("Unmarshaling yaml failed: %s", err)
		log.Logger().Debugf("%v", data)
		return nil, err
	}

	itemsMap := data
	bootstrapConfig := make([]Profile, 0)
	for k, v := range itemsMap {
		bootstrapConfig = append(bootstrapConfig, Profile{
			ProfileName: k,
			Installers:  v,
		})
		log.Logger().Infof("Successfully Loaded Profile: %s", utils.ColorInfo(k))
	}
	sort.Slice(bootstrapConfig, func(i, j int) bool {
		return bootstrapConfig[i].ProfileName < bootstrapConfig[j].ProfileName
	})
	return bootstrapConfig, nil
}

func (i Installer) GetInstaller() (package_managers.PackageManager, error) {
	return package_managers.GetInstaller(i.Name)
}
