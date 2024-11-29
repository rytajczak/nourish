export default defineAppConfig({
  ui: {
    colors: {
      primary: "green",
      neutral: "neutral",
    },
    card: {
      slots: {
        root: "shadow-md bg-elevated",
      },
    },
    skeleton: {
      base: "bg-neutral-200 dark:bg-neutral-800",
    },
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
