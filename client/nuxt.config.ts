export default defineNuxtConfig({
  future: { compatibilityVersion: 4 },
  compatibilityDate: "2024-04-03",
  modules: [
    "nuxt-auth-utils",
    "@nuxt/ui",
    "@vueuse/nuxt",
    "@nuxt/image",
    "@pinia/nuxt",
    "pinia-plugin-persistedstate/nuxt",
  ],
  css: ["./app/assets/main.css"],
  app: { pageTransition: { name: "page", mode: "out-in" } },
  ssr: false,

  runtimeConfig: {
    spoonacularApiKey: "me at the red lobster",
    spoonacularApiHost: "me at the red lobster",
    public: {
      apiUrl: "me at the red lobster",
    },
  },

  devtools: { enabled: true },
});
