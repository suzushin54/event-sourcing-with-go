package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"suzushin54/event-sourcing-with-go/cmd/di"
	"suzushin54/event-sourcing-with-go/config"
	"suzushin54/event-sourcing-with-go/pkg"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("terminated server: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	c, err := config.NewConfig()
	if err != nil {
		return fmt.Errorf("failed to build config: %v", err)
	}

	env, err := config.LoadEnv()
	if err != nil {
		return fmt.Errorf("failed to load env: %v", err)
	}
	db, err := pkg.NewDBClient(ctx, env)
	if err != nil {
		return fmt.Errorf("failed to setup mysql client: %w", err)
	}

	node, err := pkg.NewSnowflakeNode()
	if err != nil {
		return fmt.Errorf("failed to create snowflake node: %v", err)
	}
	httpServer := di.InitServer(db, node)

	if err := httpServer.Run(ctx, c.Port); err != nil {
		fmt.Printf("failed to terminate server: %v", err)
	}

	return nil
}
