"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { login, signup } from "@/actions/auth";

type AuthType = "login" | "signup";

interface AuthFormProps {
  authType: AuthType;
  redirectPath?: string;
}

export default function AuthForm({
  authType,
  redirectPath = "/",
}: AuthFormProps) {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const router = useRouter();

  const authAction = authType === "login" ? login : signup;

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);
    setError("");
    try {
      await authAction(email, password);
      router.push(redirectPath);
    } catch (err) {
      setError(
        err instanceof Error
          ? err.message
          : `${authType.charAt(0).toUpperCase() + authType.slice(1)} failed`,
      );
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <form onSubmit={handleSubmit} className="space-y-4">
      <input
        type="text"
        placeholder="Email"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
        name="email"
        className="w-full p-2 border rounded"
        disabled={isLoading}
      />
      <input
        type="password"
        placeholder="Password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
        name="password"
        className="w-full p-2 border rounded"
        disabled={isLoading}
      />
      {error && <p className="text-red-500 text-sm">{error}</p>}
      <button
        type="submit"
        className="w-full font-semibold bg-blue-500 text-white p-2 rounded hover:bg-blue-600 disabled:bg-blue-300 disabled:cursor-not-allowed"
        disabled={isLoading}
      >
        {isLoading
          ? "Processing..."
          : authType === "login"
            ? "Login"
            : "Sign Up"}
      </button>
    </form>
  );
}
