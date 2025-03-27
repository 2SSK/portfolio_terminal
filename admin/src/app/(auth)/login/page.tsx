import LoginForm from "./LoginForm";
import { LinkTxt } from "@/compnents/Link";

export default function LoginPage() {
  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <div className="bg-white p-6 rounded-lg shadow-md w-full max-w-md">
        <h1 className="text-2xl font-bold mb-6 text-center">Login</h1>
        <LoginForm />
        <p className="mt-4 text-center">
          Don&apos;t have an account?{" "}
          <LinkTxt title="Sign Up" route="/signup" />
        </p>
      </div>
    </div>
  );
}
