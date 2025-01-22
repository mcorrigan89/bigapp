import Image from "next/image";

import { createClient } from "@connectrpc/connect";
import { UserService } from "@/api/user/v1/user_pb";
import { createConnectTransport } from "@connectrpc/connect-node";

const transport = createConnectTransport({
  baseUrl: "http://localhost:3001",
  httpVersion: "2",
});

async function getUserById(id: string) {
  const client = createClient(UserService, transport);
  const res = await client.getUserById({
    id,
  });
  return res;
}
export default async function Home() {
  const response = await getUserById("30de702f-4bfb-426a-8d08-71ff696bc8de");
  return (
    <div>
      <div>Hello world</div>
      <div>{JSON.stringify(response)}</div>
    </div>
  );
}
