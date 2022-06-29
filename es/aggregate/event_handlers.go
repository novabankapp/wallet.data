package aggregate

import (
	"context"
	"github.com/google/uuid"
	es "github.com/novabankapp/common.data/eventstore"
	"github.com/novabankapp/common.data/tracing"
	"github.com/novabankapp/wallet.data/domain"
	v1 "github.com/novabankapp/wallet.data/es/events/v1"
	"github.com/novabankapp/wallet.data/es/models"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/pkg/errors"
	"time"
)

func (c *WalletProjection) onWalletCreated(ctx context.Context, evt es.Event) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "cassandraProjection.onWalletCreated")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", evt.GetAggregateID()))

	var eventData v1.WalletCreatedEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "evt.GetJsonData")
	}
	span.LogFields(log.String("WalletID", eventData.ID))

	op := models.WalletProjection{
		WalletID: GetWalletAggregateID(evt.AggregateID),
		UserID:   eventData.UserId,
		ID:       uuid.New().String(),
		Wallet: GetJsonString(domain.Wallet{
			ID:               eventData.ID,
			UserId:           eventData.UserId,
			AccountId:        eventData.AccountId,
			Balance:          eventData.Amount,
			AvailableBalance: eventData.Amount,
			CreatedAt:        time.Now(),
		}),
		WalletState: GetJsonString(domain.WalletState{
			WalletId:      eventData.ID,
			IsBlacklisted: false,
			IsLocked:      false,
		}),
		WalletTransactions: GetJsonString([]domain.WalletTransaction{}),
	}

	_, err := c.Repo.Create(ctx, op)
	if err != nil {
		return err
	}

	return nil
}

func (c *WalletProjection) onWalletCredited(ctx context.Context, evt es.Event) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "cassandraProjection.onWalletCredited")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", evt.GetAggregateID()))

	var eventData v1.WalletCreditedEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "evt.GetJsonData")
	}
	//span.LogFields(log.String("WalletID", eventData.ID))
	queries := make([]map[string]string, 1)
	m := make(map[string]string)
	m["column"] = "WalletID"
	m["compare"] = "="
	m["value"] = evt.GetAggregateID()
	queries = append(queries, m)
	ent, err := c.Repo.GetByCondition(ctx, queries)
	e := *ent
	if err != nil {
		return err
	}

	walletTransactionsP, _ := GetEntityArrayFromJsonString[domain.WalletTransaction](e.WalletTransactions)
	walletP, _ := GetEntityFromJsonString[domain.Wallet](e.Wallet)
	var wallet domain.Wallet = *walletP
	var walletTransactions []domain.WalletTransaction = *walletTransactionsP
	walletTransactions = append(walletTransactions, domain.WalletTransaction{
		DebitWalletId:  eventData.DebitWalletId,
		CreditWalletId: wallet.ID,
		Amount:         eventData.Amount,
		CreatedAt:      time.Now(),
		Description:    eventData.Description,
	})
	wallet.Balance = wallet.Balance.Add(eventData.Amount)
	e.Wallet = GetJsonString(wallet)
	e.WalletTransactions = GetJsonString(e.WalletTransactions)
	update, err := c.Repo.Update(ctx, e, e.ID)
	if err != nil {
		return err
	}
	if update {
		return nil
	} else {
		return errors.New("Not found")
	}

}

func (c *WalletProjection) onWalletDebited(ctx context.Context, evt es.Event) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "cassandraProjection.onWalletDebited")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", evt.GetAggregateID()))

	var eventData v1.WalletDebitedEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "evt.GetJsonData")
	}
	//span.LogFields(log.String("WalletID", eventData.ID))
	queries := make([]map[string]string, 1)
	m := make(map[string]string)
	m["column"] = "WalletID"
	m["compare"] = "="
	m["value"] = evt.GetAggregateID()
	queries = append(queries, m)
	ent, err := c.Repo.GetByCondition(ctx, queries)
	e := *ent
	if err != nil {
		return err
	}
	walletTransactionsP, _ := GetEntityArrayFromJsonString[domain.WalletTransaction](e.WalletTransactions)
	walletP, _ := GetEntityFromJsonString[domain.Wallet](e.Wallet)
	var wallet domain.Wallet = *walletP
	var walletTransactions []domain.WalletTransaction = *walletTransactionsP
	wallet.Balance = wallet.Balance.Sub(eventData.Amount)
	walletTransactions = append(walletTransactions, domain.WalletTransaction{
		DebitWalletId:  wallet.ID,
		CreditWalletId: eventData.CreditWalletId,
		Amount:         eventData.Amount,
		CreatedAt:      time.Now(),
		Description:    eventData.Description,
	})
	e.Wallet = GetJsonString(wallet)
	e.WalletTransactions = GetJsonString(walletTransactions)
	update, err := c.Repo.Update(ctx, e, e.ID)
	if err != nil {
		return err
	}
	if update {
		return nil
	} else {
		return errors.New("Not found")
	}
}

