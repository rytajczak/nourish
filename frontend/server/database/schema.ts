import { pgTable, varchar, timestamp, foreignKey, primaryKey } from "drizzle-orm/pg-core"
  import { sql } from "drizzle-orm"




export const auth = pgTable("auth", {
	id: varchar().primaryKey().notNull(),
	email: varchar().notNull(),
	password: varchar().notNull(),
	provider: varchar().notNull(),
	createdAt: timestamp("created_at", { mode: 'string' }).default(sql`CURRENT_TIMESTAMP`).notNull(),
	lastSignInAt: timestamp("last_sign_in_at", { mode: 'string' }).default(sql`CURRENT_TIMESTAMP`).notNull(),
});

export const profile = pgTable("profile", {
	id: varchar().primaryKey().notNull(),
	username: varchar().notNull(),
	firstName: varchar("first_name"),
	lastName: varchar("last_name"),
	diet: varchar(),
	createdAt: timestamp("created_at", { mode: 'string' }).default(sql`CURRENT_TIMESTAMP`).notNull(),
	modifiedAt: timestamp("modified_at", { mode: 'string' }).default(sql`CURRENT_TIMESTAMP`).notNull(),
},
(table) => {
	return {
		fkAuth: foreignKey({
			columns: [table.id],
			foreignColumns: [auth.id],
			name: "fk_auth"
		}).onDelete("cascade"),
	}
});

export const intolerance = pgTable("intolerance", {
	id: varchar().primaryKey().notNull(),
	name: varchar().notNull(),
});

export const likedRecipe = pgTable("liked_recipe", {
	id: varchar().primaryKey().notNull(),
	title: varchar().notNull(),
	image: varchar(),
	createdAt: timestamp("created_at", { mode: 'string' }).default(sql`CURRENT_TIMESTAMP`).notNull(),
	modifiedAt: timestamp("modified_at", { mode: 'string' }).default(sql`CURRENT_TIMESTAMP`).notNull(),
});

export const profileIntolerance = pgTable("profile_intolerance", {
	profileId: varchar("profile_id").notNull(),
	intoleranceId: varchar("intolerance_id").notNull(),
},
(table) => {
	return {
		fkProfile: foreignKey({
			columns: [table.profileId],
			foreignColumns: [profile.id],
			name: "fk_profile"
		}).onDelete("cascade"),
		fkIntolerance: foreignKey({
			columns: [table.intoleranceId],
			foreignColumns: [intolerance.id],
			name: "fk_intolerance"
		}).onDelete("cascade"),
		profileIntolerancePkey: primaryKey({ columns: [table.profileId, table.intoleranceId], name: "profile_intolerance_pkey"}),
	}
});

export const profileLikedRecipe = pgTable("profile_liked_recipe", {
	profileId: varchar("profile_id").notNull(),
	likedRecipeId: varchar("liked_recipe_id").notNull(),
},
(table) => {
	return {
		fkProfile: foreignKey({
			columns: [table.profileId],
			foreignColumns: [profile.id],
			name: "fk_profile"
		}).onDelete("cascade"),
		fkLikedRecipe: foreignKey({
			columns: [table.likedRecipeId],
			foreignColumns: [likedRecipe.id],
			name: "fk_liked_recipe"
		}).onDelete("cascade"),
		profileLikedRecipePkey: primaryKey({ columns: [table.profileId, table.likedRecipeId], name: "profile_liked_recipe_pkey"}),
	}
});
