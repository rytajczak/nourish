declare module "#auth-utils" {
  interface User {
    id: string;
    username: string;
    email: string;
    provider: string;
    picture: string;
    diet: string | null;
    calories: number | null;
    carbs: number | null;
    protein: number | null;
    fat: number | null;
    created_at: string;
    modified_at: string;
  }
}
export {};
