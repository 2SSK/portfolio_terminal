"use server";

import { cookies } from "next/headers";
import { redirect } from "next/navigation";

// Define the expected response type from your backend
interface AuthResponse {
  id: number;
  email: string;
  accessToken: string;
  refreshToken?: string; // Made optional since refresh endpoint might not return it
}

// Configuration for your API
const API_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:3000";

// Verify JWT token with your backend
async function verifyToken(token: string): Promise<AuthResponse | null> {
  try {
    const response = await fetch(`${API_URL}/api/user/verify`, {
      method: "GET",
      headers: {
        Authorization: `Bearer ${token}`,
        "Content-Type": "application/json",
      },
    });

    if (!response.ok) {
      return null;
    }

    const data = await response.json();
    return data as AuthResponse;
  } catch (error) {
    console.error("Token verification failed:", error);
    return null;
  }
}

// Server action to check authentication status
export async function checkAuth() {
  const cookieStore = await cookies(); // Await the cookies() promise
  const accessToken = cookieStore.get("accessToken")?.value;

  // If no token exists, redirect to login
  if (!accessToken) {
    redirect("/login");
  }

  // Verify the token with the backend
  const user = await verifyToken(accessToken);

  if (!user) {
    // If token is invalid or expired, try to refresh it
    const refreshToken = cookieStore.get("refreshToken")?.value;

    if (refreshToken) {
      const refreshResponse = await fetch(`${API_URL}/api/user/refresh`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ refreshToken }),
      });

      if (refreshResponse.ok) {
        const newTokens = await refreshResponse.json();
        // Update cookies with new tokens
        cookieStore.set("accessToken", newTokens.accessToken, {
          path: "/",
          httpOnly: true,
          secure: process.env.NODE_ENV === "production",
          maxAge: 7 * 24 * 60 * 60, // 7 days
        });
        return newTokens as AuthResponse;
      }
    }

    // If refresh fails or no refresh token, redirect to login
    cookieStore.delete("accessToken");
    cookieStore.delete("refreshToken");
    redirect("/login");
  }

  // If we get here, user is authenticated
  return user;
}

// Add this to your auth.ts file
export async function login(email: string, password: string) {
  // Talk to the backend to get the keys
  const response = await fetch(`${API_URL}/api/user`, {
    method: "POST", // We’re sending info, so POST
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ email, password }), // Send email and password
  });

  // If the computer says no (wrong email or password)
  if (!response.ok) {
    throw new Error("Oops! Wrong email or password!");
  }

  // Get the keys from the computer
  const data = (await response.json()) as AuthResponse;

  // Open your backpack
  const cookieStore = await cookies();

  // Put the accessToken in your backpack (it’s always there!)
  cookieStore.set("accessToken", data.accessToken, {
    path: "/",
    httpOnly: true, // Keeps it safe
    secure: process.env.NODE_ENV === "production",
    maxAge: 7 * 24 * 60 * 60, // 7 days
  });

  // Only put refreshToken in if it’s really there
  if (data.refreshToken) {
    cookieStore.set("refreshToken", data.refreshToken, {
      path: "/",
      httpOnly: true,
      secure: process.env.NODE_ENV === "production",
      maxAge: 30 * 24 * 60 * 60, // 30 days
    });
  }

  // Yay! You’re logged in!
  return data; // Give back the user info
}

// Helper function to handle logout
export async function logout() {
  const cookieStore = await cookies(); // Await the cookies() promise
  cookieStore.delete("accessToken");
  cookieStore.delete("refreshToken");
  redirect("/login");
}
