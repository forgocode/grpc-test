package server

import (
	"context"
	"fmt"
	"testing"
	"time"

	"google.golang.org/grpc"
	pb "serverMonitor/internal/logService/proto"
)

func Test_SrartLogRpc(t *testing.T) {
	go func() {
		StartLogRpc()
	}()
	time.Sleep(1 * time.Second)
	conn, err := grpc.Dial("localhost:20008", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("can't connect to 20008")
	}
	defer conn.Close()

	client := pb.NewRecordLogClient(conn)
	for i := 0; i < 10; i++ {
		result, err := client.RecordLogMsg(context.Background(), &pb.Msg{Msg: "这是日志"})
		fmt.Println(111, result, err)
		time.Sleep(1 * time.Second)
	}
}
