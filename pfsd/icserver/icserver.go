package icserver

import (
	"bytes"
	"encoding/json"
	"github.com/cpssd/paranoid/pfsd/globals"
	"log"
	"net"
	"os"
	"path"
	"strconv"
	"time"
)

var (
	// MessageChan is the channel to which incoming messages will be passed
	// Attach a listener to this channel to receive messages
	MessageChan = make(chan FileSystemMessage, 100)
	verbose     = false

	listener *net.UnixListener
)

// FileSystemMessage is the structure which represents messages coming from the client
type FileSystemMessage struct {
	Command string   `json:"command"`
	Args    []string `json:"args"`
	Data    []byte   `json:"data"`
}

// handleConnection accepts a connection and handles messages received through the connection
func handleConnection(conn net.Conn) {
	verboseLog("icserver new connection")
	defer verboseLog("icserver connection lost")

	var messageBuffer bytes.Buffer

	for {
		buffer := make([]byte, 1024)
		endOfMessage := true
		mSize, err := conn.Read(buffer)
		if err != nil {
			// connection closed
			break
		}
		data := buffer[0:mSize]
		verboseLog("icserver new message:\n" + string(data) + "\nLength: " + strconv.Itoa(len(data)))
		messageBuffer.Write(data)
		message := &FileSystemMessage{}
		if string(data[len(data)-1]) != "}" {
			endOfMessage = false
		}

		if endOfMessage {
			fullMessageData := messageBuffer.Bytes()
			messageBuffer.Reset()
			err = json.Unmarshal(fullMessageData, message)
			if err != nil {
				if err.Error() == "unexpected end of JSON input" {
					endOfMessage = false
					messageBuffer.Write(fullMessageData)
				} else {
					log.Fatalln("icserver json unmarshal error: ", err)
				}
			}
			if endOfMessage {
				MessageChan <- *message
				verboseLog("icserver new message pushed to channel: " + message.Command)
			}
		}
	}
}

// StopAccept tells the server to stop accepting connections.
// Used when terminating PFSD.
func StopAccept() {
	listener.SetDeadline(time.Now())
}

// RunServer runs the server
// give a true parameter for verbose logging
func RunServer(pfsDirectory string, verboseLogging bool) {
	defer globals.Wait.Done()
	sockFilePath := path.Join(pfsDirectory, "meta", "pfic.sock")
	deleteSockFile(sockFilePath)
	verbose = verboseLogging

	tmplis, err := net.Listen("unix", sockFilePath)
	// Need to assert to a UnixListener to access the SetDeadline method
	listener = tmplis.(*net.UnixListener)
	if err != nil {
		log.Fatalln("ic listen error: ", err)
	}

	defer listener.Close()
	defer os.Remove(sockFilePath)

	verboseLog("icserver listening on " + sockFilePath)
	for {
		conn, err := listener.Accept()
		if err != nil {
			if nerr, ok := err.(net.Error); ok && nerr.Timeout() {
				log.Println("INFO: IC server has stopped accepting connections.")
				return
			}
			log.Println("ERROR: IC accept:", err)
		}
		go handleConnection(conn)
	}
}

// deleteSockFIle deletes the .sock file if it already exists.
// if one exists already the server cannot start
func deleteSockFile(filepath string) {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return
	}
	verboseLog("trailing .sock file detected")

	err := os.Remove(filepath)
	if err != nil {
		log.Fatalln("icserver delete sock file error: ", err)
	}
	verboseLog("trailing .sock file deleted")
}

// verboseLog logs what the server is doing if the verboseLogging option was
// given when running the server
func verboseLog(message string) {
	if verbose {
		log.Println(message)
	}
}