package log

import (
	"sync"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	api "github.com/abdulmajid18/log-service/api/v1"
)

type Replicator struct {
	DialOption  []grpc.DialOption
	LocalServer api.LogClient
	logger      *zap.Logger
	mu          sync.Mutex
	servers     map[string]chan struct{}
	closed      bool
	close       chan struct{}
}
