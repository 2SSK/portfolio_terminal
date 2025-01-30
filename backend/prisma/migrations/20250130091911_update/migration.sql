/*
  Warnings:

  - The primary key for the `Database` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - You are about to drop the column `id` on the `Database` table. All the data in the column will be lost.
  - You are about to drop the column `skillId` on the `Database` table. All the data in the column will be lost.
  - The primary key for the `ProgrammingLang` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - You are about to drop the column `id` on the `ProgrammingLang` table. All the data in the column will be lost.
  - You are about to drop the column `skillId` on the `ProgrammingLang` table. All the data in the column will be lost.
  - The primary key for the `Tool` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - You are about to drop the column `id` on the `Tool` table. All the data in the column will be lost.
  - You are about to drop the column `skillId` on the `Tool` table. All the data in the column will be lost.
  - You are about to drop the `Skill` table. If the table is not empty, all the data it contains will be lost.

*/
-- DropForeignKey
ALTER TABLE "Database" DROP CONSTRAINT "Database_skillId_fkey";

-- DropForeignKey
ALTER TABLE "ProgrammingLang" DROP CONSTRAINT "ProgrammingLang_skillId_fkey";

-- DropForeignKey
ALTER TABLE "Tool" DROP CONSTRAINT "Tool_skillId_fkey";

-- AlterTable
ALTER TABLE "Database" DROP CONSTRAINT "Database_pkey",
DROP COLUMN "id",
DROP COLUMN "skillId";

-- AlterTable
ALTER TABLE "ProgrammingLang" DROP CONSTRAINT "ProgrammingLang_pkey",
DROP COLUMN "id",
DROP COLUMN "skillId";

-- AlterTable
ALTER TABLE "Tool" DROP CONSTRAINT "Tool_pkey",
DROP COLUMN "id",
DROP COLUMN "skillId";

-- DropTable
DROP TABLE "Skill";
