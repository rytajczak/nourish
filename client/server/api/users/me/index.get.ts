import { getApiUrl, UserResourceResponse } from "~~/server/utils/bff";

export default defineEventHandler(async (event) => {
  try {
    const { secure } = await requireUserSession(event);

    const response = await $fetch<UserResourceResponse>(getApiUrl(event), {
      headers: { Authorization: `Bearer ${secure?.idToken}` },
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
    return null;
  }
});
