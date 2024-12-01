export default defineAppConfig({
  ui: {
    colors: {
      primary: "green",
      neutral: "neutral",
    },
    card: {
      slots: {
        header: "p-0 sm:px-0 mx-6 py-6",
        root: "shadow-lg dark:shadow-xl p-0 sm:p-0",
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
