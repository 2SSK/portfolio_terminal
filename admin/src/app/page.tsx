import { redirect } from "next/navigation";
import { checkAuth } from "@/actions/auth";

export default async function Home() {
  const user = await checkAuth();

  if (!user) {
    return redirect("/login");
  }

  return redirect("/bio");
}
