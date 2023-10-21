CREATE TABLE "users" (
  "id" UUID PRIMARY KEY,
  "username" VARCHAR(100),
  "password" CHAR(60)
);

CREATE TABLE "posts" (
  "id" UUID PRIMARY KEY,
  "content" VARCHAR(256),
  "user_id" UUID
);

CREATE TABLE "comments" (
  "id" UUID PRIMARY KEY,
  "content" VARCHAR(256),
  "user_id" UUID,
  "post_id" UUID
);

ALTER TABLE "posts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");