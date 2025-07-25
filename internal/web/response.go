package web

import (
	"context"
	"encoding/json"
	"net/http"
)

func Respond(ctx context.Context, w http.ResponseWriter, data any, statusCode int) error {
	_ = SetStatusCode(ctx, statusCode)

	if statusCode == http.StatusNoContent || statusCode == http.StatusNotModified {
		w.WriteHeader(statusCode)
		return nil
	}

	if data == nil {
		data = struct{}{}
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(statusCode)

	if _, err := w.Write(jsonData); err != nil {
		return err
	}

	return nil
}
