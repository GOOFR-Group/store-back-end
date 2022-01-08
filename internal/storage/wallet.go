package storage

import (
	"github.com/gocraft/dbr/v2"
	"github.com/google/uuid"
)

const WalletTable = "wallet"

// tgcon - used to generate constants for each field's tag
type Wallet struct {
	ID       uuid.UUID `db:"id"`
	IDClient uuid.UUID `db:"id_client"`
	Balance  float64   `db:"balance"`
	Coin     rune      `db:"coin"`
}

func CreateWallet(t Transaction, model Wallet) error {
	_, err := t.InsertInto(WalletTable).
		Columns(WalletIDDb, WalletIDClientDb, WalletBalanceDb, WalletCoinDb).
		Record(model).
		Exec()

	return err
}

func ReadWalletByID(t Transaction, id uuid.UUID) (object Wallet, ok bool, err error) {
	err = t.Select("*").
		From(WalletTable).
		Where(WalletIDDb+" = ?", id).
		LoadOne(&object)

	switch err {
	case nil:
		ok = true
	case dbr.ErrNotFound:
		err = nil
	}
	return
}

func ReadWalletByClientID(t Transaction, id uuid.UUID) (object Wallet, ok bool, err error) {
	err = t.Select("*").
		From(WalletTable).
		Where(WalletIDClientDb+" = ?", id).
		LoadOne(&object)

	switch err {
	case nil:
		ok = true
	case dbr.ErrNotFound:
		err = nil
	}
	return
}

func UpdateWalletByID(t Transaction, model Wallet) error {
	_, err := t.Update(WalletTable).
		SetMap(map[string]interface{}{
			WalletBalanceDb: model.Balance,
			WalletCoinDb:    model.Coin,
		}).
		Where(WalletIDDb+" = ?", model.ID).
		Exec()

	return err
}
