import { eq } from "drizzle-orm";
import { v4 as uuidv4 } from "uuid";

export default defineOAuthGoogleEventHandler({
  async onSuccess(event, { user }) {
    console.log(user);
    const result = await useDB()
      .select()
      .from(tables.users)
      .where(eq(tables.users.email, user.email))
      .limit(1)
      .execute();

    if (result.length > 0) {
      user.id = result[0].id;
      await setUserSession(event, { user });
      return sendRedirect(event, "/planner");
    }

    const createdUser = await useDB()
      .insert(tables.users)
      .values({
        id: uuidv4(),
        email: user.email,
        username: user.name,
        provider: "google",
      })
      .returning({ insertedId: tables.users.id });

    const res = await $fetch<Record<string, any>>(
      `https://${process.env.RAPIDAPI_HOST}/users/connect`,
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "x-rapidapi-key": process.env.RAPIDAPI_KEY!,
          "x-rapidapi-host": process.env.RAPIDAPI_HOST!,
        },
        body: {
          username: user.name,
          firstName: user.given_name,
          lastName: user.family_name,
        },
      },
    );

    user.id = createdUser[0].insertedId;
    await setUserSession(event, { user });
    return sendRedirect(event, "/planner");
  },
});
