CREATE TABLE "users" (
  "id" serial PRIMARY KEY,
  "username" varchar NOT NULL,
  "email" varchar NOT NULL,
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
  "id" integer PRIMARY KEY,
  "user_id" integer NOT NULL,
  "last_review" timestamp NOT NULL,
  "carac_id" integer NOT NULL,
  "interval" integer,
  "created_on" timestamptz NOT NULL DEFAULT (now()),
  "nextReview" timestamp NOT NULL,
  "times_reviewed" integer NOT NULL,
  "mastery_level" integer NOT NULL,
  "consecutive_correct" integer NOT NULL,
  "easiness_correct" float64 NOT NULL
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

CREATE TABLE "radical" (
  "id" integer PRIMARY KEY,
  "strokes_number" integer NOT NULL,
  "radical" varchar NOT NULL,
  "pinyin" varchar NOT NULL,
  "meaning" varchar NOT NULL
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

COMMENT ON COLUMN "hanzi_lists"."svg" IS 'Cover image of the word list ';

COMMENT ON COLUMN "word_lists"."svg" IS 'Cover image of the word list ';

ALTER TABLE "review" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "review" ADD FOREIGN KEY ("carac_id") REFERENCES "character" ("id");
