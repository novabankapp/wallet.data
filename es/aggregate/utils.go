package aggregate

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"strings"

	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/novabankapp/common.data/eventstore"
	"github.com/opentracing/opentracing-go/log"
	"github.com/pkg/errors"
)

// GetWalletAggregateID get order aggregate id for eventstoredb
func GetWalletAggregateID(eventAggregateID string) string {
	return strings.ReplaceAll(eventAggregateID, "wallet-", "")
}

func IsAggregateNotFound(aggregate eventstore.Aggregate) bool {
	return aggregate.GetVersion() == 0
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
