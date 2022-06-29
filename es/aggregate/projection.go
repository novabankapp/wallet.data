package aggregate

import (
	"context"
	"github.com/EventStore/EventStore-Client-Go/esdb"
	es "github.com/novabankapp/common.data/eventstore"
	"github.com/novabankapp/common.data/eventstore/projections"
	"github.com/novabankapp/common.data/repositories/base"
	"github.com/novabankapp/common.data/tracing"
	v1 "github.com/novabankapp/wallet.data/es/events/v1"
	"github.com/novabankapp/wallet.data/es/models"
	"github.com/opentracing/opentracing-go/log"
	"github.com/pkg/errors"
)

const (
	CassProjection = "(CassandraDB Projection)"
)

type WalletProjection struct {
	projections.CassandraProjection
	Repo base.NoSqlRepository[models.WalletProjection]
}

func (c *WalletProjection) ProcessEvents(ctx context.Context, stream *esdb.PersistentSubscription, workerID int) error {

	for {
		event := stream.Recv()
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		if event.SubscriptionDropped != nil {
			c.Log.Errorf("(SubscriptionDropped) err: {%v}", event.SubscriptionDropped.Error)
			return errors.Wrap(event.SubscriptionDropped.Error, "Subscription Dropped")
		}

		if event.EventAppeared != nil {
			c.Log.ProjectionEvent(CassProjection, c.Cfg.CassandraProjectionGroupName, event.EventAppeared, workerID)

			err := c.When(ctx, es.NewEventFromRecorded(event.EventAppeared.Event))
			if err != nil {
				c.Log.Errorf("(CassProjection.when) err: {%v}", err)

				if err := stream.Nack(err.Error(), esdb.Nack_Retry, event.EventAppeared); err != nil {
					c.Log.Errorf("(stream.Nack) err: {%v}", err)
					return errors.Wrap(err, "stream.Nack")
				}
			}

			err = stream.Ack(event.EventAppeared)
			if err != nil {
				c.Log.Errorf("(stream.Ack) err: {%v}", err)
				return errors.Wrap(err, "stream.Ack")
			}
			c.Log.Infof("(ACK) event commit: {%v}", *event.EventAppeared.Commit)
		}
	}
}
func (c *WalletProjection) When(ctx context.Context, evt es.Event) error {
	ctx, span := tracing.StartProjectionTracerSpan(ctx, "cassandraProjection.When", evt)
	defer span.Finish()
	span.LogFields(log.String("AggregateID", evt.GetAggregateID()), log.String("EventType", evt.GetEventType()))

	switch evt.GetEventType() {

	case v1.WalletCreated:
		return c.onWalletCreated(ctx, evt)
	case v1.WalletCredited:
		return c.onWalletCredited(ctx, evt)
	case v1.WalletDebited:
		return c.onWalletDebited(ctx, evt)
	case v1.WalletCreditReserved:
		return c.onWalletCreditReserved(ctx, evt)
	case v1.WalletBlacklisted:
		return c.onWalletBlacklisted(ctx, evt)
	case v1.WalletLocked:
		return c.onWalletLocked(ctx, evt)
	case v1.WalletUnBlacklisted:
		return c.onWalletUnBlacklisted(ctx, evt)
	case v1.WalletUnlocked:
		return c.onWalletUnlocked(ctx, evt)
	case v1.WalletCreditReleased:
		return c.onWalletCreditReleased(ctx, evt)

	default:
		c.CassandraProjection.Log.Warnf("(cassandraProjection) [When unknown EventType] eventType: {%s}", evt.EventType)
		return es.ErrInvalidEventType
	}
}
