package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/cpssd/paranoid/paranoid-cli/tls"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path"
	"path/filepath"
)

//Init inits a new paranoid file system
func Init(c *cli.Context) {
	args := c.Args()
	if len(args) < 1 {
		cli.ShowCommandHelp(c, "init")
		os.Exit(0)
	}

	pfsname := args[0]

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	homeDir := usr.HomeDir

	if _, err := os.Stat(path.Join(homeDir, ".pfs")); os.IsNotExist(err) {
		err = os.Mkdir(path.Join(homeDir, ".pfs"), 0700)
		if err != nil {
			log.Fatalln("FATAL : Error making pfs directory")
		}
	}

	directory, err := filepath.Abs(path.Join(homeDir, ".pfs", pfsname))
	if err != nil {
		log.Fatalln("Given pfs-name is in incorrect format. Error : ", err)
	}
	if path.Base(directory) != args[0] {
		log.Fatalln("Given pfs-name is in incorrect format.")
	}

	if _, err := os.Stat(directory); !os.IsNotExist(err) {
		log.Fatalln("FATAL : a paranoid file system with that name already exists")
	}
	err = os.Mkdir(directory, 0700)
	if err != nil {
		log.Fatalln("FATAL : Error making pfs directory, error : ", err)
	}

	cmd := exec.Command("pfsm", "init", directory)
	err = cmd.Run()
	if err != nil {
		log.Fatalln("FATAL : ", err)
	}

	if c.Bool("unsecure") {
		log.Println("--unsecure specified. PFSD will not use TLS for its communication.")
		return
	}
	if (c.String("cert") != "") && (c.String("key") != "") {
		log.Println("INFO: Using existing certificate.")
		err = os.Link(c.String("cert"), path.Join(directory, "meta", "cert.pem"))
		if err != nil {
			log.Fatalln("FATAL: Failed to copy cert file:", err)
		}
		err = os.Link(c.String("key"), path.Join(directory, "meta", "key.pem"))
		if err != nil {
			log.Fatalln("FATAL: Failed to copy key file:", err)
		}
	} else {
		log.Println("INFO: Generating certificate.")
		fmt.Println("Generating TLS certificate. Please follow the given instructions.")
		tls.GenCertificate(directory)
	}
}
