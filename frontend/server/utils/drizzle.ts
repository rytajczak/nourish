import { drizzle } from "drizzle-orm/node-postgres";
export { sql, eq, and, or } from "drizzle-orm";

import * as schema from "../database/migrations/schema";

export const tables = schema;

export function useDB() {
  return drizzle(process.env.DATABASE_URL!, { schema });
}

export type Profile = typeof tables.profile.$inferSelect;
