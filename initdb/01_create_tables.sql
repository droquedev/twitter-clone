CREATE TABLE "users" (
  "id" UUID PRIMARY KEY,
  "username" VARCHAR(20)
);

CREATE TABLE "posts" (
  "id" UUID PRIMARY KEY,
  "content" VARCHAR(256),
  "user_id" UUID
);

ALTER TABLE "posts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

