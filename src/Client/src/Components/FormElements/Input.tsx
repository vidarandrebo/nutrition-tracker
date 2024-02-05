import { AriaAttributes, DetailedHTMLProps, FC, InputHTMLAttributes } from "react";

export interface InputProps
    extends DetailedHTMLProps<InputHTMLAttributes<HTMLInputElement>, HTMLInputElement>,
        AriaAttributes {}

export const InputPrimary: FC<InputProps> = (props) => {
    const { ...attributes } = props;

    return (
        <input
            className="shadow appearance-none border hover:border-gray-500 rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
            {...attributes}
        />
    );
};
