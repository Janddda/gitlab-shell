package handler

import (
	"context"
	"fmt"
	"os"

	pb "gitlab.com/gitlab-org/gitaly-proto/go"
	"gitlab.com/gitlab-org/gitaly/client"
)

func UploadArchive(gitalyAddress string, request *pb.SSHUploadArchiveRequest) (int32, error) {
	if gitalyAddress == "" {
		return 0, fmt.Errorf("no gitaly_address given")
	}

	conn, err := client.Dial(gitalyAddress, dialOpts())
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return client.UploadArchive(ctx, conn, os.Stdin, os.Stdout, os.Stderr, request)
}
