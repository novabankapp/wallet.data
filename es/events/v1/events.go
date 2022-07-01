package v1

import (
	es "github.com/novabankapp/common.data/eventstore"
	"github.com/shopspring/decimal"
)

const (
	WalletCreated        = "V1_WALLET_CREATED"
	WalletDebited        = "V1_WALLET_DEBITED"
	WalletCredited       = "V1_WALLET_CREDITED"
	WalletLocked         = "V1_WALLET_LOCKED"
	WalletBlacklisted    = "V1_WALLET_BLACKLISTED"
	WalletUnlocked       = "V1_WALLET_UNLOCKED"
	WalletDeleted        = "V1_WALLET_DELETED"
	WalletUnBlacklisted  = "V1_WALLET_UNBLACKLISTED"
	WalletCreditReserved = "V1_WALLET_CREDIT_RESERVED"
	WalletCreditReleased = "V1_WALLET_CREDIT_RELEASED"
)

type WalletCreatedEvent struct {
	Amount      decimal.Decimal
	Description string
	UserId      string
	AccountId   string
	ID          string
}
type WalletDebitedEvent struct {
	Amount         decimal.Decimal
	CreditWalletId string
	Description    string
}
type WalletCreditedEvent struct {
	Amount        decimal.Decimal
	DebitWalletId string
	Description   string
}
type WalletCreditReservedEvent struct {
	Amount      decimal.Decimal
	Description string
}
type WalletCreditReleasedEvent struct {
	Amount      decimal.Decimal
	Description string
}

type WalletLockedEvent struct {
	Description string
}
type WalletDeletedEvent struct {
	Description string
}
type WalletUnlockedEvent struct {
	Description string
}
type WalletBlacklistedEvent struct {
	Description string
}
type WalletUnBlacklistedEvent struct {
	Description string
}

func NewWalletCreatedEvent(aggregate es.Aggregate,
	amount decimal.Decimal,
	description string,
	userId string,
	accountId string,
	id string,
) (es.Event, error) {
	eventData := WalletCreatedEvent{
		Amount:      amount,
		AccountId:   accountId,
		UserId:      userId,
		ID:          id,
		Description: description,
	}

	event := es.NewBaseEvent(aggregate, WalletCreated)
	if err := event.SetJsonData(&eventData); err != nil {
		return es.Event{}, err
	}
	return event, nil
}

func NewWalletDebitEvent(aggregate es.Aggregate, creditWalletId string, amount decimal.Decimal, description string) (es.Event, error) {
	eventData := WalletDebitedEvent{
		Amount:         amount,
		Description:    description,
		CreditWalletId: creditWalletId,
	}
	event := es.NewBaseEvent(aggregate, WalletDebited)
	if err := event.SetJsonData(&eventData); err != nil {
		return es.Event{}, err
	}
	return event, nil
}

func NewWalletCreditEvent(aggregate es.Aggregate, debitWalletId string, amount decimal.Decimal, description string) (es.Event, error) {
	eventData := WalletCreditedEvent{
		Amount:        amount,
		Description:   description,
		DebitWalletId: debitWalletId,
	}
	event := es.NewBaseEvent(aggregate, WalletCredited)
	if err := event.SetJsonData(&eventData); err != nil {
		return es.Event{}, err
	}
	return event, nil
}
func NewWalletCreditReservedEvent(aggregate es.Aggregate, amount decimal.Decimal, description string) (es.Event, error) {
	eventData := WalletCreditReservedEvent{
		Amount:      amount,
		Description: description,
	}
	event := es.NewBaseEvent(aggregate, WalletCreditReserved)
	if err := event.SetJsonData(&eventData); err != nil {
		return es.Event{}, err
	}
	return event, nil
}
func NewWalletCreditReleasedEvent(aggregate es.Aggregate, amount decimal.Decimal, description string) (es.Event, error) {
	eventData := WalletCreditReleasedEvent{
		Amount:      amount,
		Description: description,
	}
	event := es.NewBaseEvent(aggregate, WalletCreditReleased)
	if err := event.SetJsonData(&eventData); err != nil {
		return es.Event{}, err
	}
	return event, nil
}
func NewWalletLockedEvent(aggregate es.Aggregate, description string) (es.Event, error) {
	eventData := WalletLockedEvent{
		Description: description,
	}
	event := es.NewBaseEvent(aggregate, WalletLocked)
	if err := event.SetJsonData(&eventData); err != nil {
		return es.Event{}, err
	}
	return event, nil
}
func NewWalletUnlockedEvent(aggregate es.Aggregate, description string) (es.Event, error) {
	eventData := WalletUnlockedEvent{
		Description: description,
	}
	event := es.NewBaseEvent(aggregate, WalletUnlocked)
	if err := event.SetJsonData(&eventData); err != nil {
		return es.Event{}, err
	}
	return event, nil
}
func NewWalletBlacklistedEvent(aggregate es.Aggregate, description string) (es.Event, error) {
	eventData := WalletBlacklistedEvent{
		Description: description,
	}
	event := es.NewBaseEvent(aggregate, WalletBlacklisted)
	if err := event.SetJsonData(&eventData); err != nil {
		return es.Event{}, err
	}
	return event, nil
}
func NewWalletUnBlacklistedEvent(aggregate es.Aggregate, description string) (es.Event, error) {
	eventData := WalletUnBlacklistedEvent{
		Description: description,
	}
	event := es.NewBaseEvent(aggregate, WalletUnBlacklisted)
	if err := event.SetJsonData(&eventData); err != nil {
		return es.Event{}, err
	}
	return event, nil
}
func NewWalletDeletedEvent(aggregate es.Aggregate, description string) (es.Event, error) {
	eventData := WalletDeletedEvent{
		Description: description,
	}
	event := es.NewBaseEvent(aggregate, WalletDeleted)
	if err := event.SetJsonData(&eventData); err != nil {
		return es.Event{}, err
	}
	return event, nil
}
