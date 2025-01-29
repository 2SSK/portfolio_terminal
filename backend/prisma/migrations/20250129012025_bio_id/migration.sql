-- DropIndex
DROP INDEX "Bio_name_key";

-- AlterTable
ALTER TABLE "Bio" ADD COLUMN     "id" SERIAL NOT NULL,
ADD CONSTRAINT "Bio_pkey" PRIMARY KEY ("id");
