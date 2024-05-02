
CREATE TABLE "matcher" (
  "id" SERIAL PRIMARY KEY,
  "userId" Int NOT NULL,
  "userCatId" Int NOT NULL,
  "matchCatId" Int NOT NULL,
  "message" varchar(120) NOT NULL,
  "isApproved" bool DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "matcher" ADD FOREIGN KEY ("userId") REFERENCES "user" ("id");

ALTER TABLE "matcher" ADD FOREIGN KEY ("userCatId") REFERENCES "cat" ("id");

ALTER TABLE "matcher" ADD FOREIGN KEY ("matchCatId") REFERENCES "cat" ("id");
