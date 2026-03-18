package imsg

import "mayfly-go/pkg/i18n"

var Zh_CN = map[i18n.MsgId]string{
	LogKafkaSave:   "Kafka-保存",
	LogKafkaDelete: "Kafka-删除",
	LogKafkaRunCmd: "Kafka-执行命令",
	LogUpdateDocs:  "Kafka-更新文档",
	LogDelDocs:     "Kafka-删除文档",
	LogInsertDocs:  "Kafka-插入文档",

	ErrKafkaInfoExist: "该信息已存在",
}
