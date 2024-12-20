package partition

import (
	"github.com/isi-lincoln/go-diskfs/partition/part"
	"github.com/isi-lincoln/go-diskfs/util"
)

// Table reference to a partitioning table on disk
type Table interface {
	Type() string
	Write(util.File, int64) error
	GetPartitions() []part.Partition
	Repair(diskSize uint64) error
	Verify(f util.File, diskSize uint64) error
	UUID() string
}
