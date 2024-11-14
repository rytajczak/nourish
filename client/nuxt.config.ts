export default defineNuxtConfig({
  future: { compatibilityVersion: 4 },
  compatibilityDate: "2024-04-03",
  modules: [
    "@nuxt/eslint",
    "nuxt-auth-utils",
    "@nuxt/ui",
    "@vueuse/nuxt",
    "@nuxt/image",
    "@pinia/nuxt",
    "pinia-plugin-persistedstate/nuxt",
    "@vueform/nuxt",
  ],
  css: ["./app/assets/global.css"],
  app: { pageTransition: { name: "page", mode: "out-in" } },
  ssr: false,
  runtimeConfig: {
    public: {
      apiUrl:
        process.env.NODE_ENV === "production"
          ? "https://api.helpmemealprep.com"
          : "http://api.localhost",
    },
  },
  devtools: { enabled: true },
});
