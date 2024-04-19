// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"iot_device_simulation/internal/dao/internal"
)

// internalSubTopicDao is internal type for wrapping internal DAO implements.
type internalSubTopicDao = *internal.SubTopicDao

// subTopicDao is the data access object for table sub_topic.
// You can define custom methods on it to extend its functionality as you wish.
type subTopicDao struct {
	internalSubTopicDao
}

var (
	// SubTopic is globally public accessible object for table sub_topic operations.
	SubTopic = subTopicDao{
		internal.NewSubTopicDao(),
	}
)

// Fill with you ideas below.