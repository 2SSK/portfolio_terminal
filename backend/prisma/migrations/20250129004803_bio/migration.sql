-- CreateTable
CREATE TABLE "Bio" (
    "dp" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "intro" TEXT NOT NULL,
    "description" TEXT NOT NULL
);

-- CreateIndex
CREATE UNIQUE INDEX "Bio_name_key" ON "Bio"("name");
