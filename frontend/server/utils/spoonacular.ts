const headers = {
  "Content-Type": "application/json",
  "x-rapidapi-key": process.env.RAPIDAPI_KEY!,
  "x-rapidapi-host": process.env.RAPIDAPI_HOST!,
};

type ConnResp = {
  status: string;
  username: string;
  spoonacularPassword: string;
  hash: string;
};

export const connectUser = async (
  username: string,
  firstName: string,
  lastName: string,
) => {
  const res = await $fetch<ConnResp>(
    `https://${process.env.RAPIDAPI_HOST}/users/connect`,
    {
      method: "POST",
      headers,
      body: {
        username,
        firstName,
        lastName,
      },
    },
  );

  return res;
};
