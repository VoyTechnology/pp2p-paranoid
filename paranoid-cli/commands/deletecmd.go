package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"os/user"
	"path"
	"path/filepath"
)

//Delete deletes a paranoid file system
func Delete(c *cli.Context) {
	args := c.Args()
	if len(args) < 1 {
		cli.ShowCommandHelp(c, "delete")
		os.Exit(1)
	}

	usr, err := user.Current()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pfspath, err := filepath.Abs(path.Join(usr.HomeDir, ".pfs", args[0]))
	if err != nil {
		fmt.Println("Given pfs-name is in incorrect format. Error : ", err)
		os.Exit(1)
	}
	if path.Base(pfspath) != args[0] {
		fmt.Println("Given pfs-name is in incorrect format")
		os.Exit(1)
	}

	err = os.RemoveAll(path.Join(usr.HomeDir, ".pfs", args[0]))
	if err != nil {
		fmt.Println("Could not delete given paranoid file system. Error :", err)
		os.Exit(1)
	}
}
