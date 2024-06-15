package logger

import (
	"fmt"
	"log/slog"
	"os"
)

func Init() {
	handler := slog.NewJSONHandler(
		os.Stdout,
		&slog.HandlerOptions{
			AddSource: true,
			ReplaceAttr: func(_ []string, attr slog.Attr) slog.Attr {
				if attr.Key == slog.SourceKey {
					attr.Value = slog.StringValue(
						fmt.Sprintf(
							"%s:%d",
							attr.Value.Any().(*slog.Source).File,
							attr.Value.Any().(*slog.Source).Line,
						),
					)
				}
				return attr
			},
		},
	)
	slog.SetDefault(slog.New(handler))
}
