package aggregate

import (
	"context"
	"github.com/novabankapp/common.infrastructure/tracing"
	"github.com/novabankapp/wallet.data/constants"
	eventsV1 "github.com/novabankapp/wallet.data/es/events/v1"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

func (a *WalletAggregate) CreateWallet(ctx context.Context, amount decimal.Decimal, description, userId, accountId, id string) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "WalletAggregate.CreateWallet")
	defer span.Finish()
	span.LogFields(log.String(constants.AggregateID, a.GetID()))

	event, err := eventsV1.NewWalletCreatedEvent(a, amount, description, userId, accountId, id)
	if err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "NewWalletCreatedEvent")
	}

	if err := event.SetMetadata(tracing.ExtractTextMapCarrier(span.Context())); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "SetMetadata")
	}

	return a.Apply(event)
}

func (a *WalletAggregate) CreditWallet(ctx context.Context,
	debitWalletId string,
	amount decimal.Decimal,
	description string) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "WalletAggregate.CreditWallet")
	defer span.Finish()
	span.LogFields(log.String(constants.AggregateID, a.GetID()))

	event, err := eventsV1.NewWalletCreditEvent(a, debitWalletId, amount, description)
	if err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "NewWalletCreditEvent")
	}

	if err := event.SetMetadata(tracing.ExtractTextMapCarrier(span.Context())); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "SetMetadata")
	}

	return a.Apply(event)
}
func (a *WalletAggregate) DebitWallet(ctx context.Context,
	creditWalletId string,
	amount decimal.Decimal,
	description string) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "WalletAggregate.DebitWallet")
	defer span.Finish()
	span.LogFields(log.String(constants.AggregateID, a.GetID()))

	event, err := eventsV1.NewWalletDebitEvent(a, creditWalletId, amount, description)
	if err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "NewWalletDebitEvent")
	}

	if err := event.SetMetadata(tracing.ExtractTextMapCarrier(span.Context())); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "SetMetadata")
	}

	return a.Apply(event)
}
func (a *WalletAggregate) ReserveWalletCredit(
	ctx context.Context,
	amount decimal.Decimal,
	description string) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "WalletAggregate.ReserveWalletCredit")
	defer span.Finish()
	span.LogFields(log.String(constants.AggregateID, a.GetID()))

	event, err := eventsV1.NewWalletCreditReservedEvent(a, amount, description)
	if err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "NewWalletCreditReservedEvent")
	}

	if err := event.SetMetadata(tracing.ExtractTextMapCarrier(span.Context())); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "SetMetadata")
	}

	return a.Apply(event)
}
func (a *WalletAggregate) ReleaseWalletCredit(
	ctx context.Context,
	amount decimal.Decimal,
	description string) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "WalletAggregate.ReleaseWalletCredit")
	defer span.Finish()
	span.LogFields(log.String(constants.AggregateID, a.GetID()))

	event, err := eventsV1.NewWalletCreditReleasedEvent(a, amount, description)
	if err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "NewWalletCreditReleasedEvent")
	}

	if err := event.SetMetadata(tracing.ExtractTextMapCarrier(span.Context())); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "SetMetadata")
	}

	return a.Apply(event)
}
func (a *WalletAggregate) LockWallet(ctx context.Context, description string) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "WalletAggregate.LockWallet")
	defer span.Finish()
	span.LogFields(log.String(constants.AggregateID, a.GetID()))

	event, err := eventsV1.NewWalletLockedEvent(a, description)
	if err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "NewWalletLockedEvent")
	}

	if err := event.SetMetadata(tracing.ExtractTextMapCarrier(span.Context())); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "SetMetadata")
	}

	return a.Apply(event)
}
func (a *WalletAggregate) UnlockWallet(ctx context.Context, description string) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "WalletAggregate.UnlockWallet")
	defer span.Finish()
	span.LogFields(log.String(constants.AggregateID, a.GetID()))

	event, err := eventsV1.NewWalletUnlockedEvent(a, description)
	if err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "NewWalletUnlockedEvent")
	}

	if err := event.SetMetadata(tracing.ExtractTextMapCarrier(span.Context())); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "SetMetadata")
	}

	return a.Apply(event)
}
func (a *WalletAggregate) BlacklistWallet(ctx context.Context, description string) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "WalletAggregate.BlacklistWallet")
	defer span.Finish()
	span.LogFields(log.String(constants.AggregateID, a.GetID()))

	event, err := eventsV1.NewWalletBlacklistedEvent(a, description)
	if err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "NewWalletBlacklistedEvent")
	}

	if err := event.SetMetadata(tracing.ExtractTextMapCarrier(span.Context())); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "SetMetadata")
	}

	return a.Apply(event)
}
func (a *WalletAggregate) UnBlacklistWallet(ctx context.Context, description string) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "WalletAggregate.UnBlacklistWallet")
	defer span.Finish()
	span.LogFields(log.String(constants.AggregateID, a.GetID()))

	event, err := eventsV1.NewWalletUnBlacklistedEvent(a, description)
	if err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "NewWalletUnBlacklistedEvent")
	}

	if err := event.SetMetadata(tracing.ExtractTextMapCarrier(span.Context())); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "SetMetadata")
	}

	return a.Apply(event)
}
func (a *WalletAggregate) DeleteWallet(ctx context.Context, description string) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "WalletAggregate.DeleteWallet")
	defer span.Finish()
	span.LogFields(log.String(constants.AggregateID, a.GetID()))

	event, err := eventsV1.NewWalletDeletedEvent(a, description)
	if err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "NewWalletDeletedEvent")
	}

	if err := event.SetMetadata(tracing.ExtractTextMapCarrier(span.Context())); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "SetMetadata")
	}

	return a.Apply(event)
}
