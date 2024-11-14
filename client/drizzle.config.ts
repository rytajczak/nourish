import "dotenv/config";
import { defineConfig } from "drizzle-kit";

export default defineConfig({
  dialect: "postgresql",
  schema: "./server/database/migrations/schema.ts",
  out: "./server/database/migrations",
  dbCredentials: {
    url: process.env.DATABASE_URL!,
  },
});
