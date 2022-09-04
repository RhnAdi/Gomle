import { ReactElement, HTMLProps } from "react";
import { UseFormRegisterReturn } from "react-hook-form";

interface InputProps extends HTMLProps<HTMLInputElement> {
  icon?: ReactElement;
  config?: UseFormRegisterReturn;
  py?: string | number;
}

export default function Input(props: InputProps) {
  return (
    <div className="relative flex items-center w-full">
      <input
        className={`${
          props.py ? "py-" + props.py.toString() : "py-3"
        } bg-gray-100 text-gray-800 dark:bg-gray-600 w-full px-6 block rounded-xl outline-none focus:ring-2 focus:ring-sky-500 dark:text-white focus:bg-gray-50 dark:focus:bg-gray-600`}
        {...props.config}
        {...props}
      />
      <div className="text-gray-400 absolute right-6">{props.icon || null}</div>
    </div>
  );
}
