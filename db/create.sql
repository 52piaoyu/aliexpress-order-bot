CREATE TABLE "users" (
	"id"	INTEGER NOT NULL UNIQUE,
	"aliexpress_login"	TEXT NOT NULL,
	"aliexpress_password"	TEXT NOT NULL,
	PRIMARY KEY("id")
);

CREATE TABLE "indexes" (
	"id"	INTEGER NOT NULL UNIQUE,
	"index"	TEXT NOT NULL,
	"url"	TEXT,
	"location"	TEXT,
	"last_modification"	TEXT,
	PRIMARY KEY("id")
);