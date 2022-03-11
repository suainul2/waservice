<?php

namespace Suainul\Waservice;

class SendImage extends ApiRequest
{
    private $message,$target =[],$result,$imgUrl="";
    public function __construct()
    {
        parent::__construct();
    }
    public function send()
    {
        $this->result = $this->requestApi("POST","send/image",[
            "message" =>  $this->message,
            "target" =>  $this->target,
            "imgUrl" => $this->imgUrl
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

    /**
     * Get the value of imgUrl
     */ 
    public function getImgUrl()
    {
        return $this->imgUrl;
    }

    /**
     * Set the value of imgUrl
     *
     * @return  self
     */ 
    public function setImgUrl($imgUrl)
    {
        $this->imgUrl = $imgUrl;

        return $this;
    }
}