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

	homeDirResolver := di.InitHomeDirResolver()
	geoLibResolver := di.InitMaxmindGeoLibResolver()
	geoConfResolver := di.InitMaxmindConfResolver()
	generateConfUsecase := di.InitGenerateMaxmindConfUsecase()

	fmt.Println("💃 Tango is on the scene!")

	geoConfPath, geoConfExistErr := geoConfResolver.GetPath()

	if os.IsNotExist(geoConfExistErr) {
		// generate maxmind conf file
		fmt.Println("🛠 Generating MaxMind Config file..")

		accountID := cliContext.String("account-id")
		licenseKey := cliContext.String("license-key")

		if accountID == "" || licenseKey == "" {
			fmt.Println("")
			fmt.Println("🚨 MaxMind Config file should be generated to work with Geo data")
			fmt.Println("🚨 You need to provide 'account-id' and 'license-key' arguments for this command")
			fmt.Println("🚨 Credentials can be found under https://www.maxmind.com/en/accounts/current/license-key page")

			return nil
		}

		err := generateConfUsecase.Generate(geoConfPath, accountID, licenseKey)

		if err != nil {
			fmt.Printf("🚨 Failed to generate MaxMind Config file: ", err.Error())

			return nil
		}

		fmt.Println("🛠 MaxMind Config file has been generated: ", geoConfPath)
		fmt.Println("🛠 You can customize the config file if needed (https://github.com/maxmind/geoipupdate/blob/master/doc/GeoIP.conf.md)")
	}

	fmt.Println("🛠 geo conf file: ", geoConfPath)

	geoLibPath, geoLibExistErr := geoLibResolver.GetPath()

	if !os.IsNotExist(geoLibExistErr) && !needUpdate {
		fmt.Println("🎉 geo lib has already been installed")
		fmt.Println("🛠 geo lib path: ", geoLibPath)
		return nil
	}

	if needUpdate {
		fmt.Println("💃 updating geo lib...")
	} else {
		fmt.Println("💃 installing geo lib...")
	}

	homeDir := homeDirResolver.GetPath()
	installMaxmindLibUsecase := di.InitInstallMaxmindLibUsecase()
	installMaxmindLibUsecase.Install(geoConfPath, homeDir, true)

	if needUpdate {
		fmt.Println("🎉 geo lib has been updated")
	} else {
		fmt.Println("🎉 geo lib has been installed")
	}

	fmt.Println("🛠 geo lib path: ", geoLibPath)

	return nil
}
