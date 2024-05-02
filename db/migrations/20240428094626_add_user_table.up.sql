CREATE TABLE "user" (
  "id" SERIAL PRIMARY KEY,
  "email" varchar(255) UNIQUE NOT NULL,
  "name" varchar(255) NOT NULL CHECK (length("name") >= 5),
  "password" varchar(255) NOT NULL,
  "accessToken" varchar(255) UNIQUE,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);