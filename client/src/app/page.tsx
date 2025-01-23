import { Tent } from "@/icons";
import { NavBar } from "../components/navbar";
import { FrutigerButton } from "@/frutiger/button";
import { FrutigerPanel } from "@/frutiger/panel";

export default async function Home() {
  return (
    <>
      <NavBar />
      <div className="h-screen bg-linear-to-tl from-cyan-500 to-blue-500">
        <Tent width={48} height={48} />
        <div>Hello world</div>
        <FrutigerButton>Click Me</FrutigerButton>
        <FrutigerPanel>
          <div>Panel content</div>
          <div>More stuff</div>
        </FrutigerPanel>
      </div>
    </>
  );
}
