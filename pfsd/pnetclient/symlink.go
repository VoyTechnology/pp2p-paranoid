package pnetclient

import (
	"github.com/cpssd/paranoid/pfsd/globals"
	pb "github.com/cpssd/paranoid/proto/paranoidnetwork"
	"golang.org/x/net/context"
)

func Symlink(oldPath, newPath string) {
	nodes := globals.Nodes.GetAll()
	for _, node := range nodes {
		conn, err := Dial(node)
		if err != nil {
			Log.Error("Symlink: failed to dial ", node)
			continue
		}
		defer conn.Close()

		client := pb.NewParanoidNetworkClient(conn)

		_, err = client.Symlink(context.Background(), &pb.LinkRequest{oldPath, newPath})
		if err != nil {
			Log.Error("Failed sending symkink to", node, "Error:", err)
		}
	}
}