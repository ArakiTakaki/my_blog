#!/bin/bash

# =============================================
#
# Auth test
#
# =============================================
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
curl -v  -X POST http://localhost:3000/api/auth/logout

# =============================================
#
# Post Test
#
# =============================================
# posts
curl -v http://localhost:3000/api/posts/
# 記事の投稿サンプル
curl -v -X POST \
  http://localhost:3000/api/posts/create \
  -H 'content-type: application/json' \
  -d '{
    "title": "サンプルテキスト",
    "status": "open",
    "comment_status": "sample_state",
    "password": "sample",
    "detail": {
      "content": "ssample_content"
    }
  }'

# INSERT INTO "posts" ("id","created_at","updated_at","deleted_at","title","excerpt","status","comment_status","password","user_id") VALUES ('1','2018-12-09 20:45:13','2018-12-09 20:45:13',NULL,'サンプルテキスト','サンプルテキスト','open','sample_state','sample','0')
# INSERT INTO "post_details" ("created_at","updated_at","deleted_at","post_id","content") VALUES ('2018-12-09 20:45:13','2018-12-09 20:45:13',NULL,'1','ssample_content')