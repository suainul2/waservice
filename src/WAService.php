<?php

namespace Suainul\Waservice;
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
    public function Generate()
    {
        return new ApiRequest;
    }
}
