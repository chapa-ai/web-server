CREATE TABLE IF NOT EXISTS "documents"
(
    "id" serial,
    "name" character varying NOT NULL,
    "file" boolean NOT NULL,
    "public" boolean NOT NULL,
    "token" character varying NOT NULL,
    "mime" character varying NOT NULL,
    "grant" text[] NOT NULL,
    "json"   character varying NOT NULL,
    "directory"  character varying NOT NULL,
    "created" date NOT NULL,
    CONSTRAINT "Document_pkey" PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "users"
(
    "id" SERIAL  UNIQUE NOT NULL,
    "login" character varying NOT NULL,
    "password" character varying NOT NULL,
    CONSTRAINT "pk_Users" PRIMARY KEY ("id")
    );

INSERT INTO "users"("login", "password") VALUES('admin', '$2a$12$Ag66QnSOVJWG2ySm0UJod.1M/Vjsejtp5lkcqX19gxxMIHD4C6dpC');

CREATE TABLE IF NOT EXISTS "sessions"
(
    "id" SERIAL  UNIQUE NOT NULL,
    "userId" integer NOT NULL,
    "token" character varying NOT NULL,
    CONSTRAINT "pk_Sessions" PRIMARY KEY ("id")
    );

INSERT INTO "sessions"("userId", "token") VALUES(1, '99a0833f3-a3dc-151a-b86e-01c55d54ca8a');


ALTER TABLE "sessions" ADD CONSTRAINT "fk_Sessions" FOREIGN KEY("userId")
    REFERENCES "users" ("id") ON DELETE CASCADE;


