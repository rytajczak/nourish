import { pgTable, unique, uuid, varchar, timestamp, foreignKey, integer, primaryKey } from "drizzle-orm/pg-core"
import { sql } from "drizzle-orm"



export const users = pgTable("users", {
	id: uuid().primaryKey().notNull(),
	email: varchar({ length: 255 }).notNull(),
	displayName: varchar("display_name", { length: 50 }),
	provider: varchar({ length: 50 }),
	createdAt: timestamp("created_at", { mode: 'string' }).default(sql`CURRENT_TIMESTAMP`),
	lastSignInAt: timestamp("last_sign_in_at", { mode: 'string' }).default(sql`CURRENT_TIMESTAMP`),
}, (table) => {
	return {
		usersEmailKey: unique("users_email_key").on(table.email),
	}
});

export const security = pgTable("security", {
	userId: uuid("user_id").primaryKey().notNull(),
	spoonacularUsername: varchar("spoonacular_username", { length: 100 }),
	spoonacularHash: varchar("spoonacular_hash", { length: 100 }),
	spoonacularPassword: varchar("spoonacular_password", { length: 100 }),
}, (table) => {
	return {
		securityUserIdFkey: foreignKey({
			columns: [table.userId],
			foreignColumns: [users.id],
			name: "security_user_id_fkey"
		}).onDelete("cascade"),
	}
});

export const profile = pgTable("profile", {
	userId: uuid("user_id").primaryKey().notNull(),
	username: varchar({ length: 50 }).notNull(),
	picture: varchar({ length: 100 }),
	diet: varchar({ length: 100 }),
	createdAt: timestamp("created_at", { mode: 'string' }).default(sql`CURRENT_TIMESTAMP`),
	modifiedAt: timestamp("modified_at", { mode: 'string' }).default(sql`CURRENT_TIMESTAMP`),
}, (table) => {
	return {
		profileUserIdFkey: foreignKey({
			columns: [table.userId],
			foreignColumns: [users.id],
			name: "profile_user_id_fkey"
		}).onDelete("cascade"),
	}
});

export const intolerance = pgTable("intolerance", {
	id: uuid().primaryKey().notNull(),
	name: varchar({ length: 100 }).notNull(),
}, (table) => {
	return {
		intoleranceNameKey: unique("intolerance_name_key").on(table.name),
	}
});

export const savedRecipe = pgTable("saved_recipe", {
	id: integer().primaryKey().notNull(),
	createdAt: timestamp("created_at", { mode: 'string' }).default(sql`CURRENT_TIMESTAMP`),
	modifiedAt: timestamp("modified_at", { mode: 'string' }).default(sql`CURRENT_TIMESTAMP`),
});

export const profileIntolerance = pgTable("profile_intolerance", {
	profileId: uuid("profile_id").notNull(),
	intoleranceId: uuid("intolerance_id").notNull(),
}, (table) => {
	return {
		profileIntoleranceProfileIdFkey: foreignKey({
			columns: [table.profileId],
			foreignColumns: [profile.userId],
			name: "profile_intolerance_profile_id_fkey"
		}).onDelete("cascade"),
		profileIntoleranceIntoleranceIdFkey: foreignKey({
			columns: [table.intoleranceId],
			foreignColumns: [intolerance.id],
			name: "profile_intolerance_intolerance_id_fkey"
		}).onDelete("cascade"),
		profileIntolerancePkey: primaryKey({ columns: [table.profileId, table.intoleranceId], name: "profile_intolerance_pkey"}),
	}
});

export const profileSavedRecipe = pgTable("profile_saved_recipe", {
	profileId: uuid("profile_id").notNull(),
	savedRecipeId: integer("saved_recipe_id").notNull(),
}, (table) => {
	return {
		profileSavedRecipeProfileIdFkey: foreignKey({
			columns: [table.profileId],
			foreignColumns: [profile.userId],
			name: "profile_saved_recipe_profile_id_fkey"
		}).onDelete("cascade"),
		profileSavedRecipeSavedRecipeIdFkey: foreignKey({
			columns: [table.savedRecipeId],
			foreignColumns: [savedRecipe.id],
			name: "profile_saved_recipe_saved_recipe_id_fkey"
		}).onDelete("cascade"),
		profileSavedRecipePkey: primaryKey({ columns: [table.profileId, table.savedRecipeId], name: "profile_saved_recipe_pkey"}),
	}
});
