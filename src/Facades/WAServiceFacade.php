<?php

namespace Suainul\Waservice\Facades;

use Illuminate\Support\Facades\Facade;

class WAServiceFacade extends Facade
{
    /**
     * Get the registered name of the component.
     *
     * @return string
     */
    protected static function getFacadeAccessor()
    {
        return 'waservice';
    }
}