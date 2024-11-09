interface Response {
  profile: Profile;
}

export default defineEventHandler(async (event) => {
  const id = getRouterParam(event, "id");
  if (!id) return;

  const [{ profile, intolerances, savedRecipes }] = await useDB()
    .select({
      profile: tables.profile,
      intolerances: tables.profileIntolerance,
      savedRecipes: tables.profileSavedRecipe,
    })
    .from(tables.profile)
    .leftJoin(
      tables.profileIntolerance,
      eq(tables.profileIntolerance.profileId, tables.profile.userId),
    )
    .leftJoin(
      tables.profileSavedRecipe,
      eq(tables.profileSavedRecipe.profileId, tables.profile.userId),
    )
    .where(eq(tables.profile.userId, id))
    .limit(1);

  return { profile, intolerances, savedRecipes } as Response;
});
