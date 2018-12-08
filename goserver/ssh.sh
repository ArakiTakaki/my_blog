#!/bin/bash
source .env
ssh -i $KEY $SERVER
scp -r -i $KEY "${PWD}/dist/server" $SERVER