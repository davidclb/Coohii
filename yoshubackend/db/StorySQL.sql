CREATE TABLE "users" (
  "id" serial PRIMARY KEY,
  "username" varchar NOT NULL,
  "email" varchar NOT NULL,
  "role" varchar NOT NULL,
  "password" varchar NOT NULL,
  "join_date" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "session" (
  "id" int,
  "user_id" int,
  "created_on" timestamptz DEFAULT (now()),
  "completed_on" timestamp NOT NULL,
  "duration" int,
  "completed" bool
);

CREATE TABLE "review" (
  "id" integer,
  "user_id" integer NOT NULL,
  "last_review" timestamp NOT NULL,
  "carac_id" integer NOT NULL,
  "interval" integer,
  "created_on" timestamptz NOT NULL DEFAULT (now()),
  "nextReview" timestamp NOT NULL,
  "times_reviewed" integer NOT NULL,
  "mastery_level" integer NOT NULL,
  "consecutive_correct" integer NOT NULL,
  "easiness_correct" float64 NOT NULL,
  PRIMARY KEY ("id", "interval")
);

CREATE TABLE "character" (
  "id" integer PRIMARY KEY,
  "carac_simpl" varchar NOT NULL,
  "carac_trad" varchar NOT NULL,
  "pinyin" varchar NOT NULL,
  "meaning" text NOT NULL,
  "category" varchar,
  "carac_antonym" varchar,
  "carac_similar" varchar,
  "radical_list" varchar
);

CREATE TABLE "word" (
  "id" integer PRIMARY KEY,
  "carac_simpl_list" varchar NOT NULL,
  "carac_trad_list" varchar NOT NULL,
  "pinyin" varchar NOT NULL,
  "meaning" text NOT NULL,
  "category" varchar,
  "word_antonym" varchar,
  "word_similar" varchar,
  "sentences_list" varchar
);

CREATE TABLE "radicals" (
  "id" integer PRIMARY KEY,
  "strokes_number" integer NOT NULL,
  "radical" varchar NOT NULL,
  "pinyin" varchar NOT NULL,
  "meaning" varchar NOT NULL
);

CREATE TABLE "stories" (
  "id" integer PRIMARY KEY,
  "user_id" integer NOT NULL,
  "carac_ucs_id" integer NOT NULL,
  "created_on" timestamptz NOT NULL DEFAULT (now()),
  "update_on" timestamp NOT NULL,
  "body" text NOT NULL,
  "public" tinyint NOT NULL
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
  "title" varchar NOT NULL,
  "svg" text NOT NULL
);

CREATE TABLE "word_lists" (
  "id" integer PRIMARY KEY,
  "title" varchar NOT NULL,
  "svg" text NOT NULL
);

CREATE INDEX ON "review" ("user_id");

CREATE INDEX ON "review" ("carac_id");

CREATE INDEX ON "review" ("user_id", "carac_id");

CREATE INDEX ON "stories" ("user_id");

CREATE INDEX ON "stories" ("carac_ucs_id");

CREATE INDEX ON "stories" ("user_id", "carac_ucs_id");

CREATE INDEX ON "stories_shared" ("user_id");

CREATE INDEX ON "stories_shared" ("carac_ucs_id");

CREATE INDEX ON "stories_shared" ("user_id", "carac_ucs_id");

CREATE INDEX ON "story_votes" ("voter_id");

CREATE INDEX ON "story_votes" ("carac_ucs_id");

CREATE INDEX ON "story_votes" ("voter_id", "carac_ucs_id");

COMMENT ON COLUMN "stories"."body" IS 'Content of the post';

COMMENT ON COLUMN "news"."body" IS 'Content of the post';

COMMENT ON COLUMN "hanzi_lists"."svg" IS 'Cover image of the word list ';

COMMENT ON COLUMN "word_lists"."svg" IS 'Cover image of the word list ';

ALTER TABLE "review" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "review" ADD FOREIGN KEY ("carac_id") REFERENCES "character" ("id");

ALTER TABLE "stories" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "stories" ADD FOREIGN KEY ("carac_ucs_id") REFERENCES "character" ("id");

ALTER TABLE "stories_shared" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "stories_shared" ADD FOREIGN KEY ("carac_ucs_id") REFERENCES "character" ("id");

ALTER TABLE "story_votes" ADD FOREIGN KEY ("voter_id") REFERENCES "users" ("id");

ALTER TABLE "story_votes" ADD FOREIGN KEY ("carac_ucs_id") REFERENCES "users" ("id");
