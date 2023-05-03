package api

import (
	"database/sql"
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/types"
	"go.uber.org/zap"

	"go.infratographer.com/load-balancer-api/internal/models"
)

type metadataCreateModels struct {
	pathID    string
	objDel    func(echo.Context, []qm.QueryMod) error
	objGetOne func(echo.Context, string) error
	objInsert func(echo.Context, string, string, types.JSON) (string, error)
	objList   func(echo.Context, []qm.QueryMod) error
}

func (r *Router) metadataCreateFactory(m *metadataCreateModels) func(c echo.Context) error {
	return func(c echo.Context) error {
		payload := struct {
			Namespace string     `json:"namespace"`
			Data      types.JSON `json:"data"`
		}{}

		if err := c.Bind(&payload); err != nil {
			r.logger.Error("error binding payload", zap.Error(err))
			return v1BadRequestResponse(c, err)
		}

		objID, err := r.parseUUID(c, m.pathID)
		if err != nil {
			return v1BadRequestResponse(c, err)
		}

		if err := m.objGetOne(c, objID); err != nil {
			return err
		}

		if payload.Namespace == "" {
			return v1BadRequestResponse(c, ErrNamespaceRequired)
		}

		mdID, err := m.objInsert(c, objID, payload.Namespace, payload.Data)
		if err != nil {
			r.logger.Error("error inserting metadata", zap.Error(err))
			return v1BadRequestResponse(c, err)
		}

		return v1MetadataCreatedResponse(c, mdID)
	}
}

func (r *Router) lbExists(c echo.Context, lbID string) error {
	ctx := c.Request().Context()
	// validate load balancer exists
	if _, err := models.LoadBalancers(
		models.LoadBalancerWhere.LoadBalancerID.EQ(lbID),
	).One(ctx, r.db); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			r.logger.Error("error fetching load balancer not found by id", zap.Error(err))
			return v1NotFoundResponse(c)
		}

		r.logger.Error("error fetching load balancer", zap.Error(err))

		return v1InternalServerErrorResponse(c, err)
	}

	return nil
}

func (r *Router) lbMetadataInsert(c echo.Context, lbID, ns string, data types.JSON) (string, error) {
	ctx := c.Request().Context()

	metadataModel := &models.LoadBalancerMetadatum{
		LoadBalancerID: lbID,
		Namespace:      ns,
		Data:           data,
	}

	if err := metadataModel.Insert(ctx, r.db, boil.Infer()); err != nil {
		r.logger.Error("error inserting metadata", zap.Error(err))
		return "", v1BadRequestResponse(c, err)
	}

	return metadataModel.MetadataID, nil
}

func (r *Router) createLBMetadata(c echo.Context) error {
	return r.metadataCreateFactory(&metadataCreateModels{
		pathID:    "load_balancer_id",
		objGetOne: r.lbExists,
		objInsert: r.lbMetadataInsert,
	})(c)
}

func (r *Router) originExists(c echo.Context, oID string) error {
	ctx := c.Request().Context()
	// validate load balancer exists
	if _, err := models.Origins(
		models.OriginWhere.OriginID.EQ(oID),
	).One(ctx, r.db); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			r.logger.Error("error fetching load balancer not found by id", zap.Error(err))
			return v1NotFoundResponse(c)
		}

		r.logger.Error("error fetching load balancer", zap.Error(err))

		return v1InternalServerErrorResponse(c, err)
	}

	return nil
}

func (r *Router) originMetadataInsert(c echo.Context, oID, ns string, data types.JSON) (string, error) {
	ctx := c.Request().Context()

	metadataModel := &models.OriginMetadatum{
		OriginID:  oID,
		Namespace: ns,
		Data:      data,
	}

	if err := metadataModel.Insert(ctx, r.db, boil.Infer()); err != nil {
		r.logger.Error("error inserting metadata", zap.Error(err))
		return "", v1BadRequestResponse(c, err)
	}

	return metadataModel.MetadataID, nil
}

func (r *Router) createOriginMetadata(c echo.Context) error {
	return r.metadataCreateFactory(&metadataCreateModels{
		pathID:    "origin_id",
		objGetOne: r.originExists,
		objInsert: r.originMetadataInsert,
	})(c)
}
