CREATE TYPE "state_game" AS ENUM (
  'active',
  'inactive',
  'upcoming'
);

CREATE TABLE "client" (
  "id" uuid PRIMARY KEY,
  "name" varchar NOT NULL,
  "surname" varchar NOT NULL,
  "picture" varchar NOT NULL,
  "birthdate" date NOT NULL,
  "phone_number" varchar NOT NULL,
  "vat_id" int NOT NULL,
  "active" boolean NOT NULL
);

CREATE TABLE "access" (
  "id" uuid PRIMARY KEY,
  "id_client" uuid UNIQUE NOT NULL,
  "oauth" boolean NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar
);

CREATE TABLE "client_search_history" (
  "id" uuid PRIMARY KEY,
  "id_game" uuid NOT NULL,
  "id_client" uuid NOT NULL,
  "date_time" timestamp NOT NULL
);

CREATE TABLE "wallet" (
  "id" uuid PRIMARY KEY,
  "id_client" uuid UNIQUE NOT NULL,
  "balance" float NOT NULL,
  "coin" char NOT NULL
);

CREATE TABLE "address" (
  "id" uuid PRIMARY KEY,
  "id_client" uuid UNIQUE NOT NULL,
  "street" varchar NOT NULL,
  "door_number" varchar,
  "zip_code" varchar NOT NULL,
  "city" varchar NOT NULL,
  "country" varchar NOT NULL
);

CREATE TABLE "publisher" (
  "id" uuid PRIMARY KEY,
  "name" varchar NOT NULL,
  "cover_image" varchar NOT NULL,
  "phone_number" varchar NOT NULL,
  "email" varchar NOT NULL
);

CREATE TABLE "game" (
  "id" uuid PRIMARY KEY,
  "id_publisher" uuid NOT NULL,
  "name" varchar NOT NULL,
  "price" float NOT NULL,
  "discount" float NOT NULL,
  "state" state_game NOT NULL,
  "cover_image" varchar NOT NULL,
  "release_date" date NOT NULL,
  "description" varchar NOT NULL,
  "download_link" varchar NOT NULL
);

CREATE TABLE "image_game" (
  "id" uuid PRIMARY KEY,
  "id_game" uuid NOT NULL,
  "image" varchar NOT NULL
);

CREATE TABLE "review" (
  "id" uuid PRIMARY KEY,
  "id_game" uuid NOT NULL,
  "id_client" uuid NOT NULL,
  "stars" int,
  "review" varchar
);

CREATE TABLE "tag_game" (
  "id_tag" uuid,
  "id_game" uuid,
  PRIMARY KEY ("id_tag", "id_game")
);

CREATE TABLE "tag" (
  "id" uuid PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "wishlist" (
  "id_game" uuid,
  "id_client" uuid,
  PRIMARY KEY ("id_game", "id_client")
);

CREATE TABLE "game_library" (
  "id_game" uuid,
  "id_client" uuid,
  PRIMARY KEY ("id_game", "id_client")
);

CREATE TABLE "cart" (
  "id_game" uuid,
  "id_client" uuid,
  PRIMARY KEY ("id_game", "id_client")
);

CREATE TABLE "invoice_header" (
  "id_invoice" uuid PRIMARY KEY,
  "id_client" uuid NOT NULL,
  "purchase_date" timestamp NOT NULL,
  "vat_id" int NOT NULL
);

CREATE TABLE "invoice_game" (
  "id_invoice" uuid,
  "id_game" uuid,
  "price" float NOT NULL,
  "discount" float NOT NULL,
  PRIMARY KEY ("id_invoice", "id_game")
);

CREATE TABLE "newsletter" (
  "email" varchar PRIMARY KEY
);

ALTER TABLE "client_search_history" ADD FOREIGN KEY ("id_client") REFERENCES "client" ("id") ON DELETE CASCADE;

ALTER TABLE "client_search_history" ADD FOREIGN KEY ("id_game") REFERENCES "game" ("id");

ALTER TABLE "access" ADD FOREIGN KEY ("id_client") REFERENCES "client" ("id") ON DELETE CASCADE;

ALTER TABLE "wallet" ADD FOREIGN KEY ("id_client") REFERENCES "client" ("id") ON DELETE CASCADE;

ALTER TABLE "client" ADD FOREIGN KEY ("id") REFERENCES "address" ("id_client") ON DELETE CASCADE;

ALTER TABLE "game" ADD FOREIGN KEY ("id_publisher") REFERENCES "publisher" ("id");

ALTER TABLE "tag_game" ADD FOREIGN KEY ("id_tag") REFERENCES "tag" ("id");

ALTER TABLE "tag_game" ADD FOREIGN KEY ("id_game") REFERENCES "game" ("id");

ALTER TABLE "image_game" ADD FOREIGN KEY ("id_game") REFERENCES "game" ("id");

ALTER TABLE "review" ADD FOREIGN KEY ("id_game") REFERENCES "game" ("id");

ALTER TABLE "review" ADD FOREIGN KEY ("id_client") REFERENCES "client" ("id") ON DELETE CASCADE;

ALTER TABLE "wishlist" ADD FOREIGN KEY ("id_game") REFERENCES "game" ("id");

ALTER TABLE "wishlist" ADD FOREIGN KEY ("id_client") REFERENCES "client" ("id") ON DELETE CASCADE;

ALTER TABLE "game_library" ADD FOREIGN KEY ("id_game") REFERENCES "game" ("id");

ALTER TABLE "game_library" ADD FOREIGN KEY ("id_client") REFERENCES "client" ("id") ON DELETE CASCADE;

ALTER TABLE "cart" ADD FOREIGN KEY ("id_game") REFERENCES "game" ("id");

ALTER TABLE "cart" ADD FOREIGN KEY ("id_client") REFERENCES "client" ("id") ON DELETE CASCADE;

ALTER TABLE "invoice_header" ADD FOREIGN KEY ("id_client") REFERENCES "client" ("id") ON DELETE CASCADE;

ALTER TABLE "invoice_game" ADD FOREIGN KEY ("id_invoice") REFERENCES "invoice_header" ("id_invoice");

ALTER TABLE "invoice_game" ADD FOREIGN KEY ("id_game") REFERENCES "game" ("id");

