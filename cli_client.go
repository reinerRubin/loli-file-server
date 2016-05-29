package loli

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/codegangsta/cli"
	"github.com/reinerRubin/loli-file-server/grpc_client"
	"google.golang.org/grpc"
)

// FilerCli TBD
type FilerCli struct {
	context *cli.Context
}

// NewFilerCli TBD
func NewFilerCli(c *cli.Context) *FilerCli {
	return &FilerCli{context: c}
}

// PrintFileList TBD
func (l *FilerCli) PrintFileList() error {
	return l.withConnection(func(conn *grpc.ClientConn) error {
		filerClient, err := grpcclient.NewFilerClient(conn)
		if err != nil {
			log.Fatalf("err %+v", err)
			return err
		}

		fileList, err := filerClient.FileList()
		if err != nil {
			log.Fatalf("cmd error: %v", err)
			return err
		}
		fmt.Print(fileList.PrettyStr())

		return nil
	})
}

// SaveFiles TBD
func (l *FilerCli) SaveFiles() error {
	return l.withConnection(func(conn *grpc.ClientConn) error {
		args, err := parseCopyCmdArguments(l.context)
		if err != nil {
			log.Fatalf("err %+v", err)
		}

		filerClient, err := grpcclient.NewFilerClient(conn)
		if err != nil {
			log.Fatalf("err %+v", err)
		}

		for _, url := range args.URLs {
			target := args.target
			if args.isTargetDir {
				target = path.Join(target, path.Base(url))
			}
			file, err := os.Create(target)
			if err != nil {
				log.Fatalf("can't open target file %s", err)
				return err
			}
			defer file.Close()

			err = filerClient.GetFile(url, file)
			if err != nil {
				log.Fatalf("get file error: %s", err)
			}
		}

		return nil
	})

}

func (l *FilerCli) withConnection(fn func(conn *grpc.ClientConn) error) error {
	address := l.context.Args().First()
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can't connect; %v", err)
	}
	defer conn.Close()

	return fn(conn)
}

type cpCmdArgs struct {
	address     string
	URLs        []string
	target      string
	isTargetDir bool
}

func parseCopyCmdArguments(c *cli.Context) (*cpCmdArgs, error) {
	if c.NArg() < 3 {
		return nil, errors.New("not enough args")
	}

	args := c.Args()
	address := args.First()

	urls := args[1 : len(args)-1]
	target := args[len(args)-1]

	cpCmdArgs := &cpCmdArgs{
		address: address,
		URLs:    urls,
		target:  target,
	}

	// GNU coreutils cp behavior
	fileInfo, err := os.Stat(target)
	if err != nil && os.IsExist(err) {
		return nil, err
	}

	if len(urls) > 1 {
		if err != nil {
			return nil, err
		}
		if !fileInfo.IsDir() {
			return nil, errors.New("target is not directory")
		}
		cpCmdArgs.isTargetDir = true
	} else {
		if err == nil {
			cpCmdArgs.isTargetDir = fileInfo.IsDir()
		}

		if os.IsNotExist(err) {
			dir, _ := path.Split(target)
			_, err = os.Stat(dir)
			if err != nil {
				return nil, err
			}
		} else {
			cpCmdArgs.isTargetDir = true
		}
	}

	return cpCmdArgs, nil
}
