package filetransfer

import (
	"os"
	"path/filepath"

	"github.com/kr/fs"
	pb "github.com/navinds25/styx/pkg/styxpb"
	"google.golang.org/grpc"
)

// ListDirSender is the interface for passing object to stream RemoteFile response
type ListDirSender interface {
	Send(*pb.RemoteFile) error
	grpc.ServerStream
}

// ListDir implements grpc server function for listing a directory recursively and returning the files as a stream.
func (s *FTServer) ListDir(in *pb.RemoteDirectory, stream pb.FTService_ListDirServer) error {
	listDirWalk(in.SourcePath, stream)
	return nil
}

// ListDirCondition lists a directory recursively and returns the files as a stream after evaluating a condition.
func (s *FTServer) ListDirCondition(in *pb.RemoteDirectoryCondition, stream pb.FT_ListDirConditionServer) error {
	listDirWalkCondition(in.SourcePath, stream, in.GetCondition())
	return nil
}

func checkCondition(conditionList []*pb.Condition, stat os.FileInfo) bool {
	// default state is true, so an empty condition check will pass.
	state := true
	for i := range conditionList {
		c := conditionList[i]
		switch c.GetConditionType() {
		case pb.Condition_TimeRange:
			trcondition := func(c *pb.Condition, stat os.FileInfo) bool {
				modTime := stat.ModTime().Unix()
				startTime := c.GetTimeValues().GetTimeStart()
				endTime := c.GetTimeValues().GetTimeEnd()
				if modTime > startTime && modTime < endTime {
					return true
				}
				return false
			}
			state = state && trcondition(c, stat)
		case pb.Condition_Time:
			tcondition := func(c *pb.Condition, stat os.FileInfo) bool {
				modTime := stat.ModTime().Unix()
				startTime := c.GetTimeValues().GetTimeStart()
				if modTime > startTime {
					return true
				}
				return false
			}
			state = state && tcondition(c, stat)
		case pb.Condition_Glob:
			gcondition := func(c *pb.Condition, stat os.FileInfo) bool {
				return false
			}
			state = state && gcondition(c, stat)
		case pb.Condition_Regex:
			rcondition := func(c *pb.Condition, stat os.FileInfo) bool {
				return false
			}
			state = state && rcondition(c, stat)
		}
	}
	return state
}

func listDirWalk(sourcePath string, stream ListDirSender) error {
	walker := fs.Walk(sourcePath)
	for walker.Step() {
		if err := walker.Err(); err != nil {
			continue
		}
		stat := walker.Stat()
		if stat.IsDir() {
			continue
		} else {
			if err := stream.Send(&pb.RemoteFile{
				Sourcefile: filepath.Join(walker.Path(), stat.Name()),
			}); err != nil {
				return err
			}
		}
	}
	return nil
}

func listDirWalkCondition(sourcePath string, stream ListDirSender, conditionList []*pb.Condition) error {
	walker := fs.Walk(sourcePath)
	for walker.Step() {
		if err := walker.Err(); err != nil {
			continue
		}
		stat := walker.Stat()
		if stat.IsDir() {
			continue
		} else if checkCondition(conditionList, stat) {
			if err := stream.Send(&pb.RemoteFile{
				Sourcefile: filepath.Join(walker.Path(), stat.Name()),
			}); err != nil {
				return err
			}
		}
	}
	return nil
}
