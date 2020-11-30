package logger

import (
	"context"
	"fmt"

	"github.com/giantswarm/micrologger"
)

type logger struct {
	parent micrologger.Logger
}

type ctxLogger struct {
	ctx    context.Context
	parent micrologger.Logger
}

func Wrap(l *micrologger.MicroLogger) Logger {
	return &logger{parent: l}
}

func (l *logger) Debug(format string, params ...interface{}) {
	l.parent.Log("level", "debug", "message", fmt.Sprintf(format, params...))
}

func (l *logger) Info(format string, params ...interface{}) {
	l.parent.Log("level", "info", "message", fmt.Sprintf(format, params...))
}

func (l *logger) Warning(format string, params ...interface{}) {
	l.parent.Log("level", "warning", "message", fmt.Sprintf(format, params...))
}

func (l *logger) Error(format string, params ...interface{}) {
	l.parent.Log("level", "error", "message", fmt.Sprintf(format, params...))
}

func (l *logger) WithCtx(ctx context.Context) Logger {
	return &ctxLogger{
		ctx:    ctx,
		parent: l.parent,
	}
}

func (l *ctxLogger) Debug(format string, params ...interface{}) {
	l.parent.LogCtx(l.ctx, "level", "debug", "message", fmt.Sprintf(format, params...))
}

func (l *ctxLogger) Info(format string, params ...interface{}) {
	l.parent.LogCtx(l.ctx, "level", "info", "message", fmt.Sprintf(format, params...))
}

func (l *ctxLogger) Warning(format string, params ...interface{}) {
	l.parent.LogCtx(l.ctx, "level", "warning", "message", fmt.Sprintf(format, params...))
}

func (l *ctxLogger) Error(format string, params ...interface{}) {
	l.parent.LogCtx(l.ctx, "level", "error", "message", fmt.Sprintf(format, params...))
}
