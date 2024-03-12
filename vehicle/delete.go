package vehicle

import (
	"net/http"
	"strconv"

	"github.com/jlevesy/vehicle-server/storage"
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
		return
	}

	ok, err := d.store.Vehicle().Delete(r.Context(), id)
	if err != nil {
		r.logger.Error(
			"Could not list vehicles from store",
			zap.Error(err),
		)

		httputil.ServeError(rw, http.StatusInternalServerError, err)
		return
	}
	if ok {
		rw.WriteHeader(http.StatusNoContent)
	} else {
		rw.WriteHeader(http.StatusNotFound)
	}

}
