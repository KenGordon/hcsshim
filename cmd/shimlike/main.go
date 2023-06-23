package main

import (
	"os"

	"github.com/Microsoft/go-winio"
	"github.com/Microsoft/hcsshim/cmd/shimlike/proto"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
)

var (
	usage = "shimlike <pipe address> <UVM ID>"
)

func run(cCtx cli.Context) {
	if cCtx.NArg() != 2 {
		logrus.Fatalf("Usage: %s", usage)
	}
	s := grpc.NewServer()
	pipe, err := winio.ListenPipe(cCtx.Args().First(), nil)
	if err != nil {
		logrus.Fatal(err)
	}
	defer pipe.Close()
	rs := RuntimeServer{VMID: cCtx.Args().Get(2)}
	proto.RegisterRuntimeServiceServer(s, &rs)

	// Connect to the UVM
	logrus.Info("Connecting to UVM")
	rs.connectLog()
	logrus.Info("Accepting GCS")
	rs.acceptGcs()

	// Create the gRPC server and listen on the pipe.
	// This blocks until the pipe is closed or Stop() is called.
	err = s.Serve(pipe)
	if err != nil {
		logrus.Fatal(err)
	}
}

func main() {
	app := cli.App{
		Name:      "shimlike",
		Usage:     "Connect to a UVM",
		ArgsUsage: usage,
		Action:    run,
	}
	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}