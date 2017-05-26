// GENERATED FILE (simulated)

package quota

import "io"
import "github.com/golang/protobuf/proto"

// should we use protos for these configs?
//
// this itself could be an interface (perhaps empty) that individual
// adapters could implement in a config file.
type Config proto.Message

type Adapter interface {
	io.Closer

	NewProcessor(Config, []*Template) (Processor, error)
}

type Processor interface {
	io.Closer

	Process([]*Instance) error
}
