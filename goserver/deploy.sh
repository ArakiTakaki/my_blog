#!/bin/bash
source .env
scp -r -i $KEY "${PWD}/main" $SERVER