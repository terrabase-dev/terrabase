package models

import (
	"time"

	organizationv1 "github.com/terrabase-dev/terrabase/specs/terrabase/organization/v1"
	"github.com/uptrace/bun"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Organization struct {
	bun.BaseModel `bun:"table:organization"`

	ID           string `bun:",pk"`
	Name         string `bun:",unique"`
	Subscription int32
	CreatedAt    time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt    time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

func OrganizationFromProto(org *organizationv1.Organization) *Organization {
	if org == nil {
		return &Organization{}
	}

	return &Organization{
		ID:           org.GetId(),
		Name:         org.GetName(),
		Subscription: int32(org.GetSubscription()),
	}
}

func (o *Organization) ToProto() *organizationv1.Organization {
	return &organizationv1.Organization{
		Id:           o.ID,
		Name:         o.Name,
		Subscription: organizationv1.Subscription(o.Subscription),
		CreatedAt:    timestamppb.New(o.CreatedAt.UTC()),
		UpdatedAt:    timestamppb.New(o.UpdatedAt.UTC()),
	}
}

func (o *Organization) SetUpdatedAt(time time.Time) {
	o.UpdatedAt = time
}

func (o *Organization) ModelName() string {
	return "organization"
}
