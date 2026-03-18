package imsg

import "mayfly-go/pkg/i18n"

var En = map[i18n.MsgId]string{
	LogKafkaSave:   "Kafka - Save",
	LogKafkaDelete: "Kafka-Delete",
	LogKafkaRunCmd: "Kafka - Run Cmd",
	LogUpdateDocs:  "Kafka - Update Documents",
	LogDelDocs:     "Kafka - Delete Documents",
	LogInsertDocs:  "Kafka - Insert Documents",

	ErrKafkaInfoExist: "that information already exists",
}
