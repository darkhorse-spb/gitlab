package main

import (
	"github.com/darkhorse-spb/integram"
	"github.com/darkhorse-spb/gitlab"
	"github.com/kelseyhightower/envconfig"
)

func main(){
	var cfg gitlab.Config
	envconfig.MustProcess("GITLAB", &cfg)

	integram.Register(
		cfg,
		cfg.BotConfig.Token,
	)

	integram.Run()
}