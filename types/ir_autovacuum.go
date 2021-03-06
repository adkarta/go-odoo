package types

import (
	"time"
)

type IrAutovacuum struct {
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
}

type IrAutovacuumNil struct {
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
}

var IrAutovacuumModel string = "ir.autovacuum"

type IrAutovacuums []IrAutovacuum

type IrAutovacuumsNil []IrAutovacuumNil

func (s *IrAutovacuum) NilableType_() interface{} {
	return &IrAutovacuumNil{}
}

func (ns *IrAutovacuumNil) Type_() interface{} {
	s := &IrAutovacuum{}
	return load(ns, s)
}

func (s *IrAutovacuums) NilableType_() interface{} {
	return &IrAutovacuumsNil{}
}

func (ns *IrAutovacuumsNil) Type_() interface{} {
	s := &IrAutovacuums{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrAutovacuum))
	}
	return s
}
