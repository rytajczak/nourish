import { pgTable, unique, uuid, varchar, timestamp, foreignKey, primaryKey } from "drizzle-orm/pg-core"
  import { sql } from "drizzle-orm"




export const users = pgTable("users", {
	id: uuid().primaryKey().notNull(),
	email: varchar({ length: 255 }).notNull(),
	username: varchar({ length: 50 }).notNull(),
	firstName: varchar("first_name", { length: 50 }),
	lastName: varchar("last_name", { length: 50 }),
	provider: varchar({ length: 50 }),
	createdAt: timestamp("created_at", { mode: 'string' }).default(sql`CURRENT_TIMESTAMP`),
	lastSignInAt: timestamp("last_sign_in_at", { mode: 'string' }).default(sql`CURRENT_TIMESTAMP`),
},
(table) => {
	return {
		usersEmailKey: unique("users_email_key").on(table.email),
	}
});

export const security = pgTable("security", {
	userId: uuid("user_id").primaryKey().notNull(),
	spoonacularHash: varchar("spoonacular_hash", { length: 100 }),
	spoonacularPassword: varchar("spoonacular_password", { length: 100 }),
},
(table) => {
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
	picture: varchar({ length: 100 }),
	diet: varchar({ length: 100 }),
	createdAt: timestamp("created_at", { mode: 'string' }).default(sql`CURRENT_TIMESTAMP`),
	modifiedAt: timestamp("modified_at", { mode: 'string' }).default(sql`CURRENT_TIMESTAMP`),
},
(table) => {
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
},
(table) => {
	return {
		intoleranceNameKey: unique("intolerance_name_key").on(table.name),
	}
});

export const likedRecipe = pgTable("liked_recipe", {
	id: uuid().primaryKey().notNull(),
	title: varchar({ length: 255 }).notNull(),
	image: varchar({ length: 255 }),
	createdAt: timestamp("created_at", { mode: 'string' }).default(sql`CURRENT_TIMESTAMP`),
	modifiedAt: timestamp("modified_at", { mode: 'string' }).default(sql`CURRENT_TIMESTAMP`),
});

export const profileIntolerance = pgTable("profile_intolerance", {
	profileId: uuid("profile_id").notNull(),
	intoleranceId: uuid("intolerance_id").notNull(),
},
(table) => {
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

export const profileLikedRecipe = pgTable("profile_liked_recipe", {
	profileId: uuid("profile_id").notNull(),
	likedRecipeId: uuid("liked_recipe_id").notNull(),
},
(table) => {
	return {
		profileLikedRecipeProfileIdFkey: foreignKey({
			columns: [table.profileId],
			foreignColumns: [profile.userId],
			name: "profile_liked_recipe_profile_id_fkey"
		}).onDelete("cascade"),
		profileLikedRecipeLikedRecipeIdFkey: foreignKey({
			columns: [table.likedRecipeId],
			foreignColumns: [likedRecipe.id],
			name: "profile_liked_recipe_liked_recipe_id_fkey"
		}).onDelete("cascade"),
		profileLikedRecipePkey: primaryKey({ columns: [table.profileId, table.likedRecipeId], name: "profile_liked_recipe_pkey"}),
	}
});
