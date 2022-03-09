<?php

// autoload_static.php @generated by Composer

namespace Composer\Autoload;

class ComposerStaticInit71a990ededa575c3866d1d1f21651b31
{
    public static $prefixLengthsPsr4 = array (
        'S' => 
        array (
            'Suainul\\Waservice\\' => 18,
        ),
    );

    public static $prefixDirsPsr4 = array (
        'Suainul\\Waservice\\' => 
        array (
            0 => __DIR__ . '/../..' . '/src',
        ),
    );

    public static $classMap = array (
        'Composer\\InstalledVersions' => __DIR__ . '/..' . '/composer/InstalledVersions.php',
    );

    public static function getInitializer(ClassLoader $loader)
    {
        return \Closure::bind(function () use ($loader) {
            $loader->prefixLengthsPsr4 = ComposerStaticInit71a990ededa575c3866d1d1f21651b31::$prefixLengthsPsr4;
            $loader->prefixDirsPsr4 = ComposerStaticInit71a990ededa575c3866d1d1f21651b31::$prefixDirsPsr4;
            $loader->classMap = ComposerStaticInit71a990ededa575c3866d1d1f21651b31::$classMap;

        }, null, ClassLoader::class);
    }
}
