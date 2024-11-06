-- Current sql file was generated after introspecting the database
-- If you want to run this migration please uncomment this code before executing migrations
/*
CREATE TABLE IF NOT EXISTS "users" (
	"id" uuid PRIMARY KEY NOT NULL,
	"email" varchar(255) NOT NULL,
	"username" varchar(50) NOT NULL,
	"provider" varchar(50),
	"created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
	"last_sign_in_at" timestamp DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT "users_email_key" UNIQUE("email")
);
--> statement-breakpoint
CREATE TABLE IF NOT EXISTS "security" (
	"user_id" uuid PRIMARY KEY NOT NULL,
	"spoonacular_username" varchar(100),
	"spoonacular_hash" varchar(100),
	"spoonacular_password" varchar(100)
);
--> statement-breakpoint
CREATE TABLE IF NOT EXISTS "profile" (
	"user_id" uuid PRIMARY KEY NOT NULL,
	"picture" varchar(100),
	"diet" varchar(100),
	"created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
	"modified_at" timestamp DEFAULT CURRENT_TIMESTAMP
);
--> statement-breakpoint
CREATE TABLE IF NOT EXISTS "intolerance" (
	"id" uuid PRIMARY KEY NOT NULL,
	"name" varchar(100) NOT NULL,
	CONSTRAINT "intolerance_name_key" UNIQUE("name")
);
--> statement-breakpoint
CREATE TABLE IF NOT EXISTS "liked_recipe" (
	"id" uuid PRIMARY KEY NOT NULL,
	"title" varchar(255) NOT NULL,
	"image" varchar(255),
	"created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
	"modified_at" timestamp DEFAULT CURRENT_TIMESTAMP
);
--> statement-breakpoint
CREATE TABLE IF NOT EXISTS "profile_intolerance" (
	"profile_id" uuid NOT NULL,
	"intolerance_id" uuid NOT NULL,
	CONSTRAINT "profile_intolerance_pkey" PRIMARY KEY("profile_id","intolerance_id")
);
--> statement-breakpoint
CREATE TABLE IF NOT EXISTS "profile_liked_recipe" (
	"profile_id" uuid NOT NULL,
	"liked_recipe_id" uuid NOT NULL,
	CONSTRAINT "profile_liked_recipe_pkey" PRIMARY KEY("profile_id","liked_recipe_id")
);
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "security" ADD CONSTRAINT "security_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users"("id") ON DELETE cascade ON UPDATE no action;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "profile" ADD CONSTRAINT "profile_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users"("id") ON DELETE cascade ON UPDATE no action;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "profile_intolerance" ADD CONSTRAINT "profile_intolerance_profile_id_fkey" FOREIGN KEY ("profile_id") REFERENCES "public"."profile"("user_id") ON DELETE cascade ON UPDATE no action;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "profile_intolerance" ADD CONSTRAINT "profile_intolerance_intolerance_id_fkey" FOREIGN KEY ("intolerance_id") REFERENCES "public"."intolerance"("id") ON DELETE cascade ON UPDATE no action;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "profile_liked_recipe" ADD CONSTRAINT "profile_liked_recipe_profile_id_fkey" FOREIGN KEY ("profile_id") REFERENCES "public"."profile"("user_id") ON DELETE cascade ON UPDATE no action;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "profile_liked_recipe" ADD CONSTRAINT "profile_liked_recipe_liked_recipe_id_fkey" FOREIGN KEY ("liked_recipe_id") REFERENCES "public"."liked_recipe"("id") ON DELETE cascade ON UPDATE no action;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;

*/