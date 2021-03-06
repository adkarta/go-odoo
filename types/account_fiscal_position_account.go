package types

import (
	"time"
)

type AccountFiscalPositionAccount struct {
	AccountDestId Many2One  `xmlrpc:"account_dest_id"`
	AccountSrcId  Many2One  `xmlrpc:"account_src_id"`
	CreateDate    time.Time `xmlrpc:"create_date"`
	CreateUid     Many2One  `xmlrpc:"create_uid"`
	DisplayName   string    `xmlrpc:"display_name"`
	Id            int64     `xmlrpc:"id"`
	LastUpdate    time.Time `xmlrpc:"__last_update"`
	PositionId    Many2One  `xmlrpc:"position_id"`
	WriteDate     time.Time `xmlrpc:"write_date"`
	WriteUid      Many2One  `xmlrpc:"write_uid"`
}

type AccountFiscalPositionAccountNil struct {
	AccountDestId interface{} `xmlrpc:"account_dest_id"`
	AccountSrcId  interface{} `xmlrpc:"account_src_id"`
	CreateDate    interface{} `xmlrpc:"create_date"`
	CreateUid     interface{} `xmlrpc:"create_uid"`
	DisplayName   interface{} `xmlrpc:"display_name"`
	Id            interface{} `xmlrpc:"id"`
	LastUpdate    interface{} `xmlrpc:"__last_update"`
	PositionId    interface{} `xmlrpc:"position_id"`
	WriteDate     interface{} `xmlrpc:"write_date"`
	WriteUid      interface{} `xmlrpc:"write_uid"`
}

var AccountFiscalPositionAccountModel string = "account.fiscal.position.account"

type AccountFiscalPositionAccounts []AccountFiscalPositionAccount

type AccountFiscalPositionAccountsNil []AccountFiscalPositionAccountNil

func (s *AccountFiscalPositionAccount) NilableType_() interface{} {
	return &AccountFiscalPositionAccountNil{}
}

func (ns *AccountFiscalPositionAccountNil) Type_() interface{} {
	s := &AccountFiscalPositionAccount{}
	return load(ns, s)
}

func (s *AccountFiscalPositionAccounts) NilableType_() interface{} {
	return &AccountFiscalPositionAccountsNil{}
}

func (ns *AccountFiscalPositionAccountsNil) Type_() interface{} {
	s := &AccountFiscalPositionAccounts{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountFiscalPositionAccount))
	}
	return s
}
