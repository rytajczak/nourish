const config = useRuntimeConfig();

interface UserResponse extends Record<string, string> {
  id: string;
  username: string;
  email: string;
  provider: string;
  picture: string;
  diet: string;
}

/**
 * @description Handles the Google OAuth login and redirects to the planner page.
 */
export default defineOAuthGoogleEventHandler({
  async onSuccess(event, { user, tokens }) {
    const existingUser = await $fetch<UserResponse>(
      `${config.public.apiUrl}/users/me`,
      {
        method: "GET",
        headers: {
          Authorization: `Bearer ${tokens.id_token}`,
        },
      },
    );
    if (existingUser.msg === "user not found") {
      const newUser = await $fetch<UserResponse>(
        `${config.public.apiUrl}/users/`,
        {
          method: "POST",
          headers: {
            Authorization: `Bearer ${tokens.id_token}`,
          },
          body: {
            email: user.email,
            username: user.name,
            picture: user.picture,
            provider: "google",
          },
        },
      );
      await setUserSession(event, { user: newUser });
      return sendRedirect(event, "/onboarding");
    }

    await setUserSession(event, { user: existingUser });
    return sendRedirect(event, "/onboarding");
  },
});
