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
  username: string,
  provider: string,
): Promise<string> => {
  const result = await useDB()
    .insert(tables.users)
    .values({
      id: uuidv4(),
      email,
      username,
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
  picture: string,
  diet: string,
) => {
  await useDB().insert(tables.profile).values({
    userId,
    picture,
    diet,
  });
};
