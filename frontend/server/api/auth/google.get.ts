interface ExistingUserResponse {
  id: string;
  username: string;
  email: string;
  provider: string;
  picture: string;
  diet: string;
  createdAt: string;
  modifiedAt: string;
}

/**
 * @description Handles the Google OAuth login and redirects to the planner page.
 */
export default defineOAuthGoogleEventHandler({
  async onSuccess(event, { user, tokens }) {
    setCookie(event, "id-token", tokens.id_token, {
      httpOnly: true,
      path: "/",
      maxAge: tokens.expires_in,
    });
    await setUserSession(event, { user });
    return sendRedirect(event, "/confirm");
  },
});
