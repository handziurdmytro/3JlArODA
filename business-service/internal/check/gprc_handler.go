package check

import checkpb "github.com/handziurdmytro/3JlArODA/business-service/pb/business/check"

type GRPCHandler struct {
	checkpb.UnimplementedCheckServiceServer
}
