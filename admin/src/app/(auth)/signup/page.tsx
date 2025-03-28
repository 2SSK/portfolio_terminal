import AuthForm from "@/components/AuthForm";
import { LinkTxt } from "@/components/Link";

import { Metadata } from "next";

export const metadata: Metadata = {
  title: "SignUp",
};

export default function SignupPage() {
  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <div className="bg-white p-6 rounded-lg shadow-md w-full max-w-md">
        <h1 className="text-2xl font-bold mb-6 text-center">Sign Up</h1>
        <AuthForm authType="signup" />
        <p className="mt-4 text-center">
          Already have an account? <LinkTxt title="Login" route="/login" />
        </p>
      </div>
    </div>
  );
}
