package auto_increment

import (
	"context"

	pb "github.com/ldmtam/raft-auto-increment/auto_increment/pb"
)

func (s *autoIncrement) GetSingle(ctx context.Context, req *pb.GetSingleRequest) (*pb.GetSingleResponse, error) {
	value, err := s.db.GetSingle(req.Key)
	if err != nil {
		return nil, err
	}

	return &pb.GetSingleResponse{Key: req.Key, Value: value}, nil
}

func (s *autoIncrement) GetMultiple(ctx context.Context, req *pb.GetMultipleRequest) (*pb.GetMultipleResponse, error) {
	values, err := s.db.GetMultiple(req.Key, req.Quantity)
	if err != nil {
		return nil, err
	}

	return &pb.GetMultipleResponse{Key: req.Key, Values: values}, nil
}

func (s *autoIncrement) GetLast(ctx context.Context, req *pb.GetLastRequest) (*pb.GetLastResponse, error) {
	value, err := s.db.GetLast(req.Key)
	if err != nil {
		return nil, err
	}

	return &pb.GetLastResponse{Value: value}, nil
}
