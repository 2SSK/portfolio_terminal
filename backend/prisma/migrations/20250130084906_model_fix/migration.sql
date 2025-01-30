/*
  Warnings:

  - You are about to drop the `Databases` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `Skills` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `Tools` table. If the table is not empty, all the data it contains will be lost.

*/
-- DropForeignKey
ALTER TABLE "Databases" DROP CONSTRAINT "Databases_skillId_fkey";

-- DropForeignKey
ALTER TABLE "ProgrammingLang" DROP CONSTRAINT "ProgrammingLang_skillId_fkey";

-- DropForeignKey
ALTER TABLE "Tools" DROP CONSTRAINT "Tools_skillId_fkey";

-- DropTable
DROP TABLE "Databases";

-- DropTable
DROP TABLE "Skills";

-- DropTable
DROP TABLE "Tools";

-- CreateTable
CREATE TABLE "Skill" (
    "id" SERIAL NOT NULL,

    CONSTRAINT "Skill_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Tool" (
    "id" SERIAL NOT NULL,
    "toolName" TEXT NOT NULL,
    "skillId" INTEGER,

    CONSTRAINT "Tool_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Database" (
    "id" SERIAL NOT NULL,
    "databaseName" TEXT NOT NULL,
    "skillId" INTEGER,

    CONSTRAINT "Database_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "Tool_toolName_key" ON "Tool"("toolName");

-- CreateIndex
CREATE UNIQUE INDEX "Database_databaseName_key" ON "Database"("databaseName");

-- AddForeignKey
ALTER TABLE "ProgrammingLang" ADD CONSTRAINT "ProgrammingLang_skillId_fkey" FOREIGN KEY ("skillId") REFERENCES "Skill"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Tool" ADD CONSTRAINT "Tool_skillId_fkey" FOREIGN KEY ("skillId") REFERENCES "Skill"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Database" ADD CONSTRAINT "Database_skillId_fkey" FOREIGN KEY ("skillId") REFERENCES "Skill"("id") ON DELETE SET NULL ON UPDATE CASCADE;
