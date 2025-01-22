import { createClient } from "@connectrpc/connect";
import { UserService } from "@/api/user/v1/user_pb";
import { createConnectTransport } from "@connectrpc/connect-node";

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
