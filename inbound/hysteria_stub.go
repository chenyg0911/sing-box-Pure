//go:build !with_quic

package inbound

import (
	"context"

	"github.com/inazumav/sing-box/adapter"
	C "github.com/inazumav/sing-box/constant"
	"github.com/inazumav/sing-box/log"
	"github.com/inazumav/sing-box/option"
)

func NewHysteria(ctx context.Context, router adapter.Router, logger log.ContextLogger, tag string, options option.HysteriaInboundOptions) (adapter.Inbound, error) {
	return nil, C.ErrQUICNotIncluded
}

type Hysteria struct {
	adapter.Inbound
}

func (h *Hysteria) Start() error {
	return C.ErrQUICNotIncluded
}

func (h *Hysteria) AddUsers(_ []option.HysteriaUser) error {
	return C.ErrQUICNotIncluded
}

func (h *Hysteria) DelUsers(_ []string) error {
	return C.ErrQUICNotIncluded
}

func NewHysteria2(ctx context.Context, router adapter.Router, logger log.ContextLogger, tag string, options option.Hysteria2InboundOptions) (adapter.Inbound, error) {
	return nil, C.ErrQUICNotIncluded
}
