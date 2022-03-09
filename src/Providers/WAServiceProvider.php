<?php

namespace Suainul\Waservice\Providers;

use Illuminate\Support\ServiceProvider;
use Suainul\Waservice\WAService;

class WAServiceProvider extends ServiceProvider
{
    /**
     * Perform post-registration booting of services.
     *
     * @return void
     */
    public function boot()
    {
        // Publishing is only necessary when using the CLI.
        if ($this->app->runningInConsole()) {
            $this->bootForConsole();
        }
    }

    /**
     * Register any package services.
     *
     * @return void
     */
    public function register()
    {
        $this->mergeConfigFrom(__DIR__ . '/../../config/waservice.php', 'waservice');

        // Register the service the package provides.
        $this->app->singleton('waservice', function ($app) {
            return new WAService;
        });
    }
    
    /**
     * Console-specific booting.
     *
     * @return void
     */
    protected function bootForConsole()
    {
        // Publishing the configuration file.
        $this->publishes([
            __DIR__ . '/../../config/waservice.php' => config_path('waservice.php'),
        ], 'waservice.config');
    }
}