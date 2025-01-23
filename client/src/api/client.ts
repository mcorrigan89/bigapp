import { createClient } from "@connectrpc/connect";
import { UserService } from "@/api/gen/user/v1/user_pb";
import { createConnectTransport } from "@connectrpc/connect-node";
import { cookies } from "next/headers";

const transport = createConnectTransport({
  baseUrl: "http://localhost:3001",
  httpVersion: "2",
});

export async function userById(id: string) {
  const client = createClient(UserService, transport);
  const res = await client.getUserById({
    id,
  });
  return res;
}

interface CreateUserArgs {
  email: string;
  familyName?: string;
  givenName?: string;
}

export async function createUser({
  email,
  familyName,
  givenName,
}: CreateUserArgs) {
  const client = createClient(UserService, transport);
  const res = await client.createUser({
    email,
    familyName,
    givenName,
  });
  return res;
}

export async function userByToken(token: string) {
  const client = createClient(UserService, transport);
  const res = await client.getUserBySessionToken({
    token,
  });
  return res;
}

export async function loginEmail({ email }: { email: string }) {
  const client = createClient(UserService, transport);
  const res = await client.createLoginEmail({
    email,
  });
  return res;
}

export async function loginWithRefLink({
  refLinkToken,
}: {
  refLinkToken: string;
}) {
  const client = createClient(UserService, transport);
  const res = await client.loginWithReferenceLink({
    token: refLinkToken,
  });
  return res;
}

export async function uploadImage({ file }: { file: File }) {
  const formData = new FormData();
  formData.append("image", file);

  const cookieJar = await cookies();
  const sessionToken = cookieJar.get("x-session-token");
  if (!sessionToken?.value) {
    throw new Error("Session token is missing");
  }

  const headers = new Headers();
  headers.append("x-session-token", sessionToken?.value);

  const res = await fetch("http://localhost:3001/image/upload", {
    method: "POST",
    body: formData,
    headers,
  });
  return res;
}
