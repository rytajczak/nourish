import { Recipe } from "~~/server/utils/types";

export default defineCachedEventHandler(
  async (event) => {
    const id = getRouterParam(event, "id");
    const res = await $fetch<Recipe>(`http://localhost:8080/${id}`);
    return res;
  },
  {
    maxAge: 60 * 60,
    staleMaxAge: 60 * 60,
    getKey(event) {
      return event.path;
    },
  },
);
