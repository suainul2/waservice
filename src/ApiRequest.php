<?php

namespace Suainul\Waservice;

use GuzzleHttp\Client;

class ApiRequest
{
    private $url,$api_key;

    public function __construct()
    {
        $this->url = config('waservice.url', '');
        $this->api_key = config('waservice.api_key', '');
    }
    public function connect()
    {
        return $this->requestApi("GET","api/connect");
    }
    public function getLive()
    {
        return $this->requestApi("GET","api/live");
    }
    protected function requestApi($method,$endpoint,$params = [])
    {
        $client = new Client();
        try {
            $response = $client->request($method, $this->url . $endpoint, [
                'headers' => [
                    'Authorization' =>  'Bearer '.$this->api_key,
                    'Content-Type' => 'application/json'
                ],
                'connect_timeout' => 5,
                'body' => $params
            ]);

            if ($response->getStatusCode() === 200) {
                return $response->getBody();
            }
        } catch (\Exception $exception){
            return $exception->getMessage();
        }
    }
}
