import { Icon } from "../lib/Icon";
import { IconNode, SVGIconProps } from "../lib/types";

const _iconNode: IconNode = [["ellipse",{"cx":"12","cy":"5","rx":"9","ry":"3","key":"m6lj3jdx"}],["path",{"d":"M3 12a9 3 0 0 0 5 2.69M21 9.3V5","key":"m6lj3jdx"}],["path",{"d":"M3 5v14a9 3 0 0 0 6.47 2.88M12 12v4h4","key":"m6lj3jdx"}],["path",{"d":"M13 20a5 5 0 0 0 9-3 4.5 4.5 0 0 0-4.5-4.5c-1.33 0-2.54.54-3.41 1.41L12 16","key":"m6lj3jdx"}]];

export const DatabaseBackup = (props: SVGIconProps) => (
  <Icon iconNode={_iconNode} {...props} />
);
