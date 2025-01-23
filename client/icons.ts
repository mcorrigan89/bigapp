import path from "path";
import fs from "fs";
import { optimize } from "svgo";
import { parseSync } from "svgson";

const ICONS_DIR = path.resolve(process.cwd(), "src/icons");

const pascalCase = (s: string) => {
  s = s.charAt(0).toUpperCase() + s.slice(1);
  return s.replace(/-./g, (x) => x[1].toUpperCase());
};

const svgFiles = fs
  .readdirSync(ICONS_DIR + "/svgs")
  .filter((fileName) => path.extname(fileName) === ".svg")
  .map((fileName) => ({
    fileName,
    content: fs.readFileSync(
      path.resolve(ICONS_DIR + "/svgs", fileName),
      "utf-8",
    ),
  }))
  .map((svgFile) => ({
    content: optimize(svgFile.content),
    fileName: svgFile.fileName.split(".")[0],
  }))
  .map((svgFile) => ({
    parsedData: parseSync(svgFile.content.data),
    fileName: pascalCase(svgFile.fileName),
  }))
  .map((svgFile) => {
    const _iconNode = svgFile.parsedData.children.map((child) => [
      child.name,
      child.attributes,
    ]);
    return { [svgFile.fileName]: _iconNode };
  });

svgFiles.forEach((svgFile) => {
  const key = Object.keys(svgFile)[0];
  const value = svgFile[key];
  const iconComponent = `import { Icon } from "../lib/Icon";
import { IconNode, SVGIconProps } from "../lib/types";

const _iconNode: IconNode = ${JSON.stringify(value)};

export const ${key} = (props: SVGIconProps) => (
  <Icon iconNode={_iconNode} {...props} />
);
`;

  fs.writeFileSync(ICONS_DIR + "/components/" + key + ".tsx", iconComponent);
});

const indexExportFile = svgFiles
  .sort((a, b) => Object.keys(a)[0].localeCompare(Object.keys(b)[0]))
  .map((svgFile) => `export * from "./components/${Object.keys(svgFile)[0]}";`)
  .join("\n");
fs.writeFileSync(ICONS_DIR + "/index.ts", indexExportFile);
