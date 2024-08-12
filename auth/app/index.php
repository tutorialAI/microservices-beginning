<?php

$conf = new RdKafka\Conf();
$conf->set('group.id', 'php-consumer-group');
$conf->set('metadata.broker.list', 'kafka:9092');


$consumer = new \RdKafka\Consumer($conf);
$topic = $consumer->newTopic('test', $topicConf);
$queue = $consumer->newQueue();

$topic->consumeQueueStart(0, RD_KAFKA_OFFSET_BEGINNING, $queue);
$topic->consumeQueueStart(1, RD_KAFKA_OFFSET_BEGINNING, $queue);
$topic->consumeQueueStart(2, RD_KAFKA_OFFSET_BEGINNING, $queue);
do {
    $message = $queue->consume(1000);

    if ($message === null) {
        continue;
    } elseif ($message->err === RD_KAFKA_RESP_ERR_NO_ERROR) {
        // process your message here
        var_export($message);
    } elseif ($message->err === RD_KAFKA_RESP_ERR__PARTITION_EOF) {
        echo "error";
    } elseif ($message->err === RD_KAFKA_RESP_ERR__TIMED_OUT) {
        echo "timeout";
    } else {
        // handle other errors
        $topic->consumeStop(0);
        $topic->consumeStop(1);
        $topic->consumeStop(2);
        throw new \Exception($message->errstr(), $message->err);
    }

    // trigger callback queues
    $consumer->poll(1);
} while (true);