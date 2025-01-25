import { Tent } from "@/icons";

export default async function Home() {
  return (
    <>
      <div className="h-screen bg-linear-to-tl from-cyan-500 to-blue-500">
        <Tent width={48} height={48} />
        <div>Hello world</div>
      </div>
    </>
  );
}
