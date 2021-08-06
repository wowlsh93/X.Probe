package influxdb

import (
	"context"
	"github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

type Influxdb2Connector struct {
	client   influxdb2.Client
	writeAPI api.WriteAPIBlocking
}

func NewInfluxdb2Connector(dbpath string,
	authToken string,
	organization string,
	bucket string) *Influxdb2Connector {

	client := influxdb2.NewClient(dbpath, authToken)

	return &Influxdb2Connector{
		client:   client,
		writeAPI: client.WriteAPIBlocking(organization, bucket),
	}
}

func (connector *Influxdb2Connector) Write(p *write.Point) {
	connector.writeAPI.WritePoint(context.Background(), p)
}

func (connector *Influxdb2Connector) Close() {
	connector.client.Close()
}
