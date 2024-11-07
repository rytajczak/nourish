import { relations } from "drizzle-orm/relations";
import { users, security, profile, profileIntolerance, intolerance, profileSavedRecipe, savedRecipe } from "./schema";

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
	profileSavedRecipes: many(profileSavedRecipe),
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

export const profileSavedRecipeRelations = relations(profileSavedRecipe, ({one}) => ({
	profile: one(profile, {
		fields: [profileSavedRecipe.profileId],
		references: [profile.userId]
	}),
	savedRecipe: one(savedRecipe, {
		fields: [profileSavedRecipe.savedRecipeId],
		references: [savedRecipe.id]
	}),
}));

export const savedRecipeRelations = relations(savedRecipe, ({many}) => ({
	profileSavedRecipes: many(profileSavedRecipe),
}));