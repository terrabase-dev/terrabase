package models

import (
	"time"

	"google.golang.org/protobuf/proto"
)

type TerrabaseModel[P proto.Message] interface {
	SetUpdatedAt(time time.Time)
	ToProto() P
	ModelName() string
}
