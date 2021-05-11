-- -------------------------------------------------------------
-- TablePlus 3.12.6(366)
--
-- https://tableplus.com/
--
-- Database: go_smallsite_bookings
-- Generation Time: 2564-05-11 18:30:29.7050
-- -------------------------------------------------------------


DROP TABLE IF EXISTS "public"."reservations";
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS reservations_id_seq;

-- Table Definition
CREATE TABLE "public"."reservations" (
    "id" int4 NOT NULL DEFAULT nextval('reservations_id_seq'::regclass),
    "first_name" varchar(255) NOT NULL DEFAULT ''::character varying,
    "last_name" varchar(255) NOT NULL DEFAULT ''::character varying,
    "email" varchar(255) NOT NULL,
    "phone" varchar(255) NOT NULL DEFAULT ''::character varying,
    "room_id" int4 NOT NULL,
    "status" varchar(255) NOT NULL DEFAULT 'available'::character varying,
    "start_date" date NOT NULL,
    "end_date" date NOT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    CONSTRAINT "reservations_rooms_id_fk" FOREIGN KEY ("room_id") REFERENCES "public"."rooms"("id") ON DELETE CASCADE ON UPDATE CASCADE,
    PRIMARY KEY ("id")
);

INSERT INTO "public"."reservations" ("id", "first_name", "last_name", "email", "phone", "room_id", "status", "start_date", "end_date", "created_at", "updated_at") VALUES
(1, 'Sinthorn', 'Pradutnam', 'psinthorn@gmail.com', '999 999 9999', 1, 'stay', '2021-05-13', '2021-05-15', '2021-05-11 18:06:22.847974', '2021-05-11 18:06:22.847974');
