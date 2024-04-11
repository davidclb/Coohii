CREATE TABLE "users" (
  id SERIAL PRIMARY KEY,
  username varchar NOT NULL, 
  email varchar NOT NULL,
  password varchar NOT NULL,
  join_date timestamptz DEFAULT (now())
);

CREATE TABLE "user_settings" (
  "id" integer PRIMARY KEY,
  "user_id" integer NOT NULL,
  "created_on" timestamptz DEFAULT (now()),
  "updated_on" timestamp NOT NULL
);

CREATE TABLE "review" (
  "id" integer PRIMARY KEY,
  "user_id" integer NOT NULL,
  "last_review" timestamp NOT NULL,
  "carac_ucs_id" integer NOT NULL,
  "status" varchar NOT NULL,
  "created_on" timestamptz NOT NULL DEFAULT (now()),
  "expire_date" timestamp NOT NULL,
  "total_review" integer NOT NULL,
  "failure_count" integer NOT NULL,
  "success_count" integer NOT NULL,
  "study_count" integer NOT NULL
);

CREATE TABLE "characters" (
  "ucs_id" integer PRIMARY KEY,
  "carac_simpl" varchar NOT NULL,
  "carac_trad" varchar NOT NULL,
  "pinyin" varchar NOT NULL,
  "meaning" text NOT NULL,
  "type" varchar NOT NULL,
  "index_heisig_simpl" integer NOT NULL,
  "index_heisig_trad" integer NOT NULL
);

CREATE TABLE "learned_caractere" (
  "id" integer PRIMARY KEY,
  "user_id" integer NOT NULL,
  "carac_ucs_id" integer NOT NULL
);

CREATE TABLE "custom_keywords" (
  "id" integer PRIMARY KEY,
  "user_id" integer NOT NULL,
  "carac_ucs_id" integer NOT NULL,
  "keyword" varchar NOT NULL,
  "created_on" timestamp NOT NULL,
  "update_on" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "stories" (
  "id" integer PRIMARY KEY,
  "user_id" integer NOT NULL,
  "carac_ucs_id" integer NOT NULL,
  "created_on" timestamptz NOT NULL DEFAULT (now()),
  "update_on" timestamp NOT NULL,
  "body" text NOT NULL,
  "public" integer NOT NULL
);

CREATE TABLE "stories_shared" (
  "id" integer PRIMARY KEY,
  "user_id" integer NOT NULL,
  "carac_ucs_id" integer NOT NULL,
  "created_on" timestamptz NOT NULL DEFAULT (now()),
  "update_on" timestamp NOT NULL,
  "stars" smallint NOT NULL,
  "reports" smallint NOT NULL
);

CREATE TABLE "story_votes" (
  "id" integer PRIMARY KEY,
  "author_id" integer NOT NULL,
  "voter_id" integer NOT NULL,
  "carac_ucs_id" integer NOT NULL,
  "created_on" timestamptz NOT NULL DEFAULT (now()),
  "vote" integer NOT NULL
);

CREATE TABLE "news" (
  "id" integer PRIMARY KEY,
  "subject" varchar NOT NULL,
  "created_on" timestamptz NOT NULL DEFAULT (now()),
  "update_on" timestamp NOT NULL,
  "body" text NOT NULL
);

CREATE TABLE "hanzi_lists" (
  "id" integer PRIMARY KEY,
  "subject" varchar NOT NULL,
  "created_on" timestamptz NOT NULL DEFAULT (now()),
  "update_on" timestamp NOT NULL,
  "body" text NOT NULL
);

CREATE TABLE "word_lists" (
  "id" integer PRIMARY KEY,
  "subject" varchar NOT NULL,
  "created_on" timestamptz NOT NULL DEFAULT (now()),
  "update_on" timestamp NOT NULL,
  "body" text NOT NULL
);



CREATE INDEX ON "user_settings" ("user_id");

CREATE INDEX ON "review" ("user_id");

CREATE INDEX ON "review" ("user_id", "carac_ucs_id");

CREATE INDEX ON "learned_caractere" ("user_id");

CREATE INDEX ON "learned_caractere" ("user_id", "carac_ucs_id");

CREATE INDEX ON "custom_keywords" ("user_id");

CREATE INDEX ON "custom_keywords" ("user_id", "carac_ucs_id");

CREATE INDEX ON "stories" ("user_id");

CREATE INDEX ON "stories" ("user_id", "carac_ucs_id");

CREATE INDEX ON "stories_shared" ("user_id");

CREATE INDEX ON "stories_shared" ("carac_ucs_id");

CREATE INDEX ON "stories_shared" ("user_id", "carac_ucs_id");

CREATE INDEX ON "story_votes" ("voter_id");

CREATE INDEX ON "story_votes" ("carac_ucs_id");

CREATE INDEX ON "story_votes" ("voter_id", "carac_ucs_id");

COMMENT ON COLUMN "stories"."body" IS 'Content of the post';

COMMENT ON COLUMN "news"."body" IS 'Content of the post';

ALTER TABLE "user_settings" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "review" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "review" ADD FOREIGN KEY ("carac_ucs_id") REFERENCES "characters" ("ucs_id");

ALTER TABLE "learned_caractere" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "learned_caractere" ADD FOREIGN KEY ("carac_ucs_id") REFERENCES "characters" ("ucs_id");

ALTER TABLE "custom_keywords" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "custom_keywords" ADD FOREIGN KEY ("carac_ucs_id") REFERENCES "characters" ("ucs_id");

ALTER TABLE "stories" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "stories" ADD FOREIGN KEY ("carac_ucs_id") REFERENCES "characters" ("ucs_id");

ALTER TABLE "stories_shared" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "stories_shared" ADD FOREIGN KEY ("carac_ucs_id") REFERENCES "characters" ("ucs_id");

ALTER TABLE "story_votes" ADD FOREIGN KEY ("voter_id") REFERENCES "users" ("id");

ALTER TABLE "story_votes" ADD FOREIGN KEY ("carac_ucs_id") REFERENCES "users" ("id");
