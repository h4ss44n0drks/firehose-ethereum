package substreams

import (
	"github.com/streamingfast/logging"
	"go.uber.org/zap"
)

var zlog *zap.Logger

func init() {
	zlog, _ = logging.PackageLogger("rpc-cache", "github.com/firehose-ethereum/substreams")
}
