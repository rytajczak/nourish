import { createUserSecurities, getIdFromEmail } from "~~/server/utils/users";

export default defineOAuthGoogleEventHandler({
  async onSuccess(event, { user }) {
    const existingUserId = await getIdFromEmail(user.email);

    if (existingUserId) {
      user.id = existingUserId;
    } else {
      const newUserId = await createNewUser(user.email, user.name, "google");
      const spoonacularCreds = await connectUser(
        user.name,
        user.given_name,
        user.family_name,
      );
      await createUserSecurities(newUserId, spoonacularCreds);
      await createUserProfile(newUserId, user.picture, "");
      user.id = newUserId;
    }

    await setUserSession(event, { user });
    return sendRedirect(event, "/planner");
  },
});
