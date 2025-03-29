import { checkAuth } from "@/actions/auth";
import { fetchBioData, updateBioData } from "@/actions/bio_action";
import { revalidatePath } from "next/cache";
import { redirect } from "next/navigation";
import BioView from "./BioView";

export default async function BioPage({
  searchParams,
}: {
  searchParams: { error?: string };
}) {
  // Ensure user is authenticated
  await checkAuth();

  // Fetch bio data
  const bio = await fetchBioData();

  // Server Action to handle bio update
  async function handleUpdateBio(formData: FormData) {
    "use server";
    try {
      const result = await updateBioData(formData);
      if (!result.success) {
        redirect(
          `/bio?error=${encodeURIComponent(result.error || "Failed to update bio")}`,
        );
      }
      revalidatePath("/bio"); // Revalidate cache for fresh data
      redirect("/bio"); // Redirect to bio page on success
    } catch {
      redirect(
        `/bio?error=${encodeURIComponent("An unexpected error occurred")}`,
      );
    }
  }

  return (
    <BioView
      bio={bio}
      error={
        searchParams.error ? decodeURIComponent(searchParams.error) : undefined
      }
      onUpdateBio={handleUpdateBio}
    />
  );
}

// Optional: Metadata for SEO (Next.js 13+ App Router)
export const metadata = {
  title: "Bio",
  description: "Manage your bio information",
};
