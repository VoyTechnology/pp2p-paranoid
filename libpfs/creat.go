package libpfs

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/pp2p/paranoid/libpfs/encryption"
	"github.com/pp2p/paranoid/libpfs/returncodes"
	log "github.com/pp2p/paranoid/logger"
)

//CreatCommand creates a new file with the name filePath in the pfs paranoidDirectory
func CreatCommand(paranoidDirectory, filePath string, perms os.FileMode) (returnCode returncodes.Code, returnError error) {
	log.V(1).Info("creat called on %s in %s", filePath, paranoidDirectory)

	err := GetFileSystemLock(paranoidDirectory, ExclusiveLock)
	if err != nil {
		return returncodes.EUNEXPECTED, err
	}

	defer func() {
		err := UnLockFileSystem(paranoidDirectory)
		if err != nil {
			returnCode = returncodes.EUNEXPECTED
			returnError = err
		}
	}()

	namepath := getParanoidPath(paranoidDirectory, filePath)

	fileType, err := getFileType(paranoidDirectory, namepath)
	if err != nil {
		return returncodes.EUNEXPECTED, err
	}

	if fileType != typeENOENT {
		return returncodes.EEXIST, errors.New(filePath + " already exists")
	}

	uuidbytes, err := generateNewInode()
	if err != nil {
		return returncodes.EUNEXPECTED, err
	}

	uuidstring := string(uuidbytes)

	err = ioutil.WriteFile(namepath, uuidbytes, 0600)
	if err != nil {
		return returncodes.EUNEXPECTED, errors.New("error writing name file")
	}

	nodeData := &inode{
		Mode:  perms,
		Inode: uuidstring,
		Count: 1}
	jsonData, err := json.Marshal(nodeData)
	if err != nil {
		return returncodes.EUNEXPECTED, fmt.Errorf("error marshalling json: %s", err)
	}

	err = ioutil.WriteFile(path.Join(paranoidDirectory, "inodes", uuidstring), jsonData, 0600)
	if err != nil {
		return returncodes.EUNEXPECTED, fmt.Errorf("error writing inodes file: %s", err)
	}

	contentsFile, err := os.Create(path.Join(paranoidDirectory, "contents", uuidstring))
	if err != nil {
		return returncodes.EUNEXPECTED, fmt.Errorf("error creating contents file: %s", err)
	}
	defer contentsFile.Close()

	if encryption.Encrypted {
		n, err := contentsFile.WriteAt([]byte{1}, 0)
		if err != nil {
			return returncodes.EUNEXPECTED, fmt.Errorf("error creating contents file: %s", err)
		}
		if n != 1 {
			return returncodes.EUNEXPECTED, errors.New("error writing first byte to contents file")
		}
	}

	return returncodes.OK, nil
}
