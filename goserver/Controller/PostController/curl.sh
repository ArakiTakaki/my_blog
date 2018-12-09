#!/bin/bash
# create
curl -v -X POST \
  http://localhost:3000/api/posts/create \
  -H 'content-type: application/json' \
  -d '{
    "title": "サンプルテキスト",
    "excerpt": "サンプルテキスト",
    "status": "open",
    "comment_status": "sample_state",
    "password": "sample",
    "detail": {
      "content": "ssample_content"
    }
  }'

# INSERT INTO "posts" ("id","created_at","updated_at","deleted_at","title","excerpt","status","comment_status","password","user_id") VALUES ('1','2018-12-09 20:45:13','2018-12-09 20:45:13',NULL,'サンプルテキスト','サンプルテキスト','open','sample_state','sample','0')
# INSERT INTO "post_details" ("created_at","updated_at","deleted_at","post_id","content") VALUES ('2018-12-09 20:45:13','2018-12-09 20:45:13',NULL,'1','ssample_content')