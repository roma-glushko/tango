package command

import (
	"fmt"
	"os"
	"tango/internal/di"

	"github.com/urfave/cli"
)

// InstallGeoLibCommand installs the geo lib
func InstallGeoLibCommand(cliContext *cli.Context) error {
	needUpdate := cliContext.Bool("update")
	geoLibResolver := di.InitMaxmindGeoLibResolver()

	fmt.Println("💃 Tango is on the scene!")

	geoLibPath, err := geoLibResolver.GetPath()

	if !os.IsNotExist(err) && !needUpdate {
		fmt.Println("🎉 geo lib has already been installed")
		fmt.Println("🛠 geo lib path: ", geoLibPath)
		return nil
	}

	installMaxmindLibUsecase := di.InitInstallMaxmindLibUsecase()

	if needUpdate {
		fmt.Println("💃 updating geo lib...")
	} else {
		fmt.Println("💃 installing geo lib...")
	}

	installMaxmindLibUsecase.Install()

	if needUpdate {
		fmt.Println("🎉 geo lib has been updated")
	} else {
		fmt.Println("🎉 geo lib has been installed")
	}

	fmt.Println("🛠 geo lib path: ", geoLibPath)

	return nil
}
