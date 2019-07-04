package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleKey is the name of the module
	ModuleName = "record"
	// StoreKey is the store key string for record
	StoreKey = ModuleName
	// RouterKey is the message route for record
	RouterKey = ModuleName
	// QuerierRoute is the querier route for record
	QuerierRoute = ModuleName
	// Parameter store default namestore
	DefaultParamspace = ModuleName
)
const (
	DefaultCodespace sdk.CodespaceType = ModuleName
)

var (
	RecordMaxId        uint64 = 999999999999
	RecordMinId        uint64 = 100000000000
)

const (
	IDPreStr = "record"
	Custom   = "custom"
)
const (
	QueryParams    = "params"
	QueryRecords   = "list"
	QueryRecord    = "query"
	QuerySearch    = "search"
)

const (
	TypeMsgRecord                  = "record"
)
const (
	CodeInvalidGenesis   sdk.CodeType = 102
	HashLength	                      = 64
	NameMinLength                     = 3
	NameMaxLength                     = 32
	AuthorMinLength                   = 3
	AuthorMaxLength                   = 64
	TypeMinLength                     = 3
	TypeMaxLength                     = 32
	RecordNoMinLength                 = 3
	RecordNoMaxLength                 = 32
)
