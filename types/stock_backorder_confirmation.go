package types

import (
	"time"
)

type StockBackorderConfirmation struct {
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	PickIds     []int64   `xmlrpc:"pick_ids"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type StockBackorderConfirmationNil struct {
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	PickIds     interface{} `xmlrpc:"pick_ids"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var StockBackorderConfirmationModel string = "stock.backorder.confirmation"

type StockBackorderConfirmations []StockBackorderConfirmation

type StockBackorderConfirmationsNil []StockBackorderConfirmationNil

func (s *StockBackorderConfirmation) NilableType_() interface{} {
	return &StockBackorderConfirmationNil{}
}

func (ns *StockBackorderConfirmationNil) Type_() interface{} {
	s := &StockBackorderConfirmation{}
	return load(ns, s)
}

func (s *StockBackorderConfirmations) NilableType_() interface{} {
	return &StockBackorderConfirmationsNil{}
}

func (ns *StockBackorderConfirmationsNil) Type_() interface{} {
	s := &StockBackorderConfirmations{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*StockBackorderConfirmation))
	}
	return s
}
