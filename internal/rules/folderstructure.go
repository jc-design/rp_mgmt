package rules

import (
	"path/filepath"

	"github.com/kirsle/configdir"
)

type Folderstructure struct {
	Basepath   string
	Logfiles   string
	Settings   string
	Rules      string
	Data       string
	Characters string
}

func NewFolderstructure(appname string) (Folderstructure, error) {
	// Build the folder structure to store settings and data
	// A common use case is to get a private config folder for your app to
	// place its settings files into, that are specific to the local user.
	f := Folderstructure{}
	f.Basepath = configdir.LocalConfig(appname)
	err := configdir.MakePath(f.Basepath)
	if err != nil {
		return f, err
	}

	f.Logfiles = filepath.Join(f.Basepath, "logfiles")
	err = configdir.MakePath(f.Logfiles)
	if err != nil {
		return f, err
	}

	f.Settings = filepath.Join(f.Basepath, "settings")
	err = configdir.MakePath(f.Settings)
	if err != nil {
		return f, err
	}

	f.Rules = filepath.Join(f.Basepath, "rules")
	err = configdir.MakePath(f.Rules)
	if err != nil {
		return f, err
	}

	f.Data = filepath.Join(f.Basepath, "data")
	err = configdir.MakePath(f.Data)
	if err != nil {
		return f, err
	}

	f.Characters = filepath.Join(f.Basepath, "characters")
	err = configdir.MakePath(f.Characters)
	if err != nil {
		return f, err
	}

	return f, nil
}
