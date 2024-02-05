import { AriaAttributes, ButtonHTMLAttributes, DetailedHTMLProps, FC } from "react";

export interface ButtonProps
    extends DetailedHTMLProps<ButtonHTMLAttributes<HTMLButtonElement>, HTMLButtonElement>,
        AriaAttributes {}

export const ButtonPrimary: FC<ButtonProps> = (props) => {
    const { children, ...attributes } = props;

    return (
        <button
            className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
            {...attributes}
        >
            {children}
        </button>
    );
};
