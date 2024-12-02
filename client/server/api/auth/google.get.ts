/**
 * @description Handles the Google OAuth login and redirects to the appropriate page.
 */
export default defineOAuthGoogleEventHandler({
  async onSuccess(event, { user, tokens }) {
    await setUserSession(event, {
      user: {
        email: user.email,
        username: user.name,
        firstName: user.given_name,
        lastName: user.family_name,
        provider: "google",
        picture: user.picture,
      },
      secure: { idToken: tokens.id_token },
    });
    return sendRedirect(event, "/confirm");
  },

  onError(event, error) {
    console.error("Google OAuth error:", error);
    return sendRedirect(event, "/");
  },
});
