export default defineNuxtConfig({
  modules: [
    "nuxt-auth-utils",
    "@nuxt/ui",
    "@vueuse/nuxt",
    "@nuxt/image",
  ],
  css: ["./app/assets/global.css"],
  devtools: { enabled: true },
  future: { compatibilityVersion: 4 },
  compatibilityDate: "2024-04-03",
});