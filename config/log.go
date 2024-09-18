package config

import (
	"fmt"

	"go.uber.org/zap"
)

func initializeLog() error {
	var err error
	logger, err = zap.NewDevelopment()
	if err != nil {
		return fmt.Errorf("Error configuring logger : %w", err)
	}
	return nil
}
