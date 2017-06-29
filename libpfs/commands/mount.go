package commands

import (
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"

	"github.com/pp2p/paranoid/libpfs/returncodes"
)

//MountCommand is used to notify a pfs paranoidDirectory it has been mounted.
func MountCommand(paranoidDirectory, dAddr, mountPoint string) (returnCode returncodes.Code, returnError error) {
	Log.Info("mount command called")
	Log.Verbose("mount : given paranoidDirectory = " + paranoidDirectory)

	err := ioutil.WriteFile(path.Join(paranoidDirectory, "meta", "discovery_address"), []byte(dAddr), 0600)
	if err != nil {
		return returncodes.EUNEXPECTED, fmt.Errorf("error saving discovery_address file: %s", err)
	}

	mountPoint, err = filepath.Abs(mountPoint)
	if err != nil {
		return returncodes.EUNEXPECTED, fmt.Errorf("error getting absolute path of mountpoint: %s", err)
	}

	err = ioutil.WriteFile(path.Join(paranoidDirectory, "meta", "mountpoint"), []byte(mountPoint), 0600)
	if err != nil {
		return returncodes.EUNEXPECTED, fmt.Errorf("error writing mountpoint: %s", err)
	}

	return returncodes.OK, nil
}