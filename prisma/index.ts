import { PrismaMariaDb } from "@prisma/adapter-mariadb";
import { PrismaClient } from "./generated/client/index.js";

const prismaClientSingleton = (): PrismaClient => {
  const adapter = new PrismaMariaDb(process.env.DATABASE_URL!);

  return new PrismaClient({ adapter });
};

const globalForPrisma = global as unknown as {
  prisma: PrismaClient,
};

const prisma = globalForPrisma.prisma || prismaClientSingleton();

if (process.env.NODE_ENV !== "production") {
  globalForPrisma.prisma = prisma;
}

export default prisma;
