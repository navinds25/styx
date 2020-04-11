package filetransfer

import pb "github.com/navinds25/styx/pkg/styxpb"

func (s *FTServer) TransferFile(in *pb.RemoteDirectoryTransfer, stream pb.FT_TransferFileServer) error {
	return nil
}

func (s *FTServer) TransferCondition(in *pb.RemoteDirectoryConditionTransfer, stream pb.FT_TransferConditionServer) error {
	return nil
}
