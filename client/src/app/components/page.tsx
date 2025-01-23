import Link from "next/link";

export default function Page() {
  return (
    <div className="flex flex-col p-12">
      <h1>Components</h1>
      <Link href="components/buttons">Buttons</Link>
    </div>
  );
}
