import { loginWithRefLink } from "@/api/client";
import { cookies } from "next/headers";
import { redirect } from "next/navigation";
import { NextResponse, type NextRequest } from "next/server";

export async function GET(request: NextRequest) {
  const searchParams = request.nextUrl.searchParams;
  const referenceLinkToken = searchParams.get("token");

  if (!referenceLinkToken) {
    return redirect("/");
  }

  try {
    const response = await loginWithRefLink({
      refLinkToken: referenceLinkToken,
    });
    if (!response.session?.token) {
      return redirect("/");
    }
    const cookieStore = await cookies();
    cookieStore.set("x-session-token", response.session.token);
  } catch (err) {
    console.error(err);
    return redirect("/");
  } finally {
    return redirect("/me");
  }
}
