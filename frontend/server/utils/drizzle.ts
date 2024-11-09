import { drizzle } from "drizzle-orm/node-postgres";
import * as schema from "../database/migrations/schema";
export { sql, eq, and, or } from "drizzle-orm";

export const tables = schema;

export function useDB() {
  return drizzle(process.env.DATABASE_URL!, { schema });
}

export type Profile = typeof tables.profile.$inferSelect;
