package storage

import (
	"database/sql"
	"time"

	"github.com/feast-dev/feast/backend/src/apiserver/model"
	util "github.com/feast-dev/feast/backend/src/utils"

	sq "github.com/Masterminds/squirrel"
)

// The order of the selected columns must match the order used in scan rows.
var miColumns = []string{
	"materialization_intervals.fvid",
	"materialization_intervals.start_time",
	"materialization_intervals.end_time",
}

type MIStoreInterface interface {
	CreateMINoTx(mi *model.MaterializationInterval, fvid string) (*model.MaterializationInterval, error)
	CreateMI(tx *sql.Tx, mi *model.MaterializationInterval, fvid string) (*model.MaterializationInterval, error)
	DeleteMIs(tx *sql.Tx, fvid string) error
	ListMIs(fvid string) ([]*model.MaterializationInterval, error)
}

type MIStore struct {
	db   *DB
	time util.TimeInterface
	uuid util.UUIDGeneratorInterface
}

func (s *MIStore) CreateMINoTx(mi *model.MaterializationInterval, fvid string) (*model.MaterializationInterval, error) {
	newMI := *mi
	newMI.FVId = fvid

	sql, args, err := sq.
		Insert("materialization_intervals").
		SetMap(
			sq.Eq{
				"fvid":       newMI.FVId,
				"start_time": newMI.StartTime,
				"end_time":   newMI.EndTime,
			}).
		ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to insert materialization interval: %v", err.Error())
	}

	_, err = s.db.Exec(sql, args...)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to store materialization interval.")
	}

	return &newMI, nil
}

func (s *MIStore) CreateMI(tx *sql.Tx, mi *model.MaterializationInterval, fvid string) (*model.MaterializationInterval, error) {
	newMI := *mi
	newMI.FVId = fvid

	sql, args, err := sq.
		Insert("materialization_intervals").
		SetMap(
			sq.Eq{
				"fvid":       newMI.FVId,
				"start_time": newMI.StartTime,
				"end_time":   newMI.EndTime,
			}).
		ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to insert materialization interval: %v", err.Error())
	}

	_, err = tx.Exec(sql, args...)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to store materialization interval.")
	}

	return &newMI, nil
}

func (s *MIStore) DeleteMIs(tx *sql.Tx, fvid string) error {
	sql, args, err := sq.
		Delete("materialization_intervals").
		Where(sq.Eq{"materialization_intervals.fvid": fvid}).
		ToSql()
	if err != nil {
		return util.NewInternalServerError(err, "Failed to create query to delete materialization intervals: %v", err.Error())
	}

	_, err = tx.Exec(sql, args...)
	if err != nil {
		return util.NewInternalServerError(err, "Failed to delete materialization intervals: %v", err.Error())
	}
	return nil
}

func (s *MIStore) ListMIs(fvid string) ([]*model.MaterializationInterval, error) {
	sql, args, err := sq.
		Select(miColumns...).
		From("materialization_intervals").
		Where(sq.Eq{"materialization_intervals.fvid": fvid}).
		ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to list materialization intervals: %v", err.Error())
	}

	r, err := s.db.Query(sql, args...)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to list materialization intervals: %v", err.Error())
	}
	defer r.Close()

	mis, err := s.scanRows(r)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to list materialization intervals: %v", err.Error())
	}

	return mis, nil
}

func (s *MIStore) scanRows(rows *sql.Rows) ([]*model.MaterializationInterval, error) {
	var mis []*model.MaterializationInterval
	for rows.Next() {
		var fvid string
		var start_time, end_time time.Time

		if err := rows.Scan(
			&fvid,
			&start_time,
			&end_time); err != nil {
			return nil, err
		}

		mis = append(mis, &model.MaterializationInterval{
			FVId:      fvid,
			StartTime: start_time,
			EndTime:   end_time,
		})
	}

	return mis, nil
}

func NewMIStore(db *DB, time util.TimeInterface, uuid util.UUIDGeneratorInterface) *MIStore {
	return &MIStore{db: db, time: time, uuid: uuid}
}
