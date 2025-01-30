-- CreateTable
CREATE TABLE "Skills" (
    "id" SERIAL NOT NULL,

    CONSTRAINT "Skills_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "ProgrammingLang" (
    "id" SERIAL NOT NULL,
    "languageName" TEXT NOT NULL,
    "skillId" INTEGER,

    CONSTRAINT "ProgrammingLang_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Tools" (
    "id" SERIAL NOT NULL,
    "toolName" TEXT NOT NULL,
    "skillId" INTEGER,

    CONSTRAINT "Tools_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Databases" (
    "id" SERIAL NOT NULL,
    "databaseName" TEXT NOT NULL,
    "skillId" INTEGER,

    CONSTRAINT "Databases_pkey" PRIMARY KEY ("id")
);

-- AddForeignKey
ALTER TABLE "ProgrammingLang" ADD CONSTRAINT "ProgrammingLang_skillId_fkey" FOREIGN KEY ("skillId") REFERENCES "Skills"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Tools" ADD CONSTRAINT "Tools_skillId_fkey" FOREIGN KEY ("skillId") REFERENCES "Skills"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Databases" ADD CONSTRAINT "Databases_skillId_fkey" FOREIGN KEY ("skillId") REFERENCES "Skills"("id") ON DELETE SET NULL ON UPDATE CASCADE;
