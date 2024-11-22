export default defineNuxtConfig({
  future: { compatibilityVersion: 4 },
  compatibilityDate: "2024-04-03",

  ssr: false,

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

  runtimeConfig: {
    spoonacularApiKey: "me at the red lobster",
    spoonacularApiHost: "me at the red lobster",
    public: {
      apiUrl: "me at the red lobster",
    },
  },

  icon: {
    clientBundle: {
      scan: true,
    },
  },

  devtools: { enabled: true },
});
