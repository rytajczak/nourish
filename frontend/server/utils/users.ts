import { v4 as uuidv4 } from "uuid";

/**
 * Gets the user id in the users table from an email
 * @param email email of the user
 * @returns the id of the user
 */
export const getIdFromEmail = async (email: string): Promise<string> => {
  const result = await useDB()
    .select()
    .from(tables.users)
    .where(eq(tables.users.email, email))
    .limit(1)
    .execute();

  if (result.length > 0) return result[0].id;
  return "";
};

/**
 * Create a user in the the users database
 * @param email email of the user
 * @param username username of the user
 * @param provider auth provider ('google' or 'github')
 * @returns the id of the created user
 */
export const createNewUser = async (
  email: string,
  displayName: string,
  provider: string,
): Promise<string> => {
  const result = await useDB()
    .insert(tables.users)
    .values({
      id: uuidv4(),
      email,
      displayName,
      provider,
    })
    .returning({ insertedId: tables.users.id });

  return result[0].insertedId;
};

/**
 * Create spoonacular securities for spoonacular user
 * @param userId id of the user to create the securites for
 * @param spoonCreds credentials gotten from connecting user to spoonacular
 */
export const createUserSecurities = async (
  userId: string,
  spoonCreds: {
    status: string;
    username: string;
    spoonacularPassword: string;
    hash: string;
  },
) => {
  await useDB().insert(tables.security).values({
    userId,
    spoonacularUsername: spoonCreds.username,
    spoonacularPassword: spoonCreds.spoonacularPassword,
    spoonacularHash: spoonCreds.hash,
  });
};

/**
 *
 * @param userId id of the user to create profile for
 * @param picture url of the profile picture for the user's account
 * @param diet diet of the user
 */
export const createUserProfile = async (
  userId: string,
  username: string,
  picture: string,
  diet: string,
) => {
  await useDB().insert(tables.profile).values({
    userId,
    username,
    picture,
    diet,
  });
};

/**
 * Get the user's profile
 * @param userId id of the user to get the profile for
 * @returns the user's profile
 */
export const getUserProfile = async (
  userId: string,
): Promise<Profile | null> => {
  const result = await useDB()
    .select()
    .from(tables.profile)
    .where(eq(tables.profile.userId, userId))
    .limit(1)
    .execute();

  return result[0];
};

/**
 * Save a recipe to the user's profile
 * @param userId id of the user to save the recipe to
 * @param recipeId id of the recipe to save
 */
export const saveRecipeToProfile = async (userId: string, recipeId: number) => {
  const existingRecipe = await useDB()
    .select()
    .from(tables.profileSavedRecipe)
    .where(
      and(
        eq(tables.profileSavedRecipe.profileId, userId),
        eq(tables.profileSavedRecipe.savedRecipeId, recipeId),
      ),
    )
    .limit(1)
    .execute();

  if (existingRecipe.length > 0) return;
  const existingSavedRecipe = await useDB()
    .select()
    .from(tables.savedRecipe)
    .where(eq(tables.savedRecipe.id, recipeId))
    .limit(1)
    .execute();

  if (existingSavedRecipe.length === 0) {
    await useDB().insert(tables.savedRecipe).values({
      id: recipeId,
    });
  }

  await useDB().insert(tables.profileSavedRecipe).values({
    profileId: userId,
    savedRecipeId: recipeId,
  });
};

/**
 * Get the saved recipes for a user
 * @param userId id of the user to get the saved recipes for
 * @returns an array of the ids of the saved recipes
 */
export const getSavedRecipes = async (userId: string) => {
  const result = await useDB()
    .select()
    .from(tables.profileSavedRecipe)
    .where(eq(tables.profileSavedRecipe.profileId, userId))
    .execute();

  return result.map((recipe) => recipe.savedRecipeId);
};
