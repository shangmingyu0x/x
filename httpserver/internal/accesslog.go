package internal

import (
	"github/shangmingyu0x/x/infra"
	"net/http"

	"go.uber.org/zap"
)

// AccessLog .
func AccessLog(r *http.Request, statusCode int) {
	infra.BaseLogger.Info("access_log", zap.String("path: ", r.URL.Path),
		zap.String("client_ip: ", clientIP(r)), zap.Int("response_status_code", statusCode))
}
