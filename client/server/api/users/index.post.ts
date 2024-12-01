import { getApiUrl, UserResourceResponse } from "~~/server/utils/bff";

export default defineEventHandler(async (event) => {
  try {
    const { user, secure } = await requireUserSession(event);
    const body = await readBody(event);

    const response = await $fetch<UserResourceResponse>(getApiUrl(event), {
      headers: { Authorization: `Bearer ${secure?.idToken}` },
      method: "POST",
      body: {
        ...user,
        ...body,
      },
    });

    await setUserSession(event, {
      secure: {
        spoonName: response.spoonCredential.username,
        spoonHash: response.spoonCredential.hash,
      },
    });

    return {
      profile: response.profile,
      intolerances: response.intolerances,
      savedRecipes: response.savedRecipes,
    };
  } catch (error) {
    console.log(error);
    return null;
  }
});
