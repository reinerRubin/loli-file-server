package convertor

import (
	"github.com/reinerRubin/loli-file-server/dto"
	"github.com/reinerRubin/loli-file-server/pb"
)

// DTOFileFromPb TBD
func DTOFileFromPb(pbFile *pb.File) *dto.File {
	return &dto.File{
		URL: pbFile.Url,
	}
}

// DTOFileListFromPb TBD
func DTOFileListFromPb(pbFiles []*pb.File) dto.FileList {
	dtoFileList := make(dto.FileList, len(pbFiles))
	for i := range pbFiles {
		dtoFileList[i] = DTOFileFromPb(pbFiles[i])
	}
	return dtoFileList
}

// PbFileFromDto TBD
func PbFileFromDto(dtoFile *dto.File) *pb.File {
	return &pb.File{
		Url: dtoFile.URL,
	}
}

// PbFileListFromDto TBD
func PbFileListFromDto(dtoFiles []*dto.File) []*pb.File {
	pbFileList := make([]*pb.File, len(dtoFiles))
	for i := range dtoFiles {
		pbFileList[i] = PbFileFromDto(dtoFiles[i])
	}
	return pbFileList
}
