#!/bin/bash
curl -v -X POST \
  http://localhost:3000/api/auth/login \
  -H 'content-type: application/json' \
  -d '{ "user": "hogehoge", "password": "123" }'