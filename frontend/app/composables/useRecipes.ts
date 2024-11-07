import type { RecipePreview } from "~~/server/api/recipes/search/index.get";

/**
 * A composable for managing recipes
 */
export const useRecipes = () => {
  /**
   * The recipes that the user has saved
   */
  const savedRecipes = useState<RecipePreview[]>("savedRecipes", () => []);

  /**
   * Check if a recipe is saved
   * @param recipeId id of the recipe to check
   * @returns true if the recipe is saved, false otherwise
   */
  const isSaved = (recipeId: number) => {
    return savedRecipes.value.some((recipe) => recipe.id === recipeId);
  };

  /**
   * Debounce the saving of recipes to the user's profile
   */
  const debouncedSave = useDebounce((recipes: RecipePreview[]) => {
    console.log("saving recipes", recipes);
  }, 3000);

  /**
   * Toggle the saving of a recipe to the user's profile
   * @param recipeId id of the recipe to toggle
   */
  const toggleSave = (recipe: RecipePreview) => {
    const isCurrentlySaved = savedRecipes.value.some((r) => r.id === recipe.id);

    if (isCurrentlySaved) {
      savedRecipes.value = savedRecipes.value.filter((r) => r.id !== recipe.id);
    } else {
      savedRecipes.value = [...savedRecipes.value, recipe];
    }

    debouncedSave(savedRecipes.value);
  };

  return { savedRecipes, isSaved, toggleSave };
};
