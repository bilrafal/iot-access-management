package repo_iot_simple

import (
	"context"
	"fmt"
	"iot-access-management/internal/db"
	db_factory "iot-access-management/internal/db/db_factory"
	"iot-access-management/internal/error/trace_error"
	repo_model "iot-access-management/internal/models/repo"
	"iot-access-management/internal/repo"
)

type RepoIotSimple struct {
	ctx      context.Context
	dbEngine db.DbClient
}

func NewRepoIotSimple(ctx context.Context, dbType db.DbType) repo.IotRepo {
	dbFactory := db_factory.NewDbFactorySimple(ctx, dbType)
	dbEng, err := dbFactory.GetDbClient()
	if err != nil {
		return nil
	}
	return &RepoIotSimple{
		ctx:      ctx,
		dbEngine: dbEng,
	}
}

func (r *RepoIotSimple) CreateWhiteList(whitelist repo_model.WhiteList) *trace_error.TraceError {

	dbErr := r.dbEngine.Save(db.WhiteListedDoorTableName, &whitelist)
	if dbErr != nil {
		return dbErr
	}
	return nil
}

func (r *RepoIotSimple) DeleteWhiteList(whitelist repo_model.WhiteList) *trace_error.TraceError {
	dbErr := r.dbEngine.Delete(db.WhiteListedDoorTableName, whitelist)
	if dbErr != nil {
		return dbErr
	}
	return nil
}

func (r *RepoIotSimple) ListWhiteList() ([]repo_model.WhiteList, *trace_error.TraceError) {

	whiteList, dbErr := r.dbEngine.Get(db.WhiteListedDoorTableName, nil)
	if dbErr != nil {
		return nil, dbErr
	}

	arr, ok := whiteList.([]any)
	if !ok {
		return nil, db.ErrUnexpected.From(fmt.Errorf("unexpected type: %T", whiteList))
	}
	var result []repo_model.WhiteList
	for _, wl := range arr {
		d, ok := wl.(*repo_model.WhiteList)
		if !ok {
			return nil, db.ErrUnexpected.From(fmt.Errorf("unexpected type: %T", whiteList))
		}
		result = append(result, *d)
	}
	return result, nil
}
