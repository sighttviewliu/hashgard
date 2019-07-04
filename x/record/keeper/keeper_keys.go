package keeper

import (
	"fmt"
	"github.com/hashgard/hashgard/x/record/types"
)

// Key for getting next available recordID from the store
var (
	KeyDelimiter   = ":"
	KeyNextRecordID = []byte("newRecordID")
)

// Key for getting a specific record content hash from the store
func KeyRecord(recordHash string) []byte {
	return []byte(fmt.Sprintf("hash:%s", recordHash))
}

// Key for getting records by a specific address from the store
func KeyAddressRecords(addr string) []byte {
	return []byte(fmt.Sprintf("address:%s", addr))
}

// Key for getting records by a specific author from the store
func KeyAuthorRecords(author string) []byte {
	return []byte(fmt.Sprintf("author:%s", author))
}

func KeyRecordId(id string) []byte {
	return []byte(fmt.Sprintf("id:%s", id))
}

func KeyRecordIdStr(seq uint64) string {
	return fmt.Sprintf("%s%x", types.IDPreStr, seq)
}
