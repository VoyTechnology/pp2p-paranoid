package commands

import (
	"github.com/codegangsta/cli"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path"
)

//Unmounts a paranoid file system
func Unmount(c *cli.Context) {
	args := c.Args()
	if len(args) < 1 {
		cli.ShowAppHelp(c)
		os.Exit(0)
	}

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	mountpoint, err := ioutil.ReadFile(path.Join(usr.HomeDir, "pfs", args[0], "meta", "mountpoint"))
	if err != nil {
		log.Fatalln("FATAL : Could not get mountpoint ", err)
	}

	cmd := exec.Command("fusermount", "-u", "-z", string(mountpoint))
	err = cmd.Run()
	if err != nil {
		log.Fatalln("FATAL : unmount failed ", err)
	}
}
