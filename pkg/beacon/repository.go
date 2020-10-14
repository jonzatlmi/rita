package beacon

import (
	"github.com/activecm/rita/pkg/data"
	"github.com/activecm/rita/pkg/uconn"
	"github.com/globalsign/mgo/bson"
)

// Repository for host collection
type Repository interface {
	CreateIndexes() error
	Upsert(uconnMap map[string]*uconn.Pair)
}

type updateInfo struct {
	selector bson.M
	query    bson.M
}

//update ....
type update struct {
	beacon     updateInfo
	hostIcert  updateInfo
	hostBeacon updateInfo
	uconn      updateInfo
}

//TSData ...
type TSData struct {
	Range      int64   `bson:"range"`
	Mode       int64   `bson:"mode"`
	ModeCount  int64   `bson:"mode_count"`
	Skew       float64 `bson:"skew"`
	Dispersion int64   `bson:"dispersion"`
	Duration   float64 `bson:"duration"`
}

//DSData ...
type DSData struct {
	Skew       float64 `bson:"skew"`
	Dispersion int64   `bson:"dispersion"`
	Range      int64   `bson:"range"`
	Mode       int64   `bson:"mode"`
	ModeCount  int64   `bson:"mode_count"`
}

//AnalysisView (for reporting)
type AnalysisView struct {
	data.UniqueIPPair `bson:",inline"`
	Connections       int64   `bson:"connection_count"`
	AvgBytes          float64 `bson:"avg_bytes"`
	Ts                TSData  `bson:"ts"`
	Ds                DSData  `bson:"ds"`
	Score             float64 `bson:"score"`
}

//StrobeAnalysisView (for reporting)
type StrobeAnalysisView struct {
	data.UniqueIPPair `bson:",inline"`
	ConnectionCount   int64 `bson:"connection_count"`
}
