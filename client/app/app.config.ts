export default defineAppConfig({
  ui: {
    colors: {
      primary: "green",
      neutral: "neutral",
    },
    card: { slots: { root: "shadow-lg" } },
    navigationMenu: {
      variants: {
        orientation: {
          vertical: {
            root: "text-muted",
            link: "py-3",
            childList: "m-0 border-none",
          },
        },
      },
    },
  },
});
