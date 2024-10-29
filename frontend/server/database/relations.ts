import { relations } from "drizzle-orm/relations";
import { auth, profile, profileIntolerance, intolerance, profileLikedRecipe, likedRecipe } from "./schema";

export const profileRelations = relations(profile, ({one, many}) => ({
	auth: one(auth, {
		fields: [profile.id],
		references: [auth.id]
	}),
	profileIntolerances: many(profileIntolerance),
	profileLikedRecipes: many(profileLikedRecipe),
}));

export const authRelations = relations(auth, ({many}) => ({
	profiles: many(profile),
}));

export const profileIntoleranceRelations = relations(profileIntolerance, ({one}) => ({
	profile: one(profile, {
		fields: [profileIntolerance.profileId],
		references: [profile.id]
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
		references: [profile.id]
	}),
	likedRecipe: one(likedRecipe, {
		fields: [profileLikedRecipe.likedRecipeId],
		references: [likedRecipe.id]
	}),
}));

export const likedRecipeRelations = relations(likedRecipe, ({many}) => ({
	profileLikedRecipes: many(profileLikedRecipe),
}));