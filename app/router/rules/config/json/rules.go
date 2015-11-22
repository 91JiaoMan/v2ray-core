package json

import (
	"github.com/v2ray/v2ray-core/app/point/config"
	v2net "github.com/v2ray/v2ray-core/common/net"
)

type Rule struct {
	Type        string `json:"type"`
	OutboundTag string `json:"outboundTag"`
}

func (this *Rule) Tag() *config.DetourTag {
	detourTag := config.DetourTag(this.OutboundTag)
	return &detourTag
}

func (this *Rule) Apply(dest v2net.Destination) bool {
	return false
}
