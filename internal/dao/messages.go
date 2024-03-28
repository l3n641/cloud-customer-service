// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"cloudCustomerService/internal/dao/internal"
)

// internalMessagesDao is internal type for wrapping internal DAO implements.
type internalMessagesDao = *internal.MessagesDao

// messagesDao is the data access object for table messages.
// You can define custom methods on it to extend its functionality as you wish.
type messagesDao struct {
	internalMessagesDao
}

var (
	// Messages is globally public accessible object for table messages operations.
	Messages = messagesDao{
		internal.NewMessagesDao(),
	}
)

// Fill with you ideas below.
