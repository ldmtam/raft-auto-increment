package auto_increment

import (
	"context"

	pb "github.com/ldmtam/raft-auto-increment/auto_increment/pb"
)

// GetSingle ...
func (ai *AutoIncrement) GetSingle(ctx context.Context, req *pb.GetSingleRequest) (*pb.GetSingleResponse, error) {
	value, err := ai.store.GetSingle(req.Key)
	if err != nil {
		return nil, err
	}

	return &pb.GetSingleResponse{Key: req.Key, Value: value}, nil
}

// GetMultiple ...
func (ai *AutoIncrement) GetMultiple(ctx context.Context, req *pb.GetMultipleRequest) (*pb.GetMultipleResponse, error) {
	values, err := ai.store.GetMultiple(req.Key, req.Quantity)
	if err != nil {
		return nil, err
	}

	return &pb.GetMultipleResponse{Key: req.Key, Values: values}, nil
}

// GetLast ...
func (ai *AutoIncrement) GetLast(ctx context.Context, req *pb.GetLastRequest) (*pb.GetLastResponse, error) {
	value, err := ai.store.GetLast(req.Key)
	if err != nil {
		return nil, err
	}

	return &pb.GetLastResponse{Value: value}, nil
}

// Join ...
func (ai *AutoIncrement) Join(ctx context.Context, req *pb.JoinRequest) (*pb.JoinResponse, error) {
	err := ai.store.Join(req.NodeID, req.NodeAddress)
	if err != nil {
		return nil, err
	}

	return &pb.JoinResponse{}, nil
}
