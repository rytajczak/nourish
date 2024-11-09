import { createUserSecurities, getIdFromEmail } from "~~/server/utils/users";

/**
 * @description Handles the Google OAuth login and redirects to the planner page.
 */
export default defineOAuthGoogleEventHandler({
  async onSuccess(event, { user }) {
    let userId = await getIdFromEmail(user.email);

    if (!userId) {
      userId = await createNewUser(user.email, user.name, "google");
      const spoonacularCreds = await connectUser(
        user.name,
        user.given_name,
        user.family_name,
      );
      await createUserSecurities(userId, spoonacularCreds);
      await createUserProfile(userId, user.name, user.picture, "");
    }

    user.id = userId;
    await setUserSession(event, { user });

    return sendRedirect(event, "/confirm");
  },
});
