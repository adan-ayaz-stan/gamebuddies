// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  modules: [
    "@nuxt/ui",
    "@tresjs/nuxt",
    [
      "@nuxtjs/google-fonts",
      {
        families: {
          "Big+Shoulders+Display": true,
          Iceland: true,
        },
      },
    ],
    // "@nuxtjs/kinde",
    "@vee-validate/nuxt",
    // "@prisma/nuxt",
  ],
  css: ["@/assets/css/base.css"],
});
