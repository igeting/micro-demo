package handler

import (
	"context"
	"micro-demo/file/proto/pb"
	"os"
)

type FileManager struct {
}

func (FileManager) Upload(ctx context.Context, req *pb.FileRequest, res *pb.FileResponse) error {
	err := os.WriteFile(req.Name, req.Data, os.ModePerm)
	if err != nil {
		res.Result = "upload file error"
		return err
	}
	res.Result = "upload file success"
	return nil
}
