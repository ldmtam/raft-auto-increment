package auto_increment

import (
	"context"

	pb "github.com/ldmtam/raft-auto-increment/auto_increment/pb"
)

// GetOne ...
func (ai *AutoIncrement) GetOne(ctx context.Context, req *pb.GetOneRequest) (*pb.GetOneResponse, error) {
	value, err := ai.store.GetOne(req.Key)
	if err != nil {
		return nil, err
	}

	return &pb.GetOneResponse{Key: req.Key, Value: value}, nil
}

// GetMany ...
func (ai *AutoIncrement) GetMany(ctx context.Context, req *pb.GetManyRequest) (*pb.GetManyResponse, error) {
	from, to, err := ai.store.GetMany(req.Key, req.Quantity)
	if err != nil {
		return nil, err
	}

	return &pb.GetManyResponse{Key: req.Key, From: from, To: to}, nil
}

// GetLastInserted ...
func (ai *AutoIncrement) GetLastInserted(ctx context.Context, req *pb.GetLastInsertedRequest) (*pb.GetLastInsertedResponse, error) {
	value, err := ai.store.GetLastInserted(req.Key)
	if err != nil {
		return nil, err
	}

	return &pb.GetLastInsertedResponse{Key: req.Key, Value: value}, nil
}

// Join ...
func (ai *AutoIncrement) Join(ctx context.Context, req *pb.JoinRequest) (*pb.JoinResponse, error) {
	err := ai.store.Join(req.NodeID, req.NodeAddress)
	if err != nil {
		return nil, err
	}

	return &pb.JoinResponse{}, nil
}
