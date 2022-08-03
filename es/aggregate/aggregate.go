package aggregate

import (
	es "github.com/novabankapp/common.data/eventstore"
	"github.com/novabankapp/wallet.data/domain"
	event1 "github.com/novabankapp/wallet.data/es/events/v1"
	v1 "github.com/novabankapp/wallet.data/es/events/v1"
	"github.com/pkg/errors"
	"time"
)

const (
	WalletAggregateType es.AggregateType = "wallet"
)

type WalletAggregate struct {
	*es.AggregateBase
	Wallet             *domain.Wallet
	WalletState        *domain.WalletState
	WalletLink         *domain.WalletLink
	WalletTransactions *[]domain.WalletTransaction
}

func NewWalletAggregateWithID(id string) *WalletAggregate {
	if id == "" {
		return nil
	}

	aggregate := NewWalletAggregate()
	aggregate.SetID(id)
	aggregate.Wallet.ID = id
	return aggregate
}

func NewWalletAggregate() *WalletAggregate {
	walletAggregate := &WalletAggregate{Wallet: &domain.Wallet{}}
	base := es.NewAggregateBase(walletAggregate.When)
	base.SetType(WalletAggregateType)
	walletAggregate.AggregateBase = base
	return walletAggregate
}

func (a *WalletAggregate) When(evt es.Event) error {

	switch evt.GetEventType() {

	case event1.WalletCreated:
		return a.onWalletCreated(evt)
	case event1.WalletDebited:
		return a.onWalletDebited(evt)
	case event1.WalletCredited:
		return a.onWalletCredited(evt)
	case event1.WalletBlacklisted:
		return a.onWalletBlacklisted(evt)
	case event1.WalletLocked:
		return a.onWalletLocked(evt)
	case event1.WalletCreditReleased:
		return a.onWalletCreditReleased(evt)
	case event1.WalletCreditReserved:
		return a.onWalletCreditReserved(evt)

	default:
		return es.ErrInvalidEventType
	}
}

func (a *WalletAggregate) onWalletCreated(evt es.Event) error {
	var eventData v1.WalletCreatedEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		return errors.Wrap(err, "GetJsonData")
	}

	a.Wallet.AccountId = eventData.AccountId
	a.Wallet.UserId = eventData.UserId
	a.Wallet.CreatedAt = time.Now()
	a.Wallet.Balance = eventData.Amount
	a.Wallet.ID = eventData.ID
	a.Wallet.AvailableBalance = eventData.Amount

	return nil
}

func (a *WalletAggregate) onWalletCreditReleased(evt es.Event) error {
	var eventData v1.WalletCreditReleasedEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		return errors.Wrap(err, "GetJsonData")
	}

	a.Wallet.Balance = a.Wallet.AvailableBalance.Add(eventData.Amount)
	return nil
}
func (a *WalletAggregate) onWalletCreditReserved(evt es.Event) error {
	var eventData v1.WalletCreditReservedEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		return errors.Wrap(err, "GetJsonData")
	}

	a.Wallet.Balance = a.Wallet.AvailableBalance.Sub(eventData.Amount)
	return nil
}

func (a *WalletAggregate) onWalletCredited(evt es.Event) error {
	var eventData v1.WalletCreditedEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		return errors.Wrap(err, "GetJsonData")
	}
	a.Wallet.Lock.Lock()
	defer a.Wallet.Lock.Unlock()
	a.Wallet.Balance = a.Wallet.Balance.Add(eventData.Amount)
	*a.WalletTransactions = append(*a.WalletTransactions, domain.WalletTransaction{
		DebitWalletId:  eventData.DebitWalletId,
		CreditWalletId: a.Wallet.ID,
		Amount:         eventData.Amount,
		CreatedAt:      time.Now(),
		Description:    eventData.Description,
	})
	return nil
}

func (a *WalletAggregate) onWalletDebited(evt es.Event) error {
	var eventData v1.WalletDebitedEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		return errors.Wrap(err, "GetJsonData")
	}

	a.Wallet.Lock.Lock()
	defer a.Wallet.Lock.Unlock()

	a.Wallet.Balance = a.Wallet.Balance.Sub(eventData.Amount)
	*a.WalletTransactions = append(*a.WalletTransactions, domain.WalletTransaction{
		Amount:         eventData.Amount,
		CreatedAt:      time.Now(),
		Description:    eventData.Description,
		CreditWalletId: eventData.CreditWalletId,
		DebitWalletId:  a.Wallet.ID,
	})
	return nil
}

func (a *WalletAggregate) onWalletLocked(evt es.Event) error {
	var eventData v1.WalletLockedEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		return errors.Wrap(err, "GetJsonData")
	}
	a.WalletState.IsLocked = true
	return nil
}

func (a *WalletAggregate) onWalletBlacklisted(evt es.Event) error {
	var eventData v1.WalletBlacklistedEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		return errors.Wrap(err, "GetJsonData")
	}
	a.WalletState.IsBlacklisted = true
	return nil
}
