package uuidgen

import (
	"time"
)

/*
UUID bits composition

1 bit      41 bits 				5 bits 			 5 bits			   12 bits
| 0	    |  timestamp		| 	datacenter ID |  machine ID		|  sequence number |

*/

type UUID uint64

var sqN UUID

type MetaData struct {
	DataCenterID int
	HostID       int
}

const (
	ReservedBits                = 1
	TimestampMaxBits            = 41
	DatacenterIDMaxBits         = 5
	MachineIDMaxBits            = 5
	sequenceNoMaxBits           = 12
	MachineIDBitsOffset         = sequenceNoMaxBits
	DatacenterIDBitsOffset      = MachineIDBitsOffset + MachineIDMaxBits
	TimestampBitsOffset         = DatacenterIDBitsOffset + DatacenterIDMaxBits
	uuidMaxValue           UUID = 0x7FFFFFFFFFFFFFFF
	timeStampMaxValue           = 1<<TimestampMaxBits - 1
)

func GenerateUUID(metaData MetaData) UUID {

	ts := UUID(time.Now().UnixMilli() << TimestampBitsOffset)
	dcID := UUID(metaData.DataCenterID << DatacenterIDBitsOffset)
	hostID := UUID(metaData.HostID << MachineIDBitsOffset)
	sqN += sqN

	return (uuidMaxValue&(ts|dcID) | hostID | sqN)
}

// This should be called after every milisecond
func ResetSequence() {
	sqN = 0
}
