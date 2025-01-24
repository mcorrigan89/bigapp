"use server";

import { createClient } from "@connectrpc/connect";
import { UserService } from "@/api/gen/user/v1/user_pb";
import { createConnectTransport } from "@connectrpc/connect-node";
import { cookies } from "next/headers";
import { cache } from "react";
import { env } from "@/env";

const transport = createConnectTransport({
  baseUrl: env.SERVER_URL,
  httpVersion: "2",
});

async function getHeaders() {
  const cookieJar = await cookies();
  const sessionToken = cookieJar.get("x-session-token");
  const headers = new Headers();
  headers.set("Server-Token", env.SERVER_TOKEN);
  if (sessionToken?.value) {
    headers.append("x-session-token", sessionToken.value);
  }
  return { headers, token: sessionToken?.value };
}

async function userByIdFunc(id: string) {
  const client = createClient(UserService, transport);
  const { headers } = await getHeaders();
  const res = await client.getUserById({ id }, { headers });
  return res;
}

export const userById = cache(userByIdFunc);

interface CreateUserArgs {
  email: string;
  familyName?: string;
  givenName?: string;
}

export async function createUser({ email, familyName, givenName }: CreateUserArgs) {
  const client = createClient(UserService, transport);
  const { headers } = await getHeaders();
  const res = await client.createUser({ email, familyName, givenName }, { headers });
  return res;
}

export async function userByTokenFunc(token: string) {
  const client = createClient(UserService, transport);
  const { headers } = await getHeaders();
  const res = await client.getUserBySessionToken({ token }, { headers });
  return res;
}

export const userByToken = cache(userByTokenFunc);

export async function getCurrentUserFunc() {
  const { token, headers } = await getHeaders();
  if (!token) {
    return null;
  }
  const client = createClient(UserService, transport);
  const res = await client.getUserBySessionToken({ token }, { headers });
  return res;
}

export const getCurrentUser = cache(getCurrentUserFunc);

export async function loginEmail({ email }: { email: string }) {
  const client = createClient(UserService, transport);
  const { headers } = await getHeaders();
  const res = await client.createLoginEmail({ email }, { headers });
  return res;
}

export async function loginWithRefLink({ refLinkToken }: { refLinkToken: string }) {
  const client = createClient(UserService, transport);
  const { headers } = await getHeaders();
  const res = await client.loginWithReferenceLink({ token: refLinkToken }, { headers });
  return res;
}

export async function inviteUser({ email }: { email: string }) {
  const client = createClient(UserService, transport);
  const { headers } = await getHeaders();
  const res = await client.inviteUser({ email }, { headers });
  return res;
}

export async function acceptInviteRefLink({ refLinkToken }: { refLinkToken: string }) {
  const client = createClient(UserService, transport);
  const { headers } = await getHeaders();
  const res = await client.acceptInviteReferenceLink({ token: refLinkToken }, { headers });
  return res;
}

export async function uploadImage({ file }: { file: File }) {
  const formData = new FormData();
  formData.append("image", file);

  const { headers, token } = await getHeaders();
  if (!token) {
    throw new Error("Session token is missing");
  }

  const res = await fetch(`${env.SERVER_URL}/image/upload`, {
    method: "POST",
    body: formData,
    headers,
  });
  return res;
}
