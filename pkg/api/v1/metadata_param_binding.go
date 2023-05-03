package api

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"

	"go.infratographer.com/load-balancer-api/internal/models"
)

func (r *Router) metadataParamBidningFactory(pathID string) func(c echo.Context) ([]qm.QueryMod, error) {
	return func(c echo.Context) ([]qm.QueryMod, error) {
		var (
			foundID string
			err     error
			mods    []qm.QueryMod
		)

		// optional load_balancer_id in the request path
		foundID, err = r.parseUUID(c, pathID)
		if err != nil {
			if !errors.Is(err, ErrUUIDNotFound) {
				return nil, err
			}
		} else {
			// found load_balancer_id in path so add to query mods
			mods = append(mods, models.LoadBalancerMetadatumWhere.LoadBalancerID.EQ(foundID))
			r.logger.Debug("path param", zap.String(pathID, foundID))
		}

		if foundID == "" {
			r.logger.Debug(pathID + " required in the path")
			return nil, ErrIDRequired
		}
		// query params
		queryParams := []string{"namespace", "metadata_id"}

		qpb := echo.QueryParamsBinder(c)

		for _, qp := range queryParams {
			mods = queryParamsToQueryMods(qpb, qp, mods)

			if len(c.QueryParam(qp)) > 0 {
				r.logger.Error("metadata query parameters", zap.String("query.key", qp), zap.String("query.value", c.QueryParam(qp)), zap.String(pathID, foundID))
			}
		}

		if err := qpb.BindError(); err != nil {
			return nil, err
		}

		return mods, nil
	}
}

func (r *Router) lbMetadataParamsBinding(c echo.Context) ([]qm.QueryMod, error) {
	return r.metadataParamBidningFactory("load_balancer_id")(c)
}

func (r *Router) oMetadataParamsBinding(c echo.Context) ([]qm.QueryMod, error) {
	return r.metadataParamBidningFactory("origin_id")(c)
}