func (c *WalletProjection) onWalletCreditReserved(ctx context.Context, evt es.Event) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "cassandraProjection.onWalletCreditReserved")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", evt.GetAggregateID()))

	var eventData v1.WalletCreditedEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "evt.GetJsonData")
	}
	//span.LogFields(log.String("WalletID", eventData.ID))
	queries := make([]map[string]string, 1)
	m := make(map[string]string)
	m["column"] = "WalletID"
	m["compare"] = "="
	m["value"] = evt.GetAggregateID()
	queries = append(queries, m)
	ent, err := c.Repo.GetByCondition(ctx, queries)
	e := *ent
	if err != nil {
		return err
	}

	walletP, _ := GetEntityFromJsonString[domain.Wallet](e.Wallet)
	var wallet domain.Wallet = *walletP
	wallet.Balance = wallet.AvailableBalance.Sub(eventData.Amount)
	e.Wallet = GetJsonString(wallet)
	update, err := c.Repo.Update(ctx, e, e.ID)
	if err != nil {
		return err
	}
	if update {
		return nil
	} else {
		return errors.New("Not found")
	}
}
func (c *WalletProjection) onWalletBlacklisted(ctx context.Context, evt es.Event) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "cassandraProjection.onWalletBlacklisted")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", evt.GetAggregateID()))

	var eventData v1.WalletBlacklistedEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "evt.GetJsonData")
	}
	queries := make([]map[string]string, 1)
	m := make(map[string]string)
	m["column"] = "WalletID"
	m["compare"] = "="
	m["value"] = evt.GetAggregateID()
	queries = append(queries, m)
	ent, err := c.Repo.GetByCondition(ctx, queries)
	e := *ent
	if err != nil {
		return err
	}

	walletP, _ := GetEntityFromJsonString[domain.WalletState](e.WalletState)
	var walletState = *walletP
	walletState.IsBlacklisted = true
	e.WalletState = GetJsonString(walletState)
	update, err := c.Repo.Update(ctx, e, e.ID)
	if err != nil {
		return err
	}
	if update {
		return nil
	} else {
		return errors.New("Not found")
	}

}
func (c *WalletProjection) onWalletUnBlacklisted(ctx context.Context, evt es.Event) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "cassandraProjection.onWalletUnBlacklisted")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", evt.GetAggregateID()))

	var eventData v1.WalletUnBlacklistedEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "evt.GetJsonData")
	}
	queries := make([]map[string]string, 1)
	m := make(map[string]string)
	m["column"] = "WalletID"
	m["compare"] = "="
	m["value"] = evt.GetAggregateID()
	queries = append(queries, m)
	ent, err := c.Repo.GetByCondition(ctx, queries)
	e := *ent
	if err != nil {
		return err
	}
	walletP, _ := GetEntityFromJsonString[domain.WalletState](e.WalletState)
	var walletState = *walletP
	walletState.IsBlacklisted = false
	e.WalletState = GetJsonString(walletState)
	update, err := c.Repo.Update(ctx, e, e.ID)
	if err != nil {
		return err
	}
	if update {
		return nil
	} else {
		return errors.New("Not found")
	}

}

func (c *WalletProjection) onWalletLocked(ctx context.Context, evt es.Event) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "cassandraProjection.onWalletLocked")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", evt.GetAggregateID()))

	var eventData v1.WalletLockedEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "evt.GetJsonData")
	}
	queries := make([]map[string]string, 1)
	m := make(map[string]string)
	m["column"] = "WalletID"
	m["compare"] = "="
	m["value"] = evt.GetAggregateID()
	queries = append(queries, m)
	ent, err := c.Repo.GetByCondition(ctx, queries)
	e := *ent
	if err != nil {
		return err
	}
	walletP, _ := GetEntityFromJsonString[domain.WalletState](e.WalletState)
	var walletState = *walletP
	walletState.IsLocked = true
	e.WalletState = GetJsonString(walletState)
	update, err := c.Repo.Update(ctx, e, e.ID)
	if err != nil {
		return err
	}
	if update {
		return nil
	} else {
		return errors.New("Not found")
	}
}

func (c *WalletProjection) onWalletUnlocked(ctx context.Context, evt es.Event) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "cassandraProjection.onWalletUnlocked")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", evt.GetAggregateID()))

	var eventData v1.WalletUnlockedEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "evt.GetJsonData")
	}
	queries := make([]map[string]string, 1)
	m := make(map[string]string)
	m["column"] = "WalletID"
	m["compare"] = "="
	m["value"] = evt.GetAggregateID()
	queries = append(queries, m)
	ent, err := c.Repo.GetByCondition(ctx, queries)
	e := *ent
	if err != nil {
		return err
	}
	walletP, _ := GetEntityFromJsonString[domain.WalletState](e.WalletState)
	var walletState = *walletP
	walletState.IsLocked = false
	e.WalletState = GetJsonString(walletState)
	update, err := c.Repo.Update(ctx, e, e.ID)
	if err != nil {
		return err
	}
	if update {
		return nil
	} else {
		return errors.New("Not found")
	}
}

func (c *WalletProjection) onWalletCreditReleased(ctx context.Context, evt es.Event) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "cassandraProjection.onWalletCreditReleased")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", evt.GetAggregateID()))

	var eventData v1.WalletCreditedEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "evt.GetJsonData")
	}
	//span.LogFields(log.String("WalletID", eventData.ID))
	queries := make([]map[string]string, 1)
	m := make(map[string]string)
	m["column"] = "WalletID"
	m["compare"] = "="
	m["value"] = evt.GetAggregateID()
	queries = append(queries, m)
	ent, err := c.Repo.GetByCondition(ctx, queries)
	e := *ent
	if err != nil {
		return err
	}
	walletP, _ := GetEntityFromJsonString[domain.Wallet](e.WalletState)
	var wallet domain.Wallet = *walletP
	wallet.Balance = wallet.AvailableBalance.Add(eventData.Amount)
	e.Wallet = GetJsonString(wallet)

	update, err := c.Repo.Update(ctx, e, e.ID)
	if err != nil {
		return err
	}
	if update {
		return nil
	} else {
		return errors.New("Not found")
	}
}
