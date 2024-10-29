CREATE TABLE IF NOT EXISTS "auth" (
	"id" varchar PRIMARY KEY NOT NULL,
	"email" varchar NOT NULL,
	"password" varchar NOT NULL,
	"provider" varchar NOT NULL,
	"created_at" timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	"last_sign_in_at" timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL
);
--> statement-breakpoint
CREATE TABLE IF NOT EXISTS "intolerance" (
	"id" varchar PRIMARY KEY NOT NULL,
	"name" varchar NOT NULL
);
--> statement-breakpoint
CREATE TABLE IF NOT EXISTS "liked_recipe" (
	"id" varchar PRIMARY KEY NOT NULL,
	"title" varchar NOT NULL,
	"image" varchar,
	"created_at" timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	"modified_at" timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL
);
--> statement-breakpoint
CREATE TABLE IF NOT EXISTS "profile" (
	"id" varchar PRIMARY KEY NOT NULL,
	"username" varchar NOT NULL,
	"first_name" varchar,
	"last_name" varchar,
	"diet" varchar,
	"created_at" timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	"modified_at" timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL
);
--> statement-breakpoint
CREATE TABLE IF NOT EXISTS "profile_intolerance" (
	"profile_id" varchar NOT NULL,
	"intolerance_id" varchar NOT NULL,
	CONSTRAINT "profile_intolerance_pkey" PRIMARY KEY("profile_id","intolerance_id")
);
--> statement-breakpoint
CREATE TABLE IF NOT EXISTS "profile_liked_recipe" (
	"profile_id" varchar NOT NULL,
	"liked_recipe_id" varchar NOT NULL,
	CONSTRAINT "profile_liked_recipe_pkey" PRIMARY KEY("profile_id","liked_recipe_id")
);
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "profile" ADD CONSTRAINT "fk_auth" FOREIGN KEY ("id") REFERENCES "public"."auth"("id") ON DELETE cascade ON UPDATE no action;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "profile_intolerance" ADD CONSTRAINT "fk_profile" FOREIGN KEY ("profile_id") REFERENCES "public"."profile"("id") ON DELETE cascade ON UPDATE no action;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "profile_intolerance" ADD CONSTRAINT "fk_intolerance" FOREIGN KEY ("intolerance_id") REFERENCES "public"."intolerance"("id") ON DELETE cascade ON UPDATE no action;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "profile_liked_recipe" ADD CONSTRAINT "fk_profile" FOREIGN KEY ("profile_id") REFERENCES "public"."profile"("id") ON DELETE cascade ON UPDATE no action;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "profile_liked_recipe" ADD CONSTRAINT "fk_liked_recipe" FOREIGN KEY ("liked_recipe_id") REFERENCES "public"."liked_recipe"("id") ON DELETE cascade ON UPDATE no action;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
