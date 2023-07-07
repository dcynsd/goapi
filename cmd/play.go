package cmd

import (
	"time"

	"goapi/pkg/console"
	"goapi/pkg/redis"

	"github.com/spf13/cobra"
)

var CmdPlay = &cobra.Command{
	Use:   "play",
	Short: "Likes the Go Playground, but running at our application context",
	Run:   runPlay,
}

func runPlay(cmd *cobra.Command, args []string) {
	redis.Redis.Set("hello", "hi from redis", 10*time.Second)
	console.Success(redis.Redis.Get("hello"))
}
