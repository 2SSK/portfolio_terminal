"use server";

import { cookies } from "next/headers";
import { redirect } from "next/navigation";

// Define the expected response type from your backend
interface AuthResponse {
  id: number;
  email: string;
  accessToken: string;
  refreshToken?: string;
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
  const cookieStore = await cookies();
  const accessToken = cookieStore.get("accessToken")?.value;

  if (!accessToken) {
    redirect("/login");
  }

  const user = await verifyToken(accessToken);

  if (!user) {
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
        cookieStore.set("accessToken", newTokens.accessToken, {
          path: "/",
          httpOnly: true,
          secure: process.env.NODE_ENV === "production",
          maxAge: 7 * 24 * 60 * 60,
        });
        return newTokens as AuthResponse;
      }
    }

    cookieStore.delete("accessToken");
    cookieStore.delete("refreshToken");
    redirect("/login");
  }

  return user;
}

// Login function
export async function login(email: string, password: string) {
  const response = await fetch(`${API_URL}/api/user/login`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ email, password }),
  });

  if (!response.ok) {
    throw new Error("Oops! Wrong email or password!");
  }

  const data = (await response.json()) as AuthResponse;
  const cookieStore = await cookies();

  cookieStore.set("accessToken", data.accessToken, {
    path: "/",
    httpOnly: true,
    secure: process.env.NODE_ENV === "production",
    maxAge: 7 * 24 * 60 * 60,
  });

  if (data.refreshToken) {
    cookieStore.set("refreshToken", data.refreshToken, {
      path: "/",
      httpOnly: true,
      secure: process.env.NODE_ENV === "production",
      maxAge: 30 * 24 * 60 * 60,
    });
  }

  return data;
}

// Signup function
export async function signup(email: string, password: string) {
  const response = await fetch(`${API_URL}/api/user/register`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ email, password }),
  });

  if (!response.ok) {
    const errorData = await response.json();
    throw new Error(errorData.error || "Failed to create account");
  }

  const data = (await response.json()) as AuthResponse;
  const cookieStore = await cookies();

  cookieStore.set("accessToken", data.accessToken, {
    path: "/",
    httpOnly: true,
    secure: process.env.NODE_ENV === "production",
    maxAge: 7 * 24 * 60 * 60, // 7 days
  });

  if (data.refreshToken) {
    cookieStore.set("refreshToken", data.refreshToken, {
      path: "/",
      httpOnly: true,
      secure: process.env.NODE_ENV === "production",
      maxAge: 30 * 24 * 60 * 60, // 30 days
    });
  }

  return data;
}

// Logout function
export async function logout() {
  const cookieStore = await cookies();
  cookieStore.delete("accessToken");
  cookieStore.delete("refreshToken");
  redirect("/login");
}
