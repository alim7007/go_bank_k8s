package gapi

import (
	"context"
	"log"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

type Metadata struct {
	UserAgent string
	ClientIp  string
}

const (
	GrpcGatewayUserAgent = "grpcgateway-user-agent"
	UserAgent            = "user-agent"
	xForwardedForHeader  = "x-forwarded-for"
)

func (server *Server) extractMetadata(ctx context.Context) *Metadata {
	mtdt := &Metadata{}

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		log.Printf("md: %+v\n", md)
		if userAgent := md.Get(GrpcGatewayUserAgent); len(userAgent) > 0 {
			mtdt.UserAgent = userAgent[0]
		}
		if userAgent := md.Get(UserAgent); len(userAgent) > 0 {
			mtdt.UserAgent = userAgent[0]
		}
		if clientIp := md.Get(GrpcGatewayUserAgent); len(clientIp) > 0 {
			mtdt.ClientIp = clientIp[0]
		}
	}

	if p, ok := peer.FromContext(ctx); ok {
		mtdt.ClientIp = p.Addr.String()
	}

	return mtdt
}
