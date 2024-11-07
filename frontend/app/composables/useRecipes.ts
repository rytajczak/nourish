/**
 * A composable for managing recipes
 */
export const useRecipes = () => {
  /**
   * The recipes that the user has saved
   */
  const savedRecipes = useState<number[]>("savedRecipes", () => []);

  /**
   * Check if a recipe is saved
   * @param recipeId id of the recipe to check
   * @returns true if the recipe is saved, false otherwise
   */
  const isSaved = (recipeId: number) => {
    return savedRecipes.value.includes(recipeId);
  };

  /**
   * Debounce the saving of recipes to the user's profile
   */
  const debouncedSave = useDebounce((recipeIds: number[]) => {
    console.log("saving recipes");
  }, 3000);

  /**
   * Toggle the saving of a recipe to the user's profile
   * @param recipeId id of the recipe to toggle
   */
  const toggleSave = (recipeId: number) => {
    const isCurrentlySaved = savedRecipes.value.includes(recipeId);

    if (isCurrentlySaved) {
      savedRecipes.value = savedRecipes.value.filter((id) => id !== recipeId);
    } else {
      savedRecipes.value = [...savedRecipes.value, recipeId];
    }

    debouncedSave(savedRecipes.value);
  };

  return { savedRecipes, isSaved, toggleSave };
};
