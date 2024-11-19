declare module "#auth-utils" {
  interface User {
    email: string;
    username: string;
    provider: string;
    picture: string;
  }
  interface SecureSessionData {
    idToken: string;
    spoonName?: string;
    spoonHash?: string;
  }
}
export {};
