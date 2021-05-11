-- -------------------------------------------------------------
-- TablePlus 3.12.6(366)
--
-- https://tableplus.com/
--
-- Database: go_smallsite_bookings
-- Generation Time: 2564-05-11 18:31:37.7330
-- -------------------------------------------------------------


DROP TABLE IF EXISTS "public"."users";
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS users_id_seq;

-- Table Definition
CREATE TABLE "public"."users" (
    "id" int4 NOT NULL DEFAULT nextval('users_id_seq'::regclass),
    "first_name" varchar(255) NOT NULL DEFAULT ''::character varying,
    "last_name" varchar(255) NOT NULL DEFAULT ''::character varying,
    "email" varchar(255) NOT NULL,
    "password" varchar(60) NOT NULL,
    "access_level" int4 NOT NULL DEFAULT 1,
    "status" varchar(255) NOT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    PRIMARY KEY ("id")
);

