package benchmarks_test

import (
	"testing"

	ulidpb "github.com/exaring/ulid-protobuf"
	"github.com/oklog/ulid/v2"
	"google.golang.org/protobuf/proto"
)

//go:generate protoc -I=. -I=../ --go_out=. --go_opt=paths=source_relative benchmarks.proto

func BenchmarkString(b *testing.B) {
	x := ulid.Make()

	tmp := &String{}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tmp.Id = x.String()
		outRaw, err := proto.Marshal(tmp)
		if err != nil {
			b.Errorf("failed to marshal: %v", err)
		}
		if err = proto.Unmarshal(outRaw, tmp); err != nil {
			b.Errorf("failed to unmarshal: %v", err)
		}
		_ = ulid.MustParse(tmp.Id)
	}
}

func BenchmarkNative(b *testing.B) {
	x := ulid.Make()

	tmp := &Native{}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tmp.Id = ulidpb.FromULID(x)
		outRaw, err := proto.Marshal(tmp)
		if err != nil {
			b.Errorf("failed to marshal: %v", err)
		}
		if err = proto.Unmarshal(outRaw, tmp); err != nil {
			b.Errorf("failed to unmarshal: %v", err)
		}
		_ = tmp.Id.AsULID()
	}
}
