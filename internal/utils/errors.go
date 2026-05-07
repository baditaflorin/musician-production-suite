package utils

import "log/slog"

func HandleErrorOrLogWithMessages(logger *slog.Logger, err error, errMsg string, successMsg string) {
	if err != nil {
		logger.Error(errMsg, "error", err)
		return
	}
	if successMsg != "" {
		logger.Info(successMsg)
	}
}
