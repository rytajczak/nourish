import GoogleProvider from "next-auth/providers/google";
import { NuxtAuthHandler } from "#auth";

export default NuxtAuthHandler({
  secret: "test",
  providers: [
    //@ts-ignore
    GoogleProvider.default({
      clientId:
        "582006295032-2bqikoi7afbhqaca3a115ugc32pivkpb.apps.googleusercontent.com",
      clientSecret: "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",
    }),
  ],
  callbacks: {
    /* on before signin */
    async signIn({ user, account, profile, email, credentials }) {
      return true;
    },
    /* on redirect to another url */
    async redirect({ url, baseUrl }) {
      return baseUrl;
    },
    /* on session retrival */
    async session({ session, user, token }) {
      return session;
    },
    /* on JWT token creation or mutation */
    async jwt({ token, user, account, profile }) {
      return token;
    },
  },
});
