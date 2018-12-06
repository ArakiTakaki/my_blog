CREATE TABLE IF NOT EXISTS "users"(
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "name" TEXT NOT NULL,
  "display_name" TEXT NOT NULL,
  "email" TEXT NOT NULL,
  "login" TEXT NOT NULL,
  "password" TEXT NOT NULL
);

CREATE INDEX "user_login" ON "users"("login");
