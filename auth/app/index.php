<?php
    echo 'auth 1';

    $producer = new \RdKafka\Producer();
    //$producer->setLogLevel(LOG_DEBUG);

    if ($producer->addBrokers("kafka:9093") < 1) {
        echo "Error to adding brokers";
        exit();
    }


    $topic = $producer->newTopic("test");

    
    $topic->produce(RD_KAFKA_PARTITION_UA, 0, "Message 10");
    
   

    // for ($flushRetries = 0; $flushRetries < 10; $flushRetries++) {
    //     $result = $producer->flush(10000);
    //     if (RD_KAFKA_RESP_ERR_NO_ERROR === $result) {
    //         break;
    //     }
    // }
    
    // if (RD_KAFKA_RESP_ERR_NO_ERROR !== $result) {
    //     throw new \RuntimeException('Was unable to flush, messages might be lost!');
    // }

    echo PHP_EOL . "OK";