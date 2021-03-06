<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: Person.proto

namespace Services;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * 定义list响应
 *
 * Generated from protobuf message <code>services.PersonListResp</code>
 */
class PersonListResp extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>int32 Error = 1;</code>
     */
    protected $Error = 0;
    /**
     * Generated from protobuf field <code>string Msg = 2;</code>
     */
    protected $Msg = '';
    /**
     * Generated from protobuf field <code>repeated .services.PersonRespData data = 3;</code>
     */
    private $data;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int $Error
     *     @type string $Msg
     *     @type \Services\PersonRespData[]|\Google\Protobuf\Internal\RepeatedField $data
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Person::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>int32 Error = 1;</code>
     * @return int
     */
    public function getError()
    {
        return $this->Error;
    }

    /**
     * Generated from protobuf field <code>int32 Error = 1;</code>
     * @param int $var
     * @return $this
     */
    public function setError($var)
    {
        GPBUtil::checkInt32($var);
        $this->Error = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string Msg = 2;</code>
     * @return string
     */
    public function getMsg()
    {
        return $this->Msg;
    }

    /**
     * Generated from protobuf field <code>string Msg = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setMsg($var)
    {
        GPBUtil::checkString($var, True);
        $this->Msg = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated .services.PersonRespData data = 3;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getData()
    {
        return $this->data;
    }

    /**
     * Generated from protobuf field <code>repeated .services.PersonRespData data = 3;</code>
     * @param \Services\PersonRespData[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setData($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Services\PersonRespData::class);
        $this->data = $arr;

        return $this;
    }

}

