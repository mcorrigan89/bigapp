import { Tent } from "@/icons";
import { Navbar } from "./components/navbar/navbar";

export default async function Home() {
  return (
    <>
      <Navbar />
      <div className="bg-brand">
        <Tent width={48} height={48} />
        <div>Hello world</div>
      </div>
    </>
  );
}
