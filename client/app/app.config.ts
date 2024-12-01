export default defineAppConfig({
  ui: {
    colors: {
      primary: "green",
      neutral: "neutral",
    },
    card: {
      slots: {
        root: "shadow-lg dark:shadow-xl p-0 sm:p-0 bg-white dark:bg-[#111111]",
      },
    },
    skeleton: {
      base: "bg-neutral-200 dark:bg-neutral-900",
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
