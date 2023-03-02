import { ReactElement, HTMLProps } from "react";
import { UseFormRegisterReturn } from "react-hook-form";

interface TextareaProps extends HTMLProps<HTMLTextAreaElement> {
  icon?: ReactElement;
  config?: UseFormRegisterReturn;
  py?: string | number;
}

export default function Textarea(props: TextareaProps) {
  return (
    <div className="relative flex items-center w-full">
      <textarea
        className={`${
          props.py ? "py-" + props.py.toString() : "py-3"
        } bg-slate-100 text-slate-800 dark:bg-slate-700/50 w-full px-6 block rounded-xl outline-none focus:ring-2 focus:ring-sky-500 dark:text-white focus:bg-slate-50 dark:focus:bg-slate-600 h-12`}
        {...props.config}
        {...props}
      ></textarea>
      <div className="text-slate-400 absolute right-6">
        {props.icon || null}
      </div>
    </div>
  );
}
