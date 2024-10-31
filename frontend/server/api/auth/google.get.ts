import axios from "axios";
import { eq } from "drizzle-orm";
import { v4 as uuidv4 } from "uuid";

export default defineOAuthGoogleEventHandler({
  async onSuccess(event, { user }) {
    const result = await useDB()
      .select()
      .from(tables.auth)
      .where(eq(tables.auth.email, user.email))
      .limit(1)
      .execute();

    if (result.length > 0) {
      user.id = result[0].id;
      await setUserSession(event, { user });
      return sendRedirect(event, "/planner");
    }

    const authResult = await useDB()
      .insert(tables.auth)
      .values({
        id: uuidv4(),
        email: user.email,
        provider: "google",
      })
      .returning({ insertedId: tables.auth.id });

    const headers = {
      "x-rapidapi-key": process.env.RAPIDAPI_KEY!,
      "x-rapidapi-host": process.env.RAPIDAPI_HOST!,
      "Content-Type": "application/json",
    };

    const res = await axios.post(
      `https://${process.env.RAPIDAPI_HOST}/users/connect`,
      {
        username: user.name,
        firstName: user.given_name,
        lastName: user.family_name,
        email: user.email,
      },
      { headers },
    );

    await useDB().insert(tables.profile).values({
      id: authResult[0].insertedId,
      username: res.data.username,
      firstName: user.given_name,
      lastName: user.family_name,
      spoonacularPassword: res.data.spoonacularPassword,
      hash: res.data.hash,
      picture: user.picture,
    });

    user.id = authResult[0].insertedId;
    await setUserSession(event, { user });
    return sendRedirect(event, "/planner");
  },
});
