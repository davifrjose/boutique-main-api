
CREATE TABLE "users" (
    "id" VARCHAR NOT NULL PRIMARY KEY,
    "first_name" varchar(255) NOT NULL,
    "last_name" varchar(255) NOT NULL,
    "email" varchar NOT NULL UNIQUE,
    "password" varchar(255) NOT NULL,
    "phone_number" varchar(255) NOT NULL UNIQUE,
    "created_at" timestamptz ,
    "updated_at" timestamptz
);

CREATE UNIQUE INDEX "email" ON "users" ("email");