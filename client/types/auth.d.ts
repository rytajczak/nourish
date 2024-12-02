declare module "#auth-utils" {
  interface User {
    email: string;
    username: string;
    firstName: string;
    lastName: string;
    provider: string;
    picture: string;
  }
  interface SecureSessionData {
    idToken?: string;
    spoonName?: string;
    spoonHash?: string;
  }
}
export {};
