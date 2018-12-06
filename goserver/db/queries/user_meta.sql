CREATE TABLE IF NOT EXISTS "user_meta"(
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "user_id" INTEGER,
  "key" TEXT NOT NULL INDEX,
  "value" TEXT,
  FOREIGN KEY(user_id) REFERENCES users(user_id)
);

CREATE INDEX "user_meta_key" ON "user_meta"("key");
