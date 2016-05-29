package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/codegangsta/cli"
	"github.com/reinerRubin/loli-file-server/grpc_services"
	"github.com/reinerRubin/loli-file-server/pb"
	"google.golang.org/grpc"
)

var appFlags = []cli.Flag{
	cli.IntFlag{
		Name:  "p",
		Usage: "listened port",
		Value: 50051,
	},
}

func main() {
	app := cli.NewApp()
	app.Name = "Loli-server server"
	app.Flags = appFlags
	app.UsageText = "./server /path/to/dir"

	app.Action = func(c *cli.Context) error {
		if len(c.Args()) < 1 {
			log.Fatalf("enter directory")
		}

		log.Printf("start listening %d", c.Int("p"))
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", c.Int("p")))
		if err != nil {
			log.Fatalf("failed to listen: %s", err)
		}
		s := grpc.NewServer()

		log.Println("start file dir reading")

		filer, err := grpcservices.NewFiler(c.Args().First())
		if err != nil {
			log.Fatalf("failed init server: %s", err)
		}

		log.Printf("file cnt %d", filer.FilesCnt())
		log.Println("start server")

		pb.RegisterFilerServer(s, filer)
		s.Serve(lis)

		return nil
	}

	app.Run(os.Args)
}
