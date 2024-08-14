CREATE TABLE "addresses" (
    "id" VARCHAR NOT NULL PRIMARY KEY,
    "user_id" VARCHAR NOT NULL,
    "address" varchar(255),
    "city" varchar(255),
    "province" varchar(255),
    "zip" varchar(20),
    "country" varchar(255),
    "created_at" timestamptz,
    "updated_at" timestamptz
);

ALTER TABLE
    "addresses"
ADD
    CONSTRAINT "fk_users_addresses" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;