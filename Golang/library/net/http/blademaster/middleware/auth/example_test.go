package auth_test

import (
	"fmt"

	bm "wwqdrh/handbook/librarynet/http/blademaster"
	"wwqdrh/handbook/librarynet/http/blademaster/middleware/auth"
	"wwqdrh/handbook/librarynet/metadata"
	"wwqdrh/handbook/librarynet/rpc/warden"
)

// This example create a identify middleware instance and attach to several path,
// it will validate request by specified policy and put extra information into context. e.g., `mid`.
// It provides additional handler functions to provide the identification for your business handler.
func Example() {
	authn := auth.New(&auth.Config{
		Identify:    &warden.ClientConfig{},
		DisableCSRF: false,
	})

	e := bm.DefaultServer(nil)

	// mark `/user` path as User policy
	e.GET("/user", authn.User, func(ctx *bm.Context) {
		mid := metadata.Int64(ctx, metadata.Mid)
		ctx.JSON(fmt.Sprintf("%d", mid), nil)
	})
	// mark `/mobile` path as UserMobile policy
	e.GET("/mobile", authn.UserMobile, func(ctx *bm.Context) {
		mid := metadata.Int64(ctx, metadata.Mid)
		ctx.JSON(fmt.Sprintf("%d", mid), nil)
	})
	// mark `/web` path as UserWeb policy
	e.GET("/web", authn.UserWeb, func(ctx *bm.Context) {
		mid := metadata.Int64(ctx, metadata.Mid)
		ctx.JSON(fmt.Sprintf("%d", mid), nil)
	})
	// mark `/guest` path as Guest policy
	e.GET("/guest", authn.Guest, func(ctx *bm.Context) {
		mid := metadata.Int64(ctx, metadata.Mid)
		ctx.JSON(fmt.Sprintf("%d", mid), nil)
	})

	e.Run(":18080")
}
