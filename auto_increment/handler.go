package auto_increment

import (
	"context"

	pb "github.com/ldmtam/raft-auto-increment/auto_increment/pb"
)

func (s *autoIncrement) GetSingle(ctx context.Context, req *pb.GetSingleRequest) (*pb.GetSingleResponse, error) {
	return &pb.GetSingleResponse{Key: "test", Value: 1}, nil
}

func (s *autoIncrement) GetMultiple(ctx context.Context, req *pb.GetMultipleRequest) (*pb.GetMultipleResponse, error) {
	return &pb.GetMultipleResponse{Key: "test", Value: []uint64{1, 2, 3}}, nil
}
