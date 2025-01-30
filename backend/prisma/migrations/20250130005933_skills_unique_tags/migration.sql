/*
  Warnings:

  - A unique constraint covering the columns `[databaseName]` on the table `Databases` will be added. If there are existing duplicate values, this will fail.
  - A unique constraint covering the columns `[languageName]` on the table `ProgrammingLang` will be added. If there are existing duplicate values, this will fail.
  - A unique constraint covering the columns `[toolName]` on the table `Tools` will be added. If there are existing duplicate values, this will fail.

*/
-- CreateIndex
CREATE UNIQUE INDEX "Databases_databaseName_key" ON "Databases"("databaseName");

-- CreateIndex
CREATE UNIQUE INDEX "ProgrammingLang_languageName_key" ON "ProgrammingLang"("languageName");

-- CreateIndex
CREATE UNIQUE INDEX "Tools_toolName_key" ON "Tools"("toolName");
