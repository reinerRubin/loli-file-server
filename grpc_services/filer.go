package grpcservices

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"

	"github.com/reinerRubin/loli-file-server/convertor"
	"github.com/reinerRubin/loli-file-server/dto"
	"github.com/reinerRubin/loli-file-server/pb"
	"golang.org/x/net/context"
)

const chunkSize = 2 * 1000000 // ~2mb

// Filer TBD
type Filer struct {
	files    dto.FileList
	rootPath string
}

// NewFiler TBD
func NewFiler(rootPath string) (*Filer, error) {
	var files dto.FileList
	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			relPath, err := filepath.Rel(rootPath, path)
			if err != nil {
				return err
			}
			files = append(files, &dto.File{URL: relPath})
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("can't scan directory: %s", err)
	}

	filer := &Filer{
		rootPath: rootPath,
		files:    files,
	}

	return filer, nil
}

// FileList TBD
func (f *Filer) FileList(ctx context.Context, in *pb.Empty) (*pb.FileListResponse, error) {
	return &pb.FileListResponse{
		Files: convertor.PbFileListFromDto(f.files),
	}, nil
}

func (f *Filer) filePresent(path string) bool {
	// TODO mkdir cache
	_, find := f.files.ToURLMap()[path]
	return find
}

// GetFileBody TBD
func (f *Filer) GetFileBody(r *pb.GetFileRequest, s pb.Filer_GetFileBodyServer) error {
	if !f.filePresent(r.Url) {
		return fmt.Errorf("can't find file: %s", r.Url)
	}
	filePath := path.Join(f.rootPath, r.Url)

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	data := make([]byte, chunkSize)

	for {
		data = data[:cap(data)]
		n, err := file.Read(data)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		data = data[:n]

		filePart := &pb.FilePartResponse{
			File: data,
		}
		if err := s.Send(filePart); err != nil {
			return err
		}

	}

	return nil
}

// FilesCnt TBD
func (f *Filer) FilesCnt() int {
	return len(f.files)
}
