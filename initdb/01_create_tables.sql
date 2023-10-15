CREATE TABLE "user" (
  "id" UUID PRIMARY KEY,
  "name" VARCHAR(100)
);

CREATE TABLE "post" (
  "id" UUID PRIMARY KEY,
  "content" VARCHAR(255),
  "user_id" UUID
);

ALTER TABLE "post" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

