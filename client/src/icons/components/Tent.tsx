import { Icon } from "../lib/Icon";
import { IconNode, SVGIconProps } from "../lib/types";

const _iconNode: IconNode = [["path",{"d":"M3.5 21 14 3M20.5 21 10 3M15.5 21 12 15l-3.5 6M2 21h20","key":"m68w66c8"}]];

export const Tent = (props: SVGIconProps) => (
  <Icon iconNode={_iconNode} {...props} />
);
