CREATE TYPE CatRace AS ENUM (
  'Persian',
  'Maine Coon',
  'Siamese',
  'Ragdoll',
  'Bengal',
  'Sphynx',
  'British Shorthair',
  'Abyssinian',
  'Scottish Fold',
  'Birman'
);

CREATE TYPE Gender AS ENUM (
  'male',
  'female'
);



CREATE TABLE "cat" (
  "id" SERIAL PRIMARY KEY,
  "userId" Int,
  "name" varchar(30) NOT NULL,
  "race" CatRace NOT NULL,
  "sex" Gender NOT NULL,
  "ageInMonth" Int NOT NULL CHECK ("ageInMonth" >= 1 AND "ageInMonth" <= 1000),
  "description" varchar(200) NOT NULL,
  "hasMatched" bool DEFAULT false,
  "imageUrls" text[] NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "cat" ADD FOREIGN KEY ("userId") REFERENCES "user" ("id");
