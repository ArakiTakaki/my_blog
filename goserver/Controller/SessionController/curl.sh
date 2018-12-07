#!/bin/bash
# login
curl -v -X POST \
  http://localhost:3000/api/auth/login \
  -H 'content-type: application/json' \
  -d '{ "user": "takaki", "password": "araki_admin" }'

# register
curl -v -X POST \
  http://localhost:3000/api/auth/register \
  -H 'content-type: application/json' \
  -d '{ "login": "takaki", "password": "araki_admin" }'