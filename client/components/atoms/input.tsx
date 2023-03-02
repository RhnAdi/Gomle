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
        } bg-slate-100 text-slate-800 dark:bg-slate-700/50 w-full px-6 block rounded-xl outline-none focus:ring-2 focus:ring-sky-500 dark:text-white focus:bg-slate-50 dark:focus:bg-slate-600`}
        {...props.config}
        {...props}
      />
      <div className="text-slate-400 absolute right-6">
        {props.icon || null}
      </div>
    </div>
  );
}
