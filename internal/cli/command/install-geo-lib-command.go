package command

import (
	"fmt"
	"tango/internal/di"

	"github.com/urfave/cli"
)

// InstallGeoLibCommand installs the geo lib
func InstallGeoLibCommand(cliContext *cli.Context) error {
	installMaxmindLibUsecase := di.InitInstallMaxmindLibUsecase()

	fmt.Println("💃 Tango is on the scene!")
	fmt.Println("💃 installing geo lib...")

	installMaxmindLibUsecase.Install()

	fmt.Println("🎉 geo lib has been installed")

	return nil
}
