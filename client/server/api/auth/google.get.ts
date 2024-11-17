import { User } from "#auth-utils";

const config = useRuntimeConfig();

/**
 * @description Handles the Google OAuth login and redirects to the dashboard page.
 */
export default defineOAuthGoogleEventHandler({
  async onSuccess(event, { user, tokens }) {
    setCookie(event, "idToken", tokens.id_token, {
      httpOnly: true,
      path: "/",
    });
    try {
      await $fetch<User>(`${config.public.apiUrl}/users/me`, {
        method: "GET",
        headers: {
          Authorization: `Bearer ${tokens.id_token}`,
        },
      });
      return sendRedirect(event, "/confirm");
    } catch (error) {
      await $fetch<User>(`${config.public.apiUrl}/users/`, {
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
      });
      return sendRedirect(event, "/onboarding");
    }
  },
});
