"use server";

import { cookies } from "next/headers";
import { revalidatePath } from "next/cache";

export interface BioResponse {
  image: string;
  name: string;
  title: string;
  description: string;
}

const API_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:3000";

// Fetch bio data from backend
export async function fetchBioData(): Promise<BioResponse | null> {
  const cookieStore = await cookies();
  const token = cookieStore.get("accessToken")?.value;

  if (!token) {
    throw new Error("No access token found");
  }

  try {
    const response = await fetch(`${API_URL}/api/bio`, {
      method: "GET",
      headers: {
        Authorization: `Bearer ${token}`,
        "Content-Type": "application/json",
      },
    });

    if (!response.ok) {
      throw new Error(`Failed to fetch bio: ${response.status}`);
    }

    const data = await response.json();
    if (!data.bio) {
      throw new Error("Bio data not found in response");
    }

    return data.bio as BioResponse;
  } catch (error) {
    console.error("Error fetching bio:", error);
    return null;
  }
}

// Update bio data on backend
export async function updateBioData(
  formData: FormData,
): Promise<{ success: boolean; error?: string }> {
  const cookieStore = await cookies();
  const token = cookieStore.get("accessToken")?.value;

  if (!token) {
    return { success: false, error: "No access token found" };
  }

  try {
    const response = await fetch(`${API_URL}/api/bio`, {
      method: "POST",
      headers: {
        Authorization: `Bearer ${token}`,
      },
      body: formData, // Send multipart/form-data directly
    });

    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(
        errorData.error || `Failed to update bio: ${response.status}`,
      );
    }

    const data = await response.json();
    if (!data.bio) {
      throw new Error("Updated bio not returned");
    }

    // Revalidate the bio page to reflect changes
    revalidatePath("/bio");
    return { success: true };
  } catch (error) {
    console.error("Error updating bio:", error);
    return {
      success: false,
      error: error instanceof Error ? error.message : "Unknown error",
    };
  }
}
