<?php

use Grpc\ChannelCredentials;

require_once __DIR__ . '/vendor/autoload.php';

//127.0.0.1:8002 服务端的端口
$client = new \Services\PersonServiceClient('127.0.0.1:8002', [
    'credentials' => ChannelCredentials::createInsecure()
]);

$oInfoRequest = new \Services\PersonInfoReq();
$oInfoRequest->setId(110);
list($response, $status) = $client->GetPersonInfo($oInfoRequest)->wait();
if ($status->code !== Grpc\STATUS_OK) {
    echo "ERROR: " . $status->code . ", " . $status->details . PHP_EOL;
    exit(1);
}

echo PHP_EOL . "========GetPersonInfo==========" . PHP_EOL;
echo sprintf("error = %s | msg = %s | token = %s | name = %s | login = %d",
    $response->getError(),
    $response->getMsg(),
    $response->getData()->getToken(),
    $response->getData()->getName(),
    $response->getData()->getLogin()
);

echo PHP_EOL . "========GetPersonList==========" . PHP_EOL;

$oListRequest = new \Services\PersonListReq();
$oListRequest->setName("wahaha");
list($response, $status) = $client->GetPersonList($oListRequest)->wait();
if ($status->code !== Grpc\STATUS_OK) {
    echo "ERROR: " . $status->code . ", " . $status->details . PHP_EOL;
    exit(1);
}

echo sprintf("error = %s | msg = %s ",
    $response->getError(),
    $response->getMsg()
) . PHP_EOL;

foreach ($response->getData() as $index => $data){
    echo sprintf("{$index}: token = %s | name = %s | login = %d",
        $data->getToken(),
        $data->getName(),
        $data->getLogin()
    ) . PHP_EOL;
}