<?php

require_once __DIR__ . '/vendor/autoload.php';
require_once __DIR__ . '/GPBMetadata/User.php';
require_once __DIR__ . '/User/UserFilter.php';
require_once __DIR__ . '/User/UserResponse.php';
require_once __DIR__ . '/User/UserRequest.php';
require_once __DIR__ . '/User/UserClient.php';

$client = new \User\UserClient('localhost:50051', [
    'credentials' => Grpc\ChannelCredentials::createInsecure()
]);

// Create a new user.
$user = new \User\UserRequest;
$user->setId(1);
$user->setName('Fredrik');
$user->setEmail('fredrik@example.com');
$client->CreateUser($user)->wait();

// Find the user that we just created.
$filter = new \User\UserFilter;
$filter->setName('Fredrik');
$response = $client->GetUsers($filter);

foreach ($response->responses() as $res) {
    echo sprintf("Hello %s\n", $res->getName());
    break;
}