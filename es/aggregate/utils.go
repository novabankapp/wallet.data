package aggregate

import (
	"context"
	"encoding/json"
	"github.com/EventStore/EventStore-Client-Go/esdb"
	base "github.com/novabankapp/common.data/domain/base"
	"github.com/novabankapp/common.data/eventstore"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/pkg/errors"
	"strings"
)

// GetWalletAggregateID get  aggregate id for eventstoredb
func GetWalletAggregateID(eventAggregateID string) string {
	return strings.ReplaceAll(eventAggregateID, "wallet-", "")
}

func IsAggregateNotFound(aggregate eventstore.Aggregate) bool {
	return aggregate.GetVersion() == 0
}

func GetJsonString(entity interface{}) (result string) {
	res, err := json.Marshal(entity)
	if err != nil {
		return ""
	}
	return string(res)
}
func GetEntityArrayFromJsonString[Entity base.NoSqlEntity](obj string) (result *[]Entity, error error) {
	var p []Entity
	err := json.Unmarshal([]byte(obj), &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}
func GetEntityFromJsonString[Entity base.NoSqlEntity](obj string) (result *Entity, error error) {
	var p Entity
	err := json.Unmarshal([]byte(obj), &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func LoadWalletAggregate(ctx context.Context, eventStore eventstore.AggregateStore, aggregateID string) (*WalletAggregate, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "LoadWalletAggregate")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", aggregateID))

	wallet := NewWalletAggregateWithID(aggregateID)

	err := eventStore.Exists(ctx, wallet.GetID())
	if err != nil && !errors.Is(err, esdb.ErrStreamNotFound) {
		return nil, err
	}

	if err := eventStore.Load(ctx, wallet); err != nil {
		return nil, err
	}

	return wallet, nil
}
