import { Button } from "@/components/button";

export default function ButtonsPage() {
  return (
    <div className="flex flex-col gap-4 p-12">
      <div className="flex gap-4">
        <div>
          <Button variant={"default"} size={"small"}>
            Default
          </Button>
        </div>
        <div>
          <Button variant={"default"} size={"default"}>
            Default
          </Button>
        </div>
        <div>
          <Button variant={"default"} size={"large"}>
            Default
          </Button>
        </div>
      </div>
      <div className="flex gap-4">
        <div>
          <Button variant={"positive"} size={"small"}>
            Default
          </Button>
        </div>
        <div>
          <Button variant={"positive"} size={"default"}>
            Default
          </Button>
        </div>
        <div>
          <Button variant={"positive"} size={"large"}>
            Default
          </Button>
        </div>
      </div>
      <div className="flex gap-4">
        <div>
          <Button variant={"negative"} size={"small"}>
            Default
          </Button>
        </div>
        <div>
          <Button variant={"negative"} size={"default"}>
            Default
          </Button>
        </div>
        <div>
          <Button variant={"negative"} size={"large"}>
            Default
          </Button>
        </div>
      </div>
    </div>
  );
}
