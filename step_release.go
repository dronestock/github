package main

import (
	"context"
)

type releaseStep struct {
	*plugin
}

func newReleaseStep(plugin *plugin) *releaseStep {
	return &releaseStep{
		plugin: plugin,
	}
}

func (s *releaseStep) Runnable() bool {
	return nil != s.Release
}

func (s *releaseStep) Run(_ context.Context) error {
	return s.Release.publish(s.plugin)
}
