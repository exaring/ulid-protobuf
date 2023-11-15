package ulidpb

import (
	"encoding/binary"

	"github.com/oklog/ulid/v2"
)

//go:generate protoc -I=. --go_out=paths=source_relative:. ./ulid.proto

func (x *ULID) AsULID() ulid.ULID {
	var u2 [16]byte

	binary.LittleEndian.PutUint64(u2[0:8], x.Low)
	binary.LittleEndian.PutUint64(u2[8:16], x.High)

	return ulid.ULID(u2)
}

func FromULID(u ulid.ULID) *ULID {
	ua := [16]byte(u)

	low := binary.LittleEndian.Uint64(ua[0:8])
	high := binary.LittleEndian.Uint64(ua[8:16])

	return &ULID{
		Low:  low,
		High: high,
	}
}
