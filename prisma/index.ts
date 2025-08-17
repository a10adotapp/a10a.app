import { withAccelerate } from "@prisma/extension-accelerate";
import { PrismaClient } from "./generated/client/index.js";

const prismaClientSingleton = () => {
  const prismaClient = new PrismaClient();

  return prismaClient.$extends(withAccelerate());
};

const globalForPrisma = global as unknown as {
  prisma: PrismaClient,
};

const prisma = globalForPrisma.prisma || prismaClientSingleton();

if (process.env.NODE_ENV !== "production") {
  globalForPrisma.prisma = prisma;
}

export default prisma;
