<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: Person.proto

namespace Services;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>services.PersonRespData</code>
 */
class PersonRespData extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string token = 1;</code>
     */
    protected $token = '';
    /**
     * Generated from protobuf field <code>string Name = 2;</code>
     */
    protected $Name = '';
    /**
     * Generated from protobuf field <code>int32 login = 3;</code>
     */
    protected $login = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $token
     *     @type string $Name
     *     @type int $login
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Person::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string token = 1;</code>
     * @return string
     */
    public function getToken()
    {
        return $this->token;
    }

    /**
     * Generated from protobuf field <code>string token = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setToken($var)
    {
        GPBUtil::checkString($var, True);
        $this->token = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string Name = 2;</code>
     * @return string
     */
    public function getName()
    {
        return $this->Name;
    }

    /**
     * Generated from protobuf field <code>string Name = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setName($var)
    {
        GPBUtil::checkString($var, True);
        $this->Name = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 login = 3;</code>
     * @return int
     */
    public function getLogin()
    {
        return $this->login;
    }

    /**
     * Generated from protobuf field <code>int32 login = 3;</code>
     * @param int $var
     * @return $this
     */
    public function setLogin($var)
    {
        GPBUtil::checkInt32($var);
        $this->login = $var;

        return $this;
    }

}
