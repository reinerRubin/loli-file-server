package grpcclient

import (
	"fmt"
	"io"

	"google.golang.org/grpc"

	"github.com/reinerRubin/loli-file-server/convertor"
	"github.com/reinerRubin/loli-file-server/dto"
	"github.com/reinerRubin/loli-file-server/pb"
	"golang.org/x/net/context"
)

// NewFilerClient TBD
func NewFilerClient(conn *grpc.ClientConn) (FilerClient, error) {
	return FilerClient{grpcClient: pb.NewFilerClient(conn)}, nil
}

// FilerClient TBD
type FilerClient struct {
	grpcClient pb.FilerClient
}

// FileList TBD
func (f *FilerClient) FileList() (dto.FileList, error) {
	protoFileListResponse, err := f.grpcClient.FileList(context.Background(), &pb.Empty{})
	if err != nil {
		return nil, err
	}

	return convertor.DTOFileListFromPb(protoFileListResponse.Files), nil
}

// GetFile TBDщ
func (f *FilerClient) GetFile(url string, w io.Writer) error {
	s, err := f.grpcClient.GetFileBody(context.Background(), &pb.GetFileRequest{Url: url})
	if err != nil {
		return fmt.Errorf("can't call: %s", err)
	}

	for {
		filePart, err := s.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		w.Write(filePart.File)
	}
	return nil
}
