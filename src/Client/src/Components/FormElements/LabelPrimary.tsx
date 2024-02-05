import { AriaAttributes, DetailedHTMLProps, FC, LabelHTMLAttributes } from "react";

export interface LabelProps
    extends DetailedHTMLProps<LabelHTMLAttributes<HTMLLabelElement>, HTMLLabelElement>,
        AriaAttributes {}

export const LabelPrimary: FC<LabelProps> = (props) => {
    const { children, ...attributes } = props;

    return (
        <label className="block text-gray-700 text-sm font-bold" {...attributes}>
            {children}
        </label>
    );
};
