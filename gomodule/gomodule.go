package gomodule

import (
	"fmt"
	"os"
	"strings"
)

const pathLength = 3

type GoModule struct {
	vcs     string
	owner   string
	name    string
	version string
}

func Parse(path string, version string) (GoModule, error) {
	split := strings.Split(path, string(os.PathSeparator))

	if len(split) != pathLength {
		return GoModule{}, fmt.Errorf("Path supplied is incorrect number of tokens")
	}

	module := GoModule{
		vcs:     split[0],
		owner:   split[1],
		name:    split[2],
		version: version,
	}

	return module, nil
}

func (goModule GoModule) GetVcs() string {
	return goModule.vcs
}

func (goModule GoModule) GetModuleZipName() string {
	return goModule.name + "@" + goModule.version + ".zip"
}

func (goModule GoModule) GetProjectPath() string {
	return fmt.Sprintf("%s/%s/%s", goModule.vcs, goModule.owner, goModule.name)
}
