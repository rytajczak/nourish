import { relations } from "drizzle-orm/relations";
import { users, security, profile, profileIntolerance, intolerance, profileLikedRecipe, likedRecipe } from "./schema";

export const securityRelations = relations(security, ({one}) => ({
	user: one(users, {
		fields: [security.userId],
		references: [users.id]
	}),
}));

export const usersRelations = relations(users, ({many}) => ({
	securities: many(security),
	profiles: many(profile),
}));

export const profileRelations = relations(profile, ({one, many}) => ({
	user: one(users, {
		fields: [profile.userId],
		references: [users.id]
	}),
	profileIntolerances: many(profileIntolerance),
	profileLikedRecipes: many(profileLikedRecipe),
}));

export const profileIntoleranceRelations = relations(profileIntolerance, ({one}) => ({
	profile: one(profile, {
		fields: [profileIntolerance.profileId],
		references: [profile.userId]
	}),
	intolerance: one(intolerance, {
		fields: [profileIntolerance.intoleranceId],
		references: [intolerance.id]
	}),
}));

export const intoleranceRelations = relations(intolerance, ({many}) => ({
	profileIntolerances: many(profileIntolerance),
}));

export const profileLikedRecipeRelations = relations(profileLikedRecipe, ({one}) => ({
	profile: one(profile, {
		fields: [profileLikedRecipe.profileId],
		references: [profile.userId]
	}),
	likedRecipe: one(likedRecipe, {
		fields: [profileLikedRecipe.likedRecipeId],
		references: [likedRecipe.id]
	}),
}));

export const likedRecipeRelations = relations(likedRecipe, ({many}) => ({
	profileLikedRecipes: many(profileLikedRecipe),
}));