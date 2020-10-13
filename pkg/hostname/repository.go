package hostname

import (
	"github.com/activecm/rita/pkg/data"
	"github.com/globalsign/mgo/bson"
)

// Repository for hostnames collection
type Repository interface {
	CreateIndexes() error
	Upsert(domainMap map[string]*Input)
}

//update ....
type update struct {
	selector bson.M
	query    bson.M
}

//Input ....
type Input struct {
	Host        string           //A hostname
	ResolvedIPs data.UniqueIPSet //Set of resolved UniqueIPs associated with a given hostname
	ClientIPs   data.UniqueIPSet //Set of DNS Client UniqueIPs which issued queries for a given hostname
}

//AnalysisView (for reporting)
type AnalysisView struct {
	Host              string          `bson:"host"`
	Connections       int             `bson:"conn_count"`
	UniqueConnections int             `bson:"uconn_count"`
	TotalBytes        int             `bson:"total_bytes"`
	ConnectedHosts    []data.UniqueIP `bson:"sources,omitempty"`
}
