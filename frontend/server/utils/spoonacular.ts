export const connectUser = async (
  username: string,
  firstName: string,
  lastName: string,
) => {
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
        username,
        firstName,
        lastName,
      },
    },
  );
  return res as {
    status: string;
    username: string;
    spoonacularPassword: string;
    hash: string;
  };
};
