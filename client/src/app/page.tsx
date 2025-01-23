import { Tent } from "@/icons";

export default async function Home() {
  return (
    <div className="bg-brand">
      <Tent width={48} height={48} />
      <div>Hello world</div>
    </div>
  );
}
