#!/bin/bash
source .env
scp -r -i $KEY "${PWD}/config" $SERVER
scp -r -i $KEY "${PWD}/contents" $SERVER
scp -r -i $KEY "${PWD}/controller" $SERVER
scp -r -i $KEY "${PWD}/middleware" $SERVER
scp -r -i $KEY "${PWD}/public" $SERVER
scp -r -i $KEY "${PWD}/routes" $SERVER
scp -r -i $KEY "${PWD}/lib" $SERVER
scp -r -i $KEY "${PWD}/websocket" $SERVER
scp -i $KEY "${PWD}/main.js" $SERVER
scp -i $KEY "${PWD}/package.json" $SERVER
