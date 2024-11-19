export default defineAppConfig({
  ui: {
    colors: {
      primary: "green",
      neutral: "neutral",
    },
    card: {
      slots: {
        root: "bg-[var(--ui-bg-elevated)] ring ring-[var(--ui-border)] divide-y divide-[var(--ui-border)] rounded-[calc(var(--ui-radius)*2)] shadow-md",
        header: "p-4 sm:px-4",
        body: "p-4 sm:p-4",
        footer: "p-4 sm:px-4",
      },
    },
  },
});
