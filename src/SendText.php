<?php

namespace Suainul\Waservice;

class SendText extends ApiRequest
{
    private $message,$target =[],$result;
    public function __construct()
    {
        parent::__construct();
    }
    public function send()
    {
        $this->result = $this->requestApi("POST","send/text",[
            "message" =>  $this->message,
            "target" =>  $this->target
        ]);
        return $this;
    }

    /**
     * Get the value of message
     */ 
    public function getMessage()
    {
        return $this->message;
    }

    /**
     * Set the value of message
     *
     * @return  self
     */ 
    public function setMessage($message)
    {
        $this->message = $message;

        return $this;
    }

    /**
     * Get the value of target
     */ 
    public function getTarget()
    {
        return $this->target;
    }

    /**
     * Set the value of target
     *
     * @return  self
     */ 
    public function setTarget($target =[])
    {
        $this->target = $target;

        return $this;
    }

    /**
     * Get the value of result
     */ 
    public function getResult()
    {
        return $this->result;
    }
}