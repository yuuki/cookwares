package aptpackage

import (
	"fmt"

	"github.com/yuuki/cookwares/pkg/command"
)

func Install(name string, version string) error {
	fullName := ""
	if version != "" {
		fullName = fmt.Sprintf("%s=%s", name, version)
	} else {
		fullName = name
	}
	cmd := fmt.Sprintf("DEBIAN_FRONTEND='noninteractive' apt-get -y -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confold' install %s", fullName)
	return command.Run(cmd)
}
