/*
  Warnings:

  - A unique constraint covering the columns `[contactType]` on the table `Contacts` will be added. If there are existing duplicate values, this will fail.
  - A unique constraint covering the columns `[company]` on the table `Experience` will be added. If there are existing duplicate values, this will fail.
  - A unique constraint covering the columns `[title]` on the table `Project` will be added. If there are existing duplicate values, this will fail.
  - A unique constraint covering the columns `[title]` on the table `Socials` will be added. If there are existing duplicate values, this will fail.

*/
-- CreateIndex
CREATE UNIQUE INDEX "Contacts_contactType_key" ON "Contacts"("contactType");

-- CreateIndex
CREATE UNIQUE INDEX "Experience_company_key" ON "Experience"("company");

-- CreateIndex
CREATE UNIQUE INDEX "Project_title_key" ON "Project"("title");

-- CreateIndex
CREATE UNIQUE INDEX "Socials_title_key" ON "Socials"("title");
