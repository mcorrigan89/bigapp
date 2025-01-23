import { Tent } from "@/icons";
import { NavBar } from "../components/navbar";

export default async function Home() {
  return (
    <>
      <NavBar />
      <div className="bg-negative-400 dark:bg-info-700">
        <Tent width={48} height={48} />
        <div>Hello world</div>
      </div>
    </>
  );
}
