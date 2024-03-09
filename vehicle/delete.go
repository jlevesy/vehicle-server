package vehicle

import (
	"net/http"
	"strconv"

	"github.com/cicd-lectures/vehicle-server/pkg/httputil"
	"github.com/cicd-lectures/vehicle-server/storage"
	"go.uber.org/zap"
)

type DeleteHandler struct {
	store  storage.Store
	logger *zap.Logger
}

func NewDeleteHandler(store storage.Store, logger *zap.Logger) *DeleteHandler {
	return &DeleteHandler{
		store:  store,
		logger: logger.With(zap.String("handler", "delete_vehicles")),
	}
}

func (d *DeleteHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		d.logger.Error(
			"Could not parse vehicle id",
			zap.Error(err),
		)
	}

	ok, err := d.store.Vehicle().Delete(r.Context(), id)
	if err != nil {
		d.logger.Error(
			"Could not delete a vehicle from store",
			zap.Error(err),
		)

		httputil.ServeError(rw, http.StatusInternalServerError, err)
		return
	}

	if !ok {
		httputil.ServeJSON(rw, http.StatusNotFound, &httputil.APIError{
			Code:    httputil.ErrCodeResourceNotFound,
			Message: "Vehicle not Found",
		})
		return
	}

	rw.WriteHeader(http.StatusNoContent)
	return
}
