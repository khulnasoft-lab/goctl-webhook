package main

import (
	"os"

	"github.com/khulnasoft-lab/goctl-webhook/webhook"
)

func main() {
	if err := webhook.NewCmdForward().Execute(); err != nil {
		os.Exit(1)
	}
}
