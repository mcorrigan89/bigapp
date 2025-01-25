"use server";

import { createClient } from "@connectrpc/connect";
import { UserService } from "@/api/gen/user/v1/user_pb";
import { createConnectTransport } from "@connectrpc/connect-node";
import { cookies } from "next/headers";
import { cache } from "react";
import { env } from "@/env";
import { ImageService } from "./gen/media/v1/image_pb";

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

async function userByHandleFunc(handle: string) {
  const client = createClient(UserService, transport);
  const { headers } = await getHeaders();
  const res = await client.getUserByHandle({ handle }, { headers });
  return res;
}

export const userByHandle = cache(userByHandleFunc);

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

interface UpdateUserArgs {
  id: string;
  email: string;
  familyName?: string;
  givenName?: string;
  handle: string;
}

export async function updateUser({ id, handle, email, familyName, givenName }: UpdateUserArgs) {
  const client = createClient(UserService, transport);
  const { headers } = await getHeaders();
  const res = await client.updateUser({ id, handle, email, familyName, givenName }, { headers });
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

export async function getCollectiondByIdFunc(collectionId: string) {
  const { token, headers } = await getHeaders();
  if (!token) {
    return null;
  }
  const client = createClient(ImageService, transport);
  const res = await client.getCollectionById({ collectionId }, { headers });
  return res;
}

export const getCollectiondById = cache(getCollectiondByIdFunc);

export async function getCollectiondByOwnerIdFunc(ownerId: string) {
  const { token, headers } = await getHeaders();
  if (!token) {
    return null;
  }
  const client = createClient(ImageService, transport);
  const res = await client.getCollectionByOwnerId({ ownerId }, { headers });
  return res;
}

export const getCollectiondByOwnerId = cache(getCollectiondByOwnerIdFunc);

export async function getCollectiondByOwnerTokenFunc() {
  const { token, headers } = await getHeaders();
  if (!token) {
    return null;
  }
  const client = createClient(ImageService, transport);
  const res = await client.getCollectionByOwnerToken({ token }, { headers });
  return res;
}

export const getCollectiondByOwnerToken = cache(getCollectiondByOwnerTokenFunc);

export async function createCollection(collectionName: string) {
  const { token, headers } = await getHeaders();
  if (!token) {
    return null;
  }
  const client = createClient(ImageService, transport);
  const res = await client.createCollection({ collectionName }, { headers });
  return res;
}

export async function uploadAvatarImage({ file }: { file: File }) {
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

export async function uploadImagesToCollection({ collectionId, files }: { collectionId: string; files: File[] }) {
  const formData = new FormData();
  files.forEach((file) => {
    formData.append("images", file);
  });

  const { headers, token } = await getHeaders();
  if (!token) {
    throw new Error("Session token is missing");
  }

  const res = await fetch(`${env.SERVER_URL}/image/${collectionId}/uploads`, {
    method: "POST",
    body: formData,
    headers,
  });
  return res;
}
