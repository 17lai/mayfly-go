import Api from '@/common/Api';

export const mqApi = {
    kafkaList: Api.newGet('/mq/kafka'),
    KafkaTestConn: Api.newPost('/mq/kafka/test-conn'),
    kafkaSave: Api.newPost('/mq/kafka'),
    kafkaDel: Api.newDelete('/mq/kafka/{id}'),
    kafkaGetPwd: Api.newGet('/mq/kafka/{id}/pwd'),
    kafkaTopicList: Api.newGet('/mq/kafka/{id}/getTopics'),
    kafkaTopicCreate: Api.newPost('/mq/kafka/{id}/createTopic'),
    kafkaTopicInfo: Api.newGet('/mq/kafka/{id}/{topic}/getTopicConfig'),
    kafkaTopicDelete: Api.newDelete('/mq/kafka/{id}/{topic}/deleteTopic'),
    kafkaTopicCreatePartitions: Api.newPost('/mq/kafka/{id}/createPartitions'),
    kafkaTopicProduce: Api.newPost('/mq/kafka/{id}/{topic}/produce'),
    kafkaTopicConsume: Api.newPost('/mq/kafka/{id}/{topic}/consume'),
    kafkaTopicBrokers: Api.newGet('/mq/kafka/{id}/getBrokers'),
    kafkaTopicBrokerConfig: Api.newGet('/mq/kafka/{id}/getBrokerConfig/{brokerId}'),
    kafkaGetGroups: Api.newGet('/mq/kafka/{id}/getGroups'),
    kafkaDeleteGroup: Api.newDelete('/mq/kafka/{id}/deleteGroup/{group}'),
    kafkaGetGroupMembers: Api.newGet('/mq/kafka/{id}/getGroupMembers/{group}'),
};
