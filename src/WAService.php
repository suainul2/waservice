<?php

namespace Suainul\Waservice;

use GuzzleHttp\Client;

class WAService
{
    public function SendText()
    {
        return new SendText;
    }
    public function SendImage()
    {
        return new SendImage;
    }
}
