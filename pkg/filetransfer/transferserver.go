package filetransfer

import pb "github.com/navinds25/styx/api/filetransfer"

//func (s *FTServer) TransferFile(in *pb.RemoteDirectoryTransfer, stream pb.FT_TransferFileServer) error {
//	return nil
//}
//
//func (s *FTServer) TransferCondition(in *pb.RemoteDirectoryConditionTransfer, stream pb.FT_TransferConditionServer) error {
//	return nil
//}

func (s *RemoteFTServer) TransferFile(in *pb.RemoteDirectoryTransfer, stream pb.RemoteFTService_TransferFileServer) error {
	return nil
}

func (s *RemoteFTServer) TransferCondition(in *pb.RemoteDirectoryConditionTransfer, stream pb.RemoteFTService_TransferConditionServer) error {
	return nil
}
