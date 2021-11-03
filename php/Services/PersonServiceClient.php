<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Services;

/**
 * 定义服务
 */
class PersonServiceClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * @param \Services\PersonListReq $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function GetPersonList(\Services\PersonListReq $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/services.PersonService/GetPersonList',
        $argument,
        ['\Services\PersonListResp', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \Services\PersonInfoReq $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function GetPersonInfo(\Services\PersonInfoReq $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/services.PersonService/GetPersonInfo',
        $argument,
        ['\Services\PersonInfoResp', 'decode'],
        $metadata, $options);
    }

}
